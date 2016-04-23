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

// import (
//   "github.com/go-gl/gl/v4.1-core/gl"
// )

// func createVoidBox() {
//   // Store the current matrix
//   gl.PushMatrix()

//   // Reset and transform the matrix.
//   gl.LoadIdentity()

//   // Enable/Disable features
//   gl.PushAttrib(gl.ENABLE_BIT)
//   gl.Enable(gl.TEXTURE_2D)
//   gl.Disable(gl.DEPTH_TEST)
//   gl.Disable(gl.LIGHTING)
//   gl.Disable(gl.BLEND)

//   // Just in case we set all vertices to white.
//   gl.Color4f(1,1,1,1)

//   // Render the front quad
//   gl.BindTexture(gl.TEXTURE_2D, voidbox_texture[0])
//   gl.Begin(gl.QUADS)
//       gl.TexCoord2f(0, 0)
//       gl.Vertex3f(  0.5, -0.5, -0.5 )
//       gl.TexCoord2f(1, 0)
//       gl.Vertex3f( -0.5, -0.5, -0.5 )
//       gl.TexCoord2f(1, 1)
//       gl.Vertex3f( -0.5,  0.5, -0.5 )
//       gl.TexCoord2f(0, 1)
//       gl.Vertex3f(  0.5,  0.5, -0.5 )
//   gl.End()

//   // Render the left quad
//   gl.BindTexture(gl.TEXTURE_2D, voidbox_texture[1])
//   gl.Begin(gl.QUADS)
//       gl.TexCoord2f(0, 0)
//       gl.Vertex3f(  0.5, -0.5,  0.5 )
//       gl.TexCoord2f(1, 0)
//       gl.Vertex3f(  0.5, -0.5, -0.5 )
//       gl.TexCoord2f(1, 1)
//       gl.Vertex3f(  0.5,  0.5, -0.5 )
//       gl.TexCoord2f(0, 1)
//       gl.Vertex3f(  0.5,  0.5,  0.5 )
//   gl.End()

//   // Render the back quad
//   gl.BindTexture(gl.TEXTURE_2D, voidbox_texture[2])
//   gl.Begin(gl.QUADS)
//       gl.TexCoord2f(0, 0)
//       gl.Vertex3f( -0.5, -0.5,  0.5 )
//       gl.TexCoord2f(1, 0)
//       gl.Vertex3f(  0.5, -0.5,  0.5 )
//       gl.TexCoord2f(1, 1)
//       gl.Vertex3f(  0.5,  0.5,  0.5 )
//       gl.TexCoord2f(0, 1)
//       gl.Vertex3f( -0.5,  0.5,  0.5 )

//   gl.End()

//   // Render the right quad
//   gl.BindTexture(gl.TEXTURE_2D, voidbox_texture[3])
//   gl.Begin(gl.QUADS)
//       gl.TexCoord2f(0, 0)
//       gl.Vertex3f( -0.5, -0.5, -0.5 )
//       gl.TexCoord2f(1, 0)
//       gl.Vertex3f( -0.5, -0.5,  0.5 )
//       gl.TexCoord2f(1, 1)
//       gl.Vertex3f( -0.5,  0.5,  0.5 )
//       gl.TexCoord2f(0, 1)
//       gl.Vertex3f( -0.5,  0.5, -0.5 )
//   gl.End()

//   // Render the top quad
//   gl.BindTexture(gl.TEXTURE_2D, voidbox_texture[4])
//   gl.Begin(gl.QUADS)
//       gl.TexCoord2f(0, 1)
//       gl.Vertex3f( -0.5,  0.5, -0.5 )
//       gl.TexCoord2f(0, 0)
//       gl.Vertex3f( -0.5,  0.5,  0.5 )
//       gl.TexCoord2f(1, 0)
//       gl.Vertex3f(  0.5,  0.5,  0.5 )
//       gl.TexCoord2f(1, 1)
//       gl.Vertex3f(  0.5,  0.5, -0.5 )
//   gl.End()

//   // Render the bottom quad
//   gl.BindTexture(gl.TEXTURE_2D, voidbox_texture[5])
//   gl.Begin(gl.QUADS)
//       gl.TexCoord2f(0, 0)
//       gl.Vertex3f( -0.5, -0.5, -0.5 )
//       gl.TexCoord2f(0, 1)
//       gl.Vertex3f( -0.5, -0.5,  0.5 )
//       gl.TexCoord2f(1, 1)
//       gl.Vertex3f(  0.5, -0.5,  0.5 )
//       gl.TexCoord2f(1, 0)
//       gl.Vertex3f(  0.5, -0.5, -0.5 )
//   gl.End()

//   // Restore enable bits and matrix
//   gl.PopAttrib()
//   gl.PopMatrix()
// }
