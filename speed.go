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

// Speed  Maps current scene speed
type Speed struct {
  TimePrevious float64
  TimeElapsed  float64

  Factor       float64
  Framerate    float64
}

// InstanceSpeed  Stores current scene speed
var InstanceSpeed = Speed{0.0, 0.0, 1.0, ConfigSpeedFramerateDefault}

func (speed *Speed) setFramerate(framerate int) {
  speed.Framerate = float64(framerate)
}

func getSpeed() (*Speed) {
  return &InstanceSpeed
}

func updateSpeedFactor(factorOffset float64) {
  getSpeed().Factor += factorOffset * ConfigObjectFactorSpeedChangeFactor

  // Cap down to zero? (prevents negative or very-high speeds)
  if getSpeed().Factor < 0 {
    getSpeed().Factor = 0
  } else if getSpeed().Factor > ConfigObjectFactorSpeedMaximum {
    getSpeed().Factor = ConfigObjectFactorSpeedMaximum
  }
}

func updateElaspedTime(nowTime float64) {
  speed := getSpeed()

  speed.TimeElapsed = nowTime - speed.TimePrevious
  speed.TimePrevious = nowTime
}

func shouldUpdateScene(currentTime float64) (bool) {
  speed := getSpeed()

  return currentTime - speed.TimePrevious >= 1.0 / speed.Framerate
}

func angleSince(angleTime float32, factor float64, elapsed float64) float32 {
  // angleTime in milliseconds
  // elapsed in milliseconds
  //  -> angle = (elapsed / angleTime) * ConfigObjectFullAngle
  // Important: cap angle value (circle from 0 to 360 w/ modulus)

  if angleTime == 0 {
    return 0.0
  }

  return float32(math.Mod(((ConfigObjectFactorSpeedScene * factor * elapsed) / float64(angleTime)) * ConfigObjectFullAngle, ConfigObjectFullAngle))
}

func revolutionAngleSince(object *Object, factor float64, elapsed float64) float32 {
  // revolution_time from years to milliseconds

  return angleSince((*object).Revolution * float32(ConfigTimeYearToMilliseconds), factor, elapsed * float64(ConfigTimeSecondToMilliseconds))
}

func rotationAngleSince(object *Object, factor float64, elapsed float64) float32 {
  // revolution_time from years to milliseconds

  return angleSince((*object).Rotation * float32(ConfigTimeDayToMilliseconds), factor, elapsed * float64(ConfigTimeSecondToMilliseconds))
}

func revolutionAngleSinceLast(object *Object) float32 {
  return revolutionAngleSince(object, getSpeed().Factor, getSpeed().TimeElapsed)
}

func rotationAngleSinceLast(object *Object) float32 {
  return rotationAngleSince(object, getSpeed().Factor, getSpeed().TimeElapsed)
}

func revolutionAngleSinceStart(object *Object) float32 {
  return revolutionAngleSince(object, 1.0, float64(ConfigTimeStartFromMilliseconds))
}

func rotationAngleSinceStart(object *Object) float32 {
  return rotationAngleSince(object, 1.0, float64(ConfigTimeStartFromMilliseconds))
}

func normalizedTimeFactor() float32 {
  f := ConfigTimeNormalizeFactor * float32(getSpeed().TimeElapsed)

  return f
}
