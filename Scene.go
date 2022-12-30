package assimp

//#cgo pkg-config: assimp
//#include <assimp/scene.h>
import "C"

import (
	"reflect"
	"unsafe"
)

type Node C.struct_aiNode

func (n *Node) Name() string {
	return C.GoString(&n.mName.data[0])
}

func (n *Node) Transformation() Matrix4x4 {
	return Matrix4x4(n.mTransformation)
}

func (n *Node) Parent() *Node {
	return (*Node)(n.mParent)
}

func (n *Node) NumChildren() int {
	return int(n.mNumChildren)
}

func (n *Node) Children() []*Node {
	if n.mNumChildren > 0 && n.mChildren != nil {
		var result []*Node
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(n.mNumChildren)
		header.Len = int(n.mNumChildren)
		header.Data = uintptr(unsafe.Pointer(n.mChildren))
		return result
	} else {
		return nil
	}
}

func (n *Node) NumMeshes() int {
	return int(n.mNumMeshes)
}

func (n *Node) Meshes() []int32 {
	if n.mNumMeshes > 0 && n.mMeshes != nil {
		var result []int32
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(n.mNumMeshes)
		header.Len = int(n.mNumMeshes)
		header.Data = uintptr(unsafe.Pointer(n.mMeshes))
		return result
	} else {
		return nil
	}
}

const (
	SceneFlags_Incomplete        = C.AI_SCENE_FLAGS_INCOMPLETE
	SceneFlags_Validated         = C.AI_SCENE_FLAGS_VALIDATED
	SceneFlags_ValidationWarning = C.AI_SCENE_FLAGS_VALIDATION_WARNING
	SceneFlags_NonVerboseFormat  = C.AI_SCENE_FLAGS_NON_VERBOSE_FORMAT
	SceneFlags_Terrain           = C.AI_SCENE_FLAGS_TERRAIN
)

type Scene C.struct_aiScene

func (s *Scene) Flags() uint {
	return uint(s.mFlags)
}

func (s *Scene) RootNode() *Node {
	return (*Node)(s.mRootNode)
}

func (s *Scene) NumMeshes() int {
	return int(s.mNumMeshes)
}

func (s *Scene) Meshes() []*Mesh {
	if s.mNumMeshes > 0 && s.mMeshes != nil {
		var result []*Mesh
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(s.mNumMeshes)
		header.Len = int(s.mNumMeshes)
		header.Data = uintptr(unsafe.Pointer(s.mMeshes))
		return result
	} else {
		return nil
	}
}

func (s *Scene) NumMaterials() int {
	return int(s.mNumMaterials)
}

func (s *Scene) Materials() []*Material {
	if s.mNumMaterials > 0 && s.mMaterials != nil {
		var result []*Material
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(s.mNumMaterials)
		header.Len = int(s.mNumMaterials)
		header.Data = uintptr(unsafe.Pointer(s.mMaterials))
		return result
	} else {
		return nil
	}
}

func (s *Scene) NumAnimations() int {
	return int(s.mNumAnimations)
}

func (s *Scene) Animations() []*Animation {
	if s.mNumAnimations > 0 && s.mAnimations != nil {
		var result []*Animation
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(s.mNumAnimations)
		header.Len = int(s.mNumAnimations)
		header.Data = uintptr(unsafe.Pointer(s.mAnimations))
		return result
	} else {
		return nil
	}
}

func (s *Scene) NumTextures() int {
	return int(s.mNumTextures)
}

func (s *Scene) Textures() []*Texture {
	if s.mNumTextures > 0 && s.mTextures != nil {
		var result []*Texture
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(s.mNumTextures)
		header.Len = int(s.mNumTextures)
		header.Data = uintptr(unsafe.Pointer(s.mTextures))
		return result
	} else {
		return nil
	}
}

func (s *Scene) NumLights() int {
	return int(s.mNumLights)
}

func (s *Scene) Lights() []*Light {
	if s.mNumLights > 0 && s.mLights != nil {
		var result []*Light
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(s.mNumLights)
		header.Len = int(s.mNumLights)
		header.Data = uintptr(unsafe.Pointer(s.mLights))
		return result
	} else {
		return nil
	}
}

func (s *Scene) NumCameras() int {
	return int(s.mNumCameras)
}

func (s *Scene) Cameras() []*Camera {
	if s.mNumCameras > 0 && s.mCameras != nil {
		var result []*Camera
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(s.mNumCameras)
		header.Len = int(s.mNumCameras)
		header.Data = uintptr(unsafe.Pointer(s.mCameras))
		return result
	} else {
		return nil
	}
}
