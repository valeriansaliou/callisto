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

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

type CameraData struct {
  Camera         mgl32.Mat4
  CameraUniform  int32

  PositionEye mgl32.Vec3
  PositionTarget mgl32.Vec3
}

var CAMERA CameraData

func (camera_data *CameraData) getEyeX() (position float32) {
  return camera_data.PositionEye[0]
}

func (camera_data *CameraData) getEyeY() (position float32) {
  return camera_data.PositionEye[1]
}

func (camera_data *CameraData) getEyeZ() (position float32) {
  return camera_data.PositionEye[2]
}

func (camera_data *CameraData) getTargetX() (position float32) {
  return camera_data.PositionTarget[0]
}

func (camera_data *CameraData) getTargetY() (position float32) {
  return camera_data.PositionTarget[1]
}

func (camera_data *CameraData) getTargetZ() (position float32) {
  return camera_data.PositionTarget[2]
}

func (camera_data *CameraData) moveEyeX(increment float32) {
  camera_data.PositionEye[0] += increment
}

func (camera_data *CameraData) moveEyeY(increment float32) {
  camera_data.PositionEye[1] += increment
}

func (camera_data *CameraData) moveEyeZ(increment float32) {
  camera_data.PositionEye[2] += increment
}

func (camera_data *CameraData) updateTargetX(position float32) {
  camera_data.PositionTarget[0] = position
}

func (camera_data *CameraData) updateTargetY(position float32) {
  camera_data.PositionTarget[1] = position
}

func (camera_data *CameraData) updateTargetZ(position float32) {
  camera_data.PositionTarget[2] = position
}

func getCamera() (*CameraData) {
  return &CAMERA
}

func createCamera(program uint32) {
  CAMERA.CameraUniform = gl.GetUniformLocation(program, gl.Str("cameraUniform\x00"))

  // Default camera position
  CAMERA.PositionEye = CAMERA_DEFAULT_EYE
  CAMERA.PositionTarget = CAMERA_DEFAULT_TARGET
}

func updateCamera() {
  key_state := getEventKeyState()

  // Decrease speed if diagonal move
  speed := CAMERA_MOVE_CELERITY

  if (key_state.MoveUp == true || key_state.MoveDown == true) && (key_state.MoveLeft == true || key_state.MoveRight == true) {
    speed /= math.Sqrt(2.0)
  }

  // Process camera move position (keyboard)
  if key_state.MoveUp == true {
    getCamera().moveEyeZ(float32(1.0 * speed))
  }
  if key_state.MoveDown == true {
    getCamera().moveEyeZ(float32(-1.0 * speed))
  }
  if key_state.MoveLeft == true {
    getCamera().moveEyeX(float32(1.0 * speed))
  }
  if key_state.MoveRight == true {
    getCamera().moveEyeX(float32(-1.0 * speed))
  }

  // Update overall camera position
  CAMERA.Camera = mgl32.Mat4{1.0, 0.0, 0.0, 0.0, 0.0, -1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0}

  CAMERA.Camera = CAMERA.Camera.Mul4(mgl32.Translate3D(getCamera().getEyeX(), getCamera().getEyeY(), getCamera().getEyeZ()))
}

func bindCamera() {
  gl.UniformMatrix4fv(CAMERA.CameraUniform, 1, false, &CAMERA.Camera[0])
}
