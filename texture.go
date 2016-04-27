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
  "fmt"
  "image"
  _ "image/jpeg"
  _ "image/png"
  "image/draw"
  "os"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/disintegration/imaging"
)

// Texture  Maps a texture reference to GPU internals
type Texture struct {
  Ref uint32
}

func loadTexture(name string) (Texture, error) {
  var texture = Texture{}

  // Try JPG version of texture (most used)
  imgFile, err := os.Open(fmt.Sprintf("assets/%s.jpg", name))
  if err != nil {
    // JPG not found, try PNG version of texture (less used)
    imgFile, err = os.Open(fmt.Sprintf("assets/%s.png", name))
  }
  if err != nil {
    // Open default texture (default object color)
    imgFile, err = os.Open("assets/default.png")

    if err != nil {
      panic("Failed opening fallback texture file")
    }
  }

  img, _, err := image.Decode(imgFile)
  if err != nil {
    return texture, err
  }

  // Flip image to avoid reverse textures when bound on stellar objects
  img = imaging.FlipV(img)

  rgba := image.NewRGBA(img.Bounds())
  draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

  // Generate unique texture
  gl.GenTextures(1, &texture.Ref);
  gl.ActiveTexture(texture.Ref)
  gl.BindTexture(gl.TEXTURE_2D, texture.Ref);

  gl.TexImage2D(
    gl.TEXTURE_2D,
    0,
    gl.RGBA,
    int32(rgba.Rect.Size().X),
    int32(rgba.Rect.Size().Y),
    0,
    gl.RGBA,
    gl.UNSIGNED_BYTE,
    gl.Ptr(rgba.Pix))

  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST);
  gl.GenerateMipmap(gl.TEXTURE_2D);
  gl.BindTexture(gl.TEXTURE_2D, 0);

  return texture, nil
}
