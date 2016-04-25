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
  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

type CameraData struct {
  Camera         mgl32.Mat4
  CameraUniform  int32

  PositionEye mgl32.Vec3
  PositionCenter mgl32.Vec3
  PositionUp mgl32.Vec3
}

var CAMERA CameraData

func (camera_data *CameraData) moveEyeX(position float32) {
  camera_data.PositionEye[0] = position
}

func (camera_data *CameraData) moveEyeY(position float32) {
  camera_data.PositionEye[1] = position
}

func (camera_data *CameraData) moveEyeZ(position float32) {
  camera_data.PositionEye[2] = position
}

func (camera_data *CameraData) moveCenterX(increment float32) {
  camera_data.PositionCenter[0] += increment
}

func (camera_data *CameraData) moveCenterY(increment float32) {
  camera_data.PositionCenter[1] += increment
}

func (camera_data *CameraData) moveCenterZ(increment float32) {
  camera_data.PositionCenter[2] += increment
}

func (camera_data *CameraData) moveUpX(increment float32) {
  camera_data.PositionUp[0] += increment
}

func (camera_data *CameraData) moveUpY(increment float32) {
  camera_data.PositionUp[1] += increment
}

func (camera_data *CameraData) moveUpZ(increment float32) {
  camera_data.PositionUp[2] += increment
}

func getCamera() (*CameraData) {
  return &CAMERA
}

func createCamera(program uint32) {
  CAMERA.CameraUniform = gl.GetUniformLocation(program, gl.Str("cameraUniform\x00"))

  // Default camera position
  CAMERA.PositionEye = CAMERA_DEFAULT_EYE
  CAMERA.PositionCenter = CAMERA_DEFAULT_CENTER
  CAMERA.PositionUp = CAMERA_DEFAULT_UP
}

func updateCamera() {
  key_state := getEventKeyState()

  // Camera position: Move
  if key_state.MoveUp == true {
    getCamera().moveCenterY(CAMERA_MOVE_CELERITY_FORWARD)
  }
  if key_state.MoveDown == true {
    getCamera().moveCenterY(CAMERA_MOVE_CELERITY_BACKWARD)
  }
  if key_state.MoveLeft == true {
    getCamera().moveCenterX(CAMERA_MOVE_CELERITY_BACKWARD)
  }
  if key_state.MoveRight == true {
    getCamera().moveCenterX(CAMERA_MOVE_CELERITY_FORWARD)
  }

  // Camera position: Watch
  getCamera().moveEyeX(key_state.WatchX)
  getCamera().moveEyeY(key_state.WatchY)

  // Update overall camera position
  CAMERA.Camera = mgl32.LookAtV(CAMERA.PositionEye, CAMERA.PositionCenter, CAMERA.PositionUp)
}

func bindCamera() {
  gl.UniformMatrix4fv(CAMERA.CameraUniform, 1, false, &CAMERA.Camera[0])
}
