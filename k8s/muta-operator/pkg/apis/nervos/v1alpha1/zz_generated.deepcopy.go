// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Benchmark) DeepCopyInto(out *Benchmark) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Benchmark.
func (in *Benchmark) DeepCopy() *Benchmark {
	if in == nil {
		return nil
	}
	out := new(Benchmark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	out.GraphQL = in.GraphQL
	in.Network.DeepCopyInto(&out.Network)
	out.Mempool = in.Mempool
	in.Consensus.DeepCopyInto(&out.Consensus)
	out.Executor = in.Executor
	out.Logger = in.Logger
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConsensus) DeepCopyInto(out *ConfigConsensus) {
	*out = *in
	if in.BlsPublicKeys != nil {
		in, out := &in.BlsPublicKeys, &out.BlsPublicKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConsensus.
func (in *ConfigConsensus) DeepCopy() *ConfigConsensus {
	if in == nil {
		return nil
	}
	out := new(ConfigConsensus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigExecutor) DeepCopyInto(out *ConfigExecutor) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigExecutor.
func (in *ConfigExecutor) DeepCopy() *ConfigExecutor {
	if in == nil {
		return nil
	}
	out := new(ConfigExecutor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigGenesis) DeepCopyInto(out *ConfigGenesis) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ConfigService, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigGenesis.
func (in *ConfigGenesis) DeepCopy() *ConfigGenesis {
	if in == nil {
		return nil
	}
	out := new(ConfigGenesis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigGraphQL) DeepCopyInto(out *ConfigGraphQL) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigGraphQL.
func (in *ConfigGraphQL) DeepCopy() *ConfigGraphQL {
	if in == nil {
		return nil
	}
	out := new(ConfigGraphQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigLogger) DeepCopyInto(out *ConfigLogger) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigLogger.
func (in *ConfigLogger) DeepCopy() *ConfigLogger {
	if in == nil {
		return nil
	}
	out := new(ConfigLogger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMempool) DeepCopyInto(out *ConfigMempool) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMempool.
func (in *ConfigMempool) DeepCopy() *ConfigMempool {
	if in == nil {
		return nil
	}
	out := new(ConfigMempool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMetadata) DeepCopyInto(out *ConfigMetadata) {
	*out = *in
	if in.VerifierList != nil {
		in, out := &in.VerifierList, &out.VerifierList
		*out = make([]ConfigVerifier, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMetadata.
func (in *ConfigMetadata) DeepCopy() *ConfigMetadata {
	if in == nil {
		return nil
	}
	out := new(ConfigMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigNetwork) DeepCopyInto(out *ConfigNetwork) {
	*out = *in
	if in.Bootstraps != nil {
		in, out := &in.Bootstraps, &out.Bootstraps
		*out = make([]ConfigNetworkBootstrap, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigNetwork.
func (in *ConfigNetwork) DeepCopy() *ConfigNetwork {
	if in == nil {
		return nil
	}
	out := new(ConfigNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigNetworkBootstrap) DeepCopyInto(out *ConfigNetworkBootstrap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigNetworkBootstrap.
func (in *ConfigNetworkBootstrap) DeepCopy() *ConfigNetworkBootstrap {
	if in == nil {
		return nil
	}
	out := new(ConfigNetworkBootstrap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigService) DeepCopyInto(out *ConfigService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigService.
func (in *ConfigService) DeepCopy() *ConfigService {
	if in == nil {
		return nil
	}
	out := new(ConfigService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigVerifier) DeepCopyInto(out *ConfigVerifier) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigVerifier.
func (in *ConfigVerifier) DeepCopy() *ConfigVerifier {
	if in == nil {
		return nil
	}
	out := new(ConfigVerifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyPair) DeepCopyInto(out *KeyPair) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyPair.
func (in *KeyPair) DeepCopy() *KeyPair {
	if in == nil {
		return nil
	}
	out := new(KeyPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Muta) DeepCopyInto(out *Muta) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Muta.
func (in *Muta) DeepCopy() *Muta {
	if in == nil {
		return nil
	}
	out := new(Muta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Muta) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutaList) DeepCopyInto(out *MutaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Muta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutaList.
func (in *MutaList) DeepCopy() *MutaList {
	if in == nil {
		return nil
	}
	out := new(MutaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MutaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutaSpec) DeepCopyInto(out *MutaSpec) {
	*out = *in
	if in.Chaos != nil {
		in, out := &in.Chaos, &out.Chaos
		*out = make([]ChaosType, len(*in))
		copy(*out, *in)
	}
	if in.Benchmark != nil {
		in, out := &in.Benchmark, &out.Benchmark
		*out = new(Benchmark)
		**out = **in
	}
	in.Config.DeepCopyInto(&out.Config)
	in.Genesis.DeepCopyInto(&out.Genesis)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutaSpec.
func (in *MutaSpec) DeepCopy() *MutaSpec {
	if in == nil {
		return nil
	}
	out := new(MutaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutaStatus) DeepCopyInto(out *MutaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutaStatus.
func (in *MutaStatus) DeepCopy() *MutaStatus {
	if in == nil {
		return nil
	}
	out := new(MutaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeCrypto) DeepCopyInto(out *NodeCrypto) {
	*out = *in
	if in.Keypairs != nil {
		in, out := &in.Keypairs, &out.Keypairs
		*out = make([]KeyPair, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeCrypto.
func (in *NodeCrypto) DeepCopy() *NodeCrypto {
	if in == nil {
		return nil
	}
	out := new(NodeCrypto)
	in.DeepCopyInto(out)
	return out
}