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

var SPEED_TIME_PREVIOUS float64 = 0.0
var SPEED_TIME_ELAPSED float64 = 0.0

func updateElaspedTime(nowTime float64) {
  SPEED_TIME_ELAPSED = nowTime - SPEED_TIME_PREVIOUS
  SPEED_TIME_PREVIOUS = nowTime
}

func revolutionAngleSinceLast(object *Object) float32 {
  // revolution_time from years to milliseconds
  // SPEED_TIME_ELAPSED in milliseconds
  //  -> angle = (SPEED_TIME_ELAPSED / revolution_time) * OBJECT_REVOLUTION_FULL_ANGLE

  return float32(OBJECT_FACTOR_SPEED_SCENE * SPEED_TIME_ELAPSED) / ((*object).Revolution * float32(TIME_YEAR_TO_MILLISECONDS)) * float32(OBJECT_REVOLUTION_FULL_ANGLE)
}

func rotationAngleSinceLast(object *Object) float32 {
  // rotation_time from days to milliseconds
  // SPEED_TIME_ELAPSED in milliseconds
  //  -> angle = (SPEED_TIME_ELAPSED / rotation_time) * OBJECT_ROTATION_FULL_ANGLE

  return float32(OBJECT_FACTOR_SPEED_SCENE * SPEED_TIME_ELAPSED) / ((*object).Rotation * float32(TIME_DAY_TO_MILLISECONDS)) * float32(OBJECT_ROTATION_FULL_ANGLE)
}
