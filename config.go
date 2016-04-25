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
  "github.com/go-gl/mathgl/mgl32"
)

// Math
const (
  MATH_DEG_TO_RAD float64 = math.Pi / 180
)

// Time
const (
  TIME_DAY_TO_MILLISECONDS int = 24 * 60 * 60 * 1000
  TIME_YEAR_TO_MILLISECONDS int = 365 * 24 * 60 * 60 * 1000
)

// Window
const (
  WINDOW_WIDTH int = 1200
  WINDOW_HEIGHT int = 800

  WINDOW_TITLE string = "Callisto - Solar System Simulator"
)

// Camera
var (
  CAMERA_DEFAULT_EYE mgl32.Vec3 = mgl32.Vec3{8, 3, 3}
  CAMERA_DEFAULT_CENTER mgl32.Vec3 = mgl32.Vec3{0, 0, 0}
  CAMERA_DEFAULT_UP mgl32.Vec3 = mgl32.Vec3{0, 1, 0}

  CAMERA_MOVE_CELERITY_FORWARD float32 = 0.01
  CAMERA_MOVE_CELERITY_BACKWARD float32 = -0.01

  CAMERA_WATCH_REDUCER float32 = 0.01
)

// Object
const (
  OBJECT_TEXTURE_PHI_MAX int = 90
  OBJECT_TEXTURE_THETA_MAX int = 360
  OBJECT_TEXTURE_STEP_LATITUDE int = 3
  OBJECT_TEXTURE_STEP_LONGITUDE int = 6

  OBJECT_ROTATION_FULL_ANGLE float64 = 2.0 * math.Pi
  OBJECT_REVOLUTION_FULL_ANGLE float64 = 2.0 * math.Pi

  OBJECT_FACTOR_RADIUS float32 = 0.00001436781609
  OBJECT_FACTOR_DISTANCE float32 = 0.00000000488536842
  OBJECT_FACTOR_SPEED_ROTATION float64 = 60 * 60 * 1000 * 4
  OBJECT_FACTOR_SPEED_REVOLUTION float64 = 60 * 60 * 1000 * 4 * 100
)
