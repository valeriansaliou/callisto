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

type Circle ObjectElement

func generateCircleFilledFromObject(object *Object) (Circle) {
  return generateCircleFilled(object.Distance, object.Radius, object.Radiate)
}

func generateCircleFromObject(object *Object) (Circle) {
  return generateCircle(object.Radius, object.Radiate)
}

func generateCircleFilled(radius float32, thickness float32, radiate bool) (Circle) {
  var (
    circle           Circle

    i                int
    j                int
    k                int
    l                int

    normal_direction float32

    radius_inside_n  float32
    radius_outside_n float32
    nb_vertices      int32

    angle            float64
    angle_max        int
  )

  angle_max = 360

  circle.Vertices = make([]float32, 3 * angle_max * 2)
  circle.VerticeNormals = make([]float32, 3 * angle_max * 2)
  circle.Indices = make([]int32, 3 * (angle_max + 1) * 2)
  circle.TextureCoords = make([]float32, 2 * angle_max * 2)

  i = 0
  j = 0
  k = 0
  l = 0

  radius_inside_n = normalizeObjectSize(radius)
  radius_outside_n = normalizeObjectSize(radius + thickness)

  // Normal is -1 if sun, which is the light source, to avoid any self-shadow effect
  if radiate == true {
    normal_direction = -1.0
  } else {
    normal_direction = 1.0
  }

  for angle = 0.0; angle < float64(angle_max); angle++ {
    // Generate inner circle object
    generateCircleObject(&circle, radius_inside_n, angle, normal_direction, &nb_vertices, &i, &j, &k, &l)

    // Generate outer circle object? (if not last)
    generateCircleObject(&circle, radius_outside_n, angle, normal_direction, &nb_vertices, &i, &j, &k, &l)
  }

  return circle
}

func generateCircle(radius float32, radiate bool) (Circle) {
  var (
    circle           Circle

    i                int
    j                int
    k                int
    l                int

    normal_direction float32

    radius_n         float32
    nb_vertices      int32

    angle            float64
    angle_max        int
  )

  angle_max = 360

  circle.Vertices = make([]float32, 3 * angle_max * 2)
  circle.VerticeNormals = make([]float32, 3 * angle_max * 2)
  circle.Indices = make([]int32, 3 * (angle_max + 1) * 2)
  circle.TextureCoords = make([]float32, 2 * angle_max * 2)

  i = 0
  j = 0
  k = 0
  l = 0

  radius_n = normalizeObjectSize(radius)

  // Normal is -1 if sun, which is the light source, to avoid any self-shadow effect
  if radiate == true {
    normal_direction = -1.0
  } else {
    normal_direction = 1.0
  }

  for angle = 0.0; angle < float64(angle_max); angle++ {
    generateCircleObject(&circle, radius_n, angle, normal_direction, &nb_vertices, &i, &j, &k, &l)
  }

  return circle
}

func generateCircleObject(circle *Circle, radius_n float32, angle float64, normal_direction float32, nb_vertices *int32, i *int, j *int, k *int, l *int) {
  var (
    vertex_position_x  float32
    vertex_position_y  float32
    vertex_position_z  float32
  )

  // Generate inside circle vertices
  vertex_position_x = float32(math.Cos(MATH_DEG_TO_RAD * angle))
  vertex_position_y = 0.0
  vertex_position_z = float32(math.Sin(MATH_DEG_TO_RAD * angle))

  // Bind inside circle vertices
  circle.Vertices[*i] = radius_n * vertex_position_x
  circle.Vertices[*i + 1] = radius_n * vertex_position_y
  circle.Vertices[*i + 2] = radius_n * vertex_position_z

  *i += 3

  // Bind circle vertice normals
  circle.VerticeNormals[*j] = normal_direction * vertex_position_x
  circle.VerticeNormals[*j + 1] = normal_direction * vertex_position_y
  circle.VerticeNormals[*j + 2] = normal_direction * vertex_position_z

  *j += 3

  // Bind circle indices
  circle.Indices[*k] = *nb_vertices
  circle.Indices[*k + 1] = *nb_vertices + 1
  circle.Indices[*k + 2] = *nb_vertices + 2

  *k += 3

  *nb_vertices += 1.0

  // Bind circle texture coordinates
  circle.TextureCoords[*l] = 0.0
  circle.TextureCoords[*l + 1] = 0.0

  *l += 2
}
