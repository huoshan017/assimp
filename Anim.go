package assimp

//#cgo pkg-config: assimp
//#include <assimp/anim.h>
import "C"

import (
	"reflect"
	"unsafe"
)

type VectorKey C.struct_aiVectorKey

func (v VectorKey) Time() float64 {
	return float64(v.mTime)
}

func (v VectorKey) Value() Vector3 {
	return Vector3(v.mValue)
}

type QuatKey C.struct_aiQuatKey

func (v QuatKey) Time() float64 {
	return float64(v.mTime)
}

func (v QuatKey) Value() Quaternion {
	return Quaternion(v.mValue)
}

type MeshKey C.struct_aiMeshKey

func (v MeshKey) Time() float64 {
	return float64(v.mTime)
}

func (v MeshKey) Value() int {
	return int(v.mValue)
}

type AnimBehaviour C.enum_aiAnimBehaviour

const (
	AnimBehaviour_Default  AnimBehaviour = C.aiAnimBehaviour_DEFAULT
	AnimBehaviour_Constant AnimBehaviour = C.aiAnimBehaviour_CONSTANT
	AnimBehaviour_Linear   AnimBehaviour = C.aiAnimBehaviour_LINEAR
	AnimBehaviour_Repeat   AnimBehaviour = C.aiAnimBehaviour_REPEAT
)

type NodeAnim C.struct_aiNodeAnim

func (n *NodeAnim) Name() string {
	return C.GoString(&n.mNodeName.data[0])
}

func (n *NodeAnim) NumPositionKeys() int {
	return int(n.mNumPositionKeys)
}

func (n *NodeAnim) PositionKeys() []VectorKey {
	if n.mNumPositionKeys > 0 && n.mPositionKeys != nil {
		var result []VectorKey
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(n.mNumPositionKeys)
		header.Len = int(n.mNumPositionKeys)
		header.Data = uintptr(unsafe.Pointer(n.mPositionKeys))
		return result
	} else {
		return nil
	}
}

func (n *NodeAnim) NumRotationKeys() int {
	return int(n.mNumRotationKeys)
}

func (n *NodeAnim) RotationKeys() []QuatKey {
	if n.mNumRotationKeys > 0 && n.mRotationKeys != nil {
		var result []QuatKey
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(n.mNumRotationKeys)
		header.Len = int(n.mNumRotationKeys)
		header.Data = uintptr(unsafe.Pointer(n.mRotationKeys))
		return result
	} else {
		return nil
	}
}

func (n *NodeAnim) NumScalingKeys() int {
	return int(n.mNumScalingKeys)
}

func (n *NodeAnim) ScalingKeys() []VectorKey {
	if n.mNumScalingKeys > 0 && n.mScalingKeys != nil {
		var result []VectorKey
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(n.mNumScalingKeys)
		header.Len = int(n.mNumScalingKeys)
		header.Data = uintptr(unsafe.Pointer(n.mScalingKeys))
		return result
	} else {
		return nil
	}
}

func (n *NodeAnim) PreState() AnimBehaviour {
	return AnimBehaviour(n.mPreState)
}

func (n *NodeAnim) PostState() AnimBehaviour {
	return AnimBehaviour(n.mPostState)
}

type MeshAnim C.struct_aiMeshAnim

func (n *MeshAnim) Name() string {
	return C.GoString(&n.mName.data[0])
}

func (n *MeshAnim) NumKeys() int {
	return int(n.mNumKeys)
}

func (n *MeshAnim) Keys() []MeshKey {
	if n.mNumKeys > 0 && n.mKeys != nil {
		var result []MeshKey
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(n.mNumKeys)
		header.Len = int(n.mNumKeys)
		header.Data = uintptr(unsafe.Pointer(n.mKeys))
		return result
	} else {
		return nil
	}
}

type Animation C.struct_aiAnimation

func (a *Animation) Name() string {
	return C.GoString(&a.mName.data[0])
}

func (a *Animation) Duration() float64 {
	return float64(a.mDuration)
}

func (a *Animation) TicksPerSecond() float64 {
	return float64(a.mTicksPerSecond)
}

func (a *Animation) NumChannels() int {
	return int(a.mNumChannels)
}

func (a *Animation) Channels() []*NodeAnim {
	if a.mNumChannels > 0 && a.mChannels != nil {
		var result []*NodeAnim
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(a.mNumChannels)
		header.Len = int(a.mNumChannels)
		header.Data = uintptr(unsafe.Pointer(a.mChannels))
		return result
	} else {
		return nil
	}
}

func (a *Animation) NumMeshChannels() int {
	return int(a.mNumMeshChannels)
}

func (a *Animation) MeshChannels() []*MeshAnim {
	if a.mNumMeshChannels > 0 && a.mMeshChannels != nil {
		var result []*MeshAnim
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(a.mNumMeshChannels)
		header.Len = int(a.mNumMeshChannels)
		header.Data = uintptr(unsafe.Pointer(a.mMeshChannels))
		return result
	} else {
		return nil
	}
}
