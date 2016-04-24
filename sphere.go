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

type Sphere struct {
  Vertices      []float32
  Indices       []int32
  TextureCoords []float32
}

func generateSphere(name string, radius float32) (Sphere) {
  var (
    sphere           Sphere

    i                int
    j                int
    k                int

    radius_n         float32
    nb_vertices      float32

    unary_size_full  int
    unary_size_short int

    res_longitude    float32

    longitude        int
    latitude         int

    longitude_r_f    float64
    latitude_r_f     float64
  )

  unary_size_full = (2 * OBJECT_TEXTURE_PHI_MAX / OBJECT_TEXTURE_STEP_LATITUDE + 1) * (OBJECT_TEXTURE_THETA_MAX / OBJECT_TEXTURE_STEP_LONGITUDE + 1)
  unary_size_short = (2 * OBJECT_TEXTURE_PHI_MAX / OBJECT_TEXTURE_STEP_LATITUDE) * (OBJECT_TEXTURE_THETA_MAX / OBJECT_TEXTURE_STEP_LONGITUDE)

  sphere.Vertices = make([]float32, 3 * unary_size_full)
  sphere.Indices = make([]int32, 6 * unary_size_short)
  sphere.TextureCoords = make([]float32, 2 * unary_size_full)

  i = 0
  j = 0
  k = 0

  radius_n = radius * float32(OBJECT_FACTOR_RADIUS)
  nb_vertices = 0.0
  res_longitude = float32(OBJECT_TEXTURE_THETA_MAX) / float32(OBJECT_TEXTURE_STEP_LONGITUDE) + 1.0;

  // Map sphere data
  for latitude = -90; latitude <= OBJECT_TEXTURE_PHI_MAX; latitude += OBJECT_TEXTURE_STEP_LATITUDE {
    for longitude = 0; longitude <= OBJECT_TEXTURE_THETA_MAX; longitude += OBJECT_TEXTURE_STEP_LONGITUDE {
      // Convert latitude & longitude to radians
      longitude_r_f = float64(MATH_DEG_TO_RAD) * float64(longitude)
      latitude_r_f = float64(MATH_DEG_TO_RAD) * float64(latitude)

      // Bind sphere vertices
      sphere.Vertices[i] = radius_n * float32(math.Sin(longitude_r_f) * math.Cos(latitude_r_f))
      sphere.Vertices[i + 1] = radius_n * float32(math.Sin(latitude_r_f))
      sphere.Vertices[i + 2] = radius_n * float32(math.Cos(latitude_r_f) * math.Cos(longitude_r_f))

      i += 3

      // Bind sphere indices
      if (longitude != OBJECT_TEXTURE_THETA_MAX) {
        if (latitude < OBJECT_TEXTURE_PHI_MAX) {
          sphere.Indices[j] = int32(nb_vertices)
          sphere.Indices[j + 1] = int32(nb_vertices + 1.0)
          sphere.Indices[j + 2] = int32(nb_vertices + 1.0 + res_longitude)

          sphere.Indices[j + 3] = int32(nb_vertices)
          sphere.Indices[j + 4] = int32(nb_vertices + 1.0 + res_longitude)
          sphere.Indices[j + 5] = int32(nb_vertices + res_longitude)

          j += 6
        }
      }

      nb_vertices += 1.0

      // Bind sphere texture coordinates
      sphere.TextureCoords[k] = float32(longitude) / float32(OBJECT_TEXTURE_THETA_MAX)
      sphere.TextureCoords[k + 1] = float32(90 + latitude) / float32(90 + OBJECT_TEXTURE_PHI_MAX)

      k += 2
    }
  }

  return sphere
}
