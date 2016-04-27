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

  PositionEye    mgl32.Vec3
  PositionTarget mgl32.Vec3

  InertiaDrag    float64
  InertiaTurn    float64
}

var __CAMERA CameraData

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

func (camera_data *CameraData) moveTargetX(increment float32) {
  camera_data.PositionTarget[0] += increment
}

func (camera_data *CameraData) moveTargetY(increment float32) {
  camera_data.PositionTarget[1] += increment
}

func (camera_data *CameraData) moveTargetZ(increment float32) {
  camera_data.PositionTarget[2] += increment
}

func (camera_data *CameraData) defaultEye() {
  camera_data.PositionEye = CAMERA_DEFAULT_EYE
}

func (camera_data *CameraData) defaultTarget() {
  camera_data.PositionTarget = CAMERA_DEFAULT_TARGET
}

func (camera_data *CameraData) defaultInertia() {
  camera_data.InertiaDrag = 0.0
  camera_data.InertiaTurn = 0.0
}

func getCamera() (*CameraData) {
  return &__CAMERA
}

func createCamera(program uint32) {
  camera := getCamera()

  camera.CameraUniform = gl.GetUniformLocation(program, gl.Str("cameraUniform\x00"))

  // Default inertia (none)
  camera.defaultInertia()

  // Default camera position
  camera.defaultEye()
  camera.defaultTarget()
}

func produceInertia(inertia *float64, increment float64, celerity float64) {
  *inertia += increment * celerity

  // Cap inertia to maximum value
  if *inertia > celerity {
    *inertia = celerity
  } else if *inertia < -1.0 * celerity {
    *inertia = -1.0 * celerity
  }
}

func consumeInertia(inertia *float64) (float64) {
  if *inertia > 0 {
    *inertia += CAMERA_INERTIA_CONSUME_FORWARD
  } else if *inertia < 0 {
    *inertia += CAMERA_INERTIA_CONSUME_BACKWARD
  }

  return *inertia
}

func processEventCameraEye() {
  var (
    celerity     float64
    rotation_x   float64
    rotation_y   float64

    inertia_drag float64
    inertia_turn float64
  )

  camera := getCamera()
  key_state := getEventKeyState()

  // Decrease speed if diagonal move
  if key_state.MoveTurbo == true {
    celerity = CAMERA_MOVE_CELERITY_TURBO
  } else {
    celerity = CAMERA_MOVE_CELERITY_CRUISE
  }

  if (key_state.MoveUp == true || key_state.MoveDown == true) && (key_state.MoveLeft == true || key_state.MoveRight == true) {
    celerity /= math.Sqrt(2.0)
  }

  // Acquire rotation around axis
  rotation_x = float64(camera.getTargetX())
  rotation_y = float64(camera.getTargetY())

  // Process camera move position (keyboard)
  if key_state.MoveUp == true {
    produceInertia(&(camera.InertiaDrag), CAMERA_INERTIA_PRODUCE_FORWARD, celerity)
  }
  if key_state.MoveDown == true {
    produceInertia(&(camera.InertiaDrag), CAMERA_INERTIA_PRODUCE_BACKWARD, celerity)
  }
  if key_state.MoveLeft == true {
    produceInertia(&(camera.InertiaTurn), CAMERA_INERTIA_PRODUCE_FORWARD, celerity)
  }
  if key_state.MoveRight == true {
    produceInertia(&(camera.InertiaTurn), CAMERA_INERTIA_PRODUCE_BACKWARD, celerity)
  }

  // Apply new position with inertia
  inertia_drag = consumeInertia(&(camera.InertiaDrag))
  inertia_turn = consumeInertia(&(camera.InertiaTurn))

  camera.moveEyeX(float32(inertia_drag * -1.0 * math.Sin(rotation_y) + inertia_turn * math.Cos(rotation_y)))
  camera.moveEyeZ(float32(inertia_drag * math.Cos(rotation_y) + inertia_turn * math.Sin(rotation_y)))
  camera.moveEyeY(float32(inertia_drag * math.Sin(rotation_x)))

  // Translation: walk
  camera.Camera = camera.Camera.Mul4(mgl32.Translate3D(camera.getEyeX(), camera.getEyeY(), camera.getEyeZ()))
}

func processEventCameraTarget() {
  camera := getCamera()
  key_state := getEventKeyState()

  camera.moveTargetX(key_state.WatchY * float32(math.Pi) * CAMERA_TARGET_AMORTIZE_FACTOR)
  camera.moveTargetY(key_state.WatchX * float32(math.Pi) * 2 * CAMERA_TARGET_AMORTIZE_FACTOR)

  // Rotation: view
  camera.Camera = camera.Camera.Mul4(mgl32.HomogRotate3D(camera.getTargetX(), mgl32.Vec3{1, 0, 0}))
  camera.Camera = camera.Camera.Mul4(mgl32.HomogRotate3D(camera.getTargetY(), mgl32.Vec3{0, 1, 0}))
  camera.Camera = camera.Camera.Mul4(mgl32.HomogRotate3D(camera.getTargetZ(), mgl32.Vec3{0, 0, 1}))
}

func updateCamera() {
  // Update overall camera position
  getCamera().Camera = mgl32.Ident4()

  // Process new eye watch target in scene
  processEventCameraTarget()

  // Process new eye position in scene
  processEventCameraEye()
}

func resetCamera() {
  camera := getCamera()

  // Reset camera modifiers
  resetMouseCursor()

  // Reset camera itself
  camera.defaultInertia()

  camera.defaultEye()
  camera.defaultTarget()
}

func bindCamera() {
  camera := getCamera()

  gl.UniformMatrix4fv(camera.CameraUniform, 1, false, &(camera.Camera[0]))
}
