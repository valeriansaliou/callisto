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
  TIME_SECOND_TO_MILLISECONDS int = 1000
  TIME_HOUR_TO_MILLISECONDS int = 60 * 60 * TIME_SECOND_TO_MILLISECONDS
  TIME_DAY_TO_MILLISECONDS int = 24 * TIME_HOUR_TO_MILLISECONDS
  TIME_YEAR_TO_MILLISECONDS int = 365 * TIME_DAY_TO_MILLISECONDS

  TIME_START_FROM_MILLISECONDS int = TIME_YEAR_TO_MILLISECONDS

  TIME_NORMALIZE_FACTOR float32 = 64.0
)

// Window
const (
  WINDOW_WIDTH int = 1200
  WINDOW_HEIGHT int = 800

  WINDOW_TITLE string = "Callisto - Solar System Simulator"
)

// Controls
const (
  CONTROLS_ENABLE_KEY bool = true
  CONTROLS_ENABLE_MOUSE bool = true
)

// Speed
const (
  SPEED_FRAMERATE float64 = 60
)

// Projection
var (
  PROJECTION_FIELD_NEAR float32 = 0.1
  PROJECTION_FIELD_FAR float32 = 9999999999999999999.0
)

// Camera
var (
  CAMERA_DEFAULT_EYE mgl32.Vec3 = mgl32.Vec3{1600, -1800, -5657}
  CAMERA_DEFAULT_TARGET mgl32.Vec3 = mgl32.Vec3{0.255, 0.650, 0.000}

  CAMERA_MOVE_CELERITY_CRUISE float64 = 5.0
  CAMERA_MOVE_CELERITY_TURBO float64 = 20.0

  CAMERA_INERTIA_PRODUCE_FORWARD float64 = 0.05
  CAMERA_INERTIA_PRODUCE_BACKWARD float64 = -0.05
  CAMERA_INERTIA_CONSUME_FORWARD float64 = -0.04
  CAMERA_INERTIA_CONSUME_BACKWARD float64 = 0.04

  CAMERA_TARGET_AMORTIZE_FACTOR float32 = 0.005
)

// Object
const (
  OBJECT_TEXTURE_PHI_MAX int = 90
  OBJECT_TEXTURE_THETA_MAX int = 360
  OBJECT_TEXTURE_STEP_LATITUDE int = 3
  OBJECT_TEXTURE_STEP_LONGITUDE int = 6

  OBJECT_FULL_ANGLE float64 = 4.0 * math.Pi

  OBJECT_FACTOR_SIZE float64 = 0.25
  OBJECT_FACTOR_SPEED_SCENE float64 = float64(TIME_HOUR_TO_MILLISECONDS) / 100.0
  OBJECT_FACTOR_SPEED_MAXIMUM float64 = 200.0
)
