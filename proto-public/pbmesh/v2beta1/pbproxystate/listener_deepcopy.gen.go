// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package pbproxystate

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Listener within kubernetes types, where deepcopy-gen is used.
func (in *Listener) DeepCopyInto(out *Listener) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listener. Required by controller-gen.
func (in *Listener) DeepCopy() *Listener {
	if in == nil {
		return nil
	}
	out := new(Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Listener. Required by controller-gen.
func (in *Listener) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Router within kubernetes types, where deepcopy-gen is used.
func (in *Router) DeepCopyInto(out *Router) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Router. Required by controller-gen.
func (in *Router) DeepCopy() *Router {
	if in == nil {
		return nil
	}
	out := new(Router)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Router. Required by controller-gen.
func (in *Router) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Match within kubernetes types, where deepcopy-gen is used.
func (in *Match) DeepCopyInto(out *Match) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Match. Required by controller-gen.
func (in *Match) DeepCopy() *Match {
	if in == nil {
		return nil
	}
	out := new(Match)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Match. Required by controller-gen.
func (in *Match) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CidrRange within kubernetes types, where deepcopy-gen is used.
func (in *CidrRange) DeepCopyInto(out *CidrRange) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CidrRange. Required by controller-gen.
func (in *CidrRange) DeepCopy() *CidrRange {
	if in == nil {
		return nil
	}
	out := new(CidrRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CidrRange. Required by controller-gen.
func (in *CidrRange) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using L4Destination within kubernetes types, where deepcopy-gen is used.
func (in *L4Destination) DeepCopyInto(out *L4Destination) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4Destination. Required by controller-gen.
func (in *L4Destination) DeepCopy() *L4Destination {
	if in == nil {
		return nil
	}
	out := new(L4Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new L4Destination. Required by controller-gen.
func (in *L4Destination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using L7DestinationRoute within kubernetes types, where deepcopy-gen is used.
func (in *L7DestinationRoute) DeepCopyInto(out *L7DestinationRoute) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7DestinationRoute. Required by controller-gen.
func (in *L7DestinationRoute) DeepCopy() *L7DestinationRoute {
	if in == nil {
		return nil
	}
	out := new(L7DestinationRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new L7DestinationRoute. Required by controller-gen.
func (in *L7DestinationRoute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using L7Destination within kubernetes types, where deepcopy-gen is used.
func (in *L7Destination) DeepCopyInto(out *L7Destination) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7Destination. Required by controller-gen.
func (in *L7Destination) DeepCopy() *L7Destination {
	if in == nil {
		return nil
	}
	out := new(L7Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new L7Destination. Required by controller-gen.
func (in *L7Destination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SNIDestination within kubernetes types, where deepcopy-gen is used.
func (in *SNIDestination) DeepCopyInto(out *SNIDestination) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNIDestination. Required by controller-gen.
func (in *SNIDestination) DeepCopy() *SNIDestination {
	if in == nil {
		return nil
	}
	out := new(SNIDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SNIDestination. Required by controller-gen.
func (in *SNIDestination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
