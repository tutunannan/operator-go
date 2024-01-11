//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alhpa1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationClass) DeepCopyInto(out *AuthenticationClass) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationClass.
func (in *AuthenticationClass) DeepCopy() *AuthenticationClass {
	if in == nil {
		return nil
	}
	out := new(AuthenticationClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationProvider) DeepCopyInto(out *AuthenticationProvider) {
	*out = *in
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = new(OIDCProvider)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationProvider.
func (in *AuthenticationProvider) DeepCopy() *AuthenticationProvider {
	if in == nil {
		return nil
	}
	out := new(AuthenticationProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Database) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseConnection) DeepCopyInto(out *DatabaseConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseConnection.
func (in *DatabaseConnection) DeepCopy() *DatabaseConnection {
	if in == nil {
		return nil
	}
	out := new(DatabaseConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseConnectionCredentialSpec) DeepCopyInto(out *DatabaseConnectionCredentialSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseConnectionCredentialSpec.
func (in *DatabaseConnectionCredentialSpec) DeepCopy() *DatabaseConnectionCredentialSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseConnectionCredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseConnectionList) DeepCopyInto(out *DatabaseConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseConnectionList.
func (in *DatabaseConnectionList) DeepCopy() *DatabaseConnectionList {
	if in == nil {
		return nil
	}
	out := new(DatabaseConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseConnectionSpec) DeepCopyInto(out *DatabaseConnectionSpec) {
	*out = *in
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(DatabaseProvider)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseConnectionSpec.
func (in *DatabaseConnectionSpec) DeepCopy() *DatabaseConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseCredentialSpec) DeepCopyInto(out *DatabaseCredentialSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseCredentialSpec.
func (in *DatabaseCredentialSpec) DeepCopy() *DatabaseCredentialSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseCredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseList) DeepCopyInto(out *DatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Database, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseList.
func (in *DatabaseList) DeepCopy() *DatabaseList {
	if in == nil {
		return nil
	}
	out := new(DatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseProvider) DeepCopyInto(out *DatabaseProvider) {
	*out = *in
	if in.Mysq != nil {
		in, out := &in.Mysq, &out.Mysq
		*out = new(MysqlProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Postgres != nil {
		in, out := &in.Postgres, &out.Postgres
		*out = new(PostgresProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Redis != nil {
		in, out := &in.Redis, &out.Redis
		*out = new(RedisProvider)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseProvider.
func (in *DatabaseProvider) DeepCopy() *DatabaseProvider {
	if in == nil {
		return nil
	}
	out := new(DatabaseProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(DatabaseCredentialSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlProvider) DeepCopyInto(out *MysqlProvider) {
	*out = *in
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(DatabaseConnectionCredentialSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlProvider.
func (in *MysqlProvider) DeepCopy() *MysqlProvider {
	if in == nil {
		return nil
	}
	out := new(MysqlProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCProvider) DeepCopyInto(out *OIDCProvider) {
	*out = *in
	if in.ResponseType != nil {
		in, out := &in.ResponseType, &out.ResponseType
		*out = make([]ResponseType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCProvider.
func (in *OIDCProvider) DeepCopy() *OIDCProvider {
	if in == nil {
		return nil
	}
	out := new(OIDCProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresProvider) DeepCopyInto(out *PostgresProvider) {
	*out = *in
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(DatabaseConnectionCredentialSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresProvider.
func (in *PostgresProvider) DeepCopy() *PostgresProvider {
	if in == nil {
		return nil
	}
	out := new(PostgresProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisProvider) DeepCopyInto(out *RedisProvider) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisProvider.
func (in *RedisProvider) DeepCopy() *RedisProvider {
	if in == nil {
		return nil
	}
	out := new(RedisProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bucket) DeepCopyInto(out *S3Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bucket.
func (in *S3Bucket) DeepCopy() *S3Bucket {
	if in == nil {
		return nil
	}
	out := new(S3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketCredential) DeepCopyInto(out *S3BucketCredential) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketCredential.
func (in *S3BucketCredential) DeepCopy() *S3BucketCredential {
	if in == nil {
		return nil
	}
	out := new(S3BucketCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketList) DeepCopyInto(out *S3BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketList.
func (in *S3BucketList) DeepCopy() *S3BucketList {
	if in == nil {
		return nil
	}
	out := new(S3BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketSpec) DeepCopyInto(out *S3BucketSpec) {
	*out = *in
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(S3BucketCredential)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketSpec.
func (in *S3BucketSpec) DeepCopy() *S3BucketSpec {
	if in == nil {
		return nil
	}
	out := new(S3BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Connection) DeepCopyInto(out *S3Connection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Connection.
func (in *S3Connection) DeepCopy() *S3Connection {
	if in == nil {
		return nil
	}
	out := new(S3Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3Connection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConnectionList) DeepCopyInto(out *S3ConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3Connection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConnectionList.
func (in *S3ConnectionList) DeepCopy() *S3ConnectionList {
	if in == nil {
		return nil
	}
	out := new(S3ConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3ConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ConnectionSpec) DeepCopyInto(out *S3ConnectionSpec) {
	*out = *in
	if in.S3Credential != nil {
		in, out := &in.S3Credential, &out.S3Credential
		*out = new(S3Credential)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ConnectionSpec.
func (in *S3ConnectionSpec) DeepCopy() *S3ConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(S3ConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Credential) DeepCopyInto(out *S3Credential) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Credential.
func (in *S3Credential) DeepCopy() *S3Credential {
	if in == nil {
		return nil
	}
	out := new(S3Credential)
	in.DeepCopyInto(out)
	return out
}