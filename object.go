/* Callisto - Yet another Solar System simulator
 *
 * Copyright (c) 2016, Valerian Saliou <valerian@valeriansaliou.name>
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *   * Redistributions of source code must retain the above copyright notice,
 *     this list of conditions and the following disclaimer.
 *   * Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
  "math"
)

// Object format: [vertex<X, Y, Z>]

func generateSphere(longitudes int, latitudes int) ([]float32, []int32) {
  var (
    i               int
    j               int

    longitude       int
    latitude        int

    longitudes_f    float64
    latitudes_f     float64

    longitudes_i    int32
    latitudes_i     int32

    angle_longitude float64
    cos_longitude   float64
    sin_longitude   float64

    angle_latitude  float64
    cos_latitude    float64
    sin_latitude    float64
  )

  vertices := make([]float32, longitudes * latitudes * 3)
  indices := make([]int32, longitudes * (latitudes - 1) * 6)

  i = 0
  j = 0

  longitudes_f = float64(longitudes)
  latitudes_f = float64(latitudes)

  longitudes_i = int32(longitudes)
  latitudes_i = int32(latitudes)

  // Build sphere vertices
  for longitude = 0; longitude < longitudes; longitude++ {
    angle_longitude = (2.0 * float64(longitude) * math.Pi) / longitudes_f

    cos_longitude = math.Cos(angle_longitude)
    sin_longitude = math.Sin(angle_longitude)

    for latitude = 0; latitude < latitudes; latitude++ {
      angle_latitude = (float64(latitude) / (latitudes_f - 1.0) - 0.5) * math.Pi

      cos_latitude = math.Cos(angle_latitude)
      sin_latitude = math.Sin(angle_latitude)

      // Append sphere vertex
      vertices[i] = float32(cos_latitude * sin_longitude)
      vertices[i + 1] = float32(sin_latitude)
      vertices[i + 2] = float32(cos_latitude * cos_longitude)

      i += 3
    }
  }

  // Build sphere indices
  for longitude = 0; longitude < longitudes; longitude++ {
    for latitude = 0; latitude < (latitudes - 1); latitude++ {
      // ABC triangle
      indices[j] = int32(latitude) + int32(longitude) * latitudes_i
      indices[j + 1] = int32(latitude) + ((int32(longitude) + 1) % longitudes_i) * latitudes_i
      indices[j + 2] = (int32(latitude) + 1) + ((int32(longitude) + 1) % longitudes_i) * latitudes_i

      // ACD triangle
      indices[j + 3] = int32(latitude) + int32(longitude) * latitudes_i
      indices[j + 4] = (int32(latitude) + 1) + ((int32(longitude) + 1) % longitudes_i) * latitudes_i
      indices[j + 5] = (int32(latitude) + 1) + int32(longitude) * latitudes_i

      j += 6
    }
  }

  return vertices, indices
}
