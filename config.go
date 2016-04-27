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
  ConfigMathDegreeToRadian float64 = math.Pi / 180
)

// Time
const (
  ConfigTimeSecondToMilliseconds int = 1000
  ConfigTimeHourToMilliseconds int = 60 * 60 * ConfigTimeSecondToMilliseconds
  ConfigTimeDayToMilliseconds int = 24 * ConfigTimeHourToMilliseconds
  ConfigTimeYearToMilliseconds int = 365 * ConfigTimeDayToMilliseconds

  ConfigTimeStartFromMilliseconds int = ConfigTimeYearToMilliseconds

  ConfigTimeNormalizeFactor float32 = 64.0
)

// Window
const (
  ConfigWindowTitle string = "Callisto - Solar System Simulator"
)

// Speed
const (
  ConfigSpeedFramerateDefault float64 = 60
)

// Projection
var (
  ConfigProjectionFieldNear float32 = 0.1
  ConfigProjectionFieldFar float32 = 9999999999999999999.0
)

// Camera
var (
  ConfigCameraDefaultEye = mgl32.Vec3{1600, -1800, -5657}
  ConfigCameraDefaultTarget = mgl32.Vec3{0.255, 0.650, 0.000}

  ConfigCameraMoveCelerityCruise = 5.0
  ConfigCameraMoveCelerityTurbo = 50.0

  ConfigCameraInertiaProduceForward = 0.05
  ConfigCameraInertiaProduceBackward = -0.05
  ConfigCameraInertiaConsumeForward = -0.04
  ConfigCameraInertiaConsumeBackward = 0.04

  ConfigCameraTargetAmortizeFactor float32 = 0.005
)

// Object
const (
  ConfigObjectTexturePhiMax int = 90
  ConfigObjectTextureThetaMax int = 360
  ConfigObjectTextureStepLatitude int = 3
  ConfigObjectTextureStepLongitude int = 6

  ConfigObjectFullAngle float64 = 4.0 * math.Pi

  ConfigObjectFactorSize float64 = 0.25
  ConfigObjectFactorSpeedScene float64 = float64(ConfigTimeHourToMilliseconds) / 100.0
  ConfigObjectFactorSpeedMaximum float64 = 200.0
)
