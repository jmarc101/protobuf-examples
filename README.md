#+title: Protocol Buffer
* Protocol Buffers (Protobuf) Quick Reference

** `protoc` and Flags

*** Specify Output Directory

You can specify an output directory for the generated Go code:

#+BEGIN_SRC shell
protoc --go_out=output_dir path/to/yourfile.proto
#+END_SRC
Replace output_dir with the desired directory path.

*** Multiple .proto Files

You can generate code from multiple .proto files at once:

#+BEGIN_SRC shell
protoc --go_opt=module=example.com/myapp --go_out=. path/to/yourfile1.proto path/to/yourfile2.proto
#+END_SRC

This generates Go code for both yourfile1.proto and yourfile2.proto.

*** Using Subdirectories

You can organize your .proto files into subdirectories and generate code accordingly:

##+BEGIN_SRC shell
protoc --go_out=output_dir path/to/proto/*.proto
#+END_SRC

This processes all .proto files in the specified directory and its subdirectories, generating Go code in the output_dir.

+BEGIN_SRC shell
protoc --go_opt=module=example.com/myapp --go_out=. path/to/proto/*.proto
#+END_SRC

This will process .proto and generate the code in the option go_package specified in .proto file

*** Set Go Package Name

You can set the Go package name for the generated code:

#+BEGIN_SRC shell
protoc --go_opt=module=example.com/myapp --go_out=plugins=grpc:. path/to/yourfile.proto
#+END_SRC

In this example, we use the plugins=grpc option to enable gRPC code generation and set the Go package name.*** Common Command-Line Flags

- `-I` or `--proto_path`: Specifies the directory in which to search for `.proto` files. You can use this flag to specify the import directories.

- `--proto_path=path1,path2`: Allows specifying multiple import directories, separated by commas.

- `--go_out`: Specifies the output directory for Go code generation. For Go, you can use this flag with the `plugins=grpc` option to enable gRPC support.

- `--java_out`: Specifies the output directory for Java code generation.

- `--python_out`: Specifies the output directory for Python code generation.

- `--cpp_out`: Specifies the output directory for C++ code generation.

- `--csharp_out`: Specifies the output directory for C# code generation.

- `--swift_out`: Specifies the output directory for Swift code generation.

*** Import Directories

The `-I` or `--proto_path` flag is used to specify import directories for `.proto` files. Import directories are essential for resolving imports and including `.proto` files referenced by other files. You can specify one or more directories to search for `.proto` files.

In this example, we specify two import directories (`proto_dir` and `common_proto`) to search for `.proto` files.

*** `go_opt=module=<GoModule>` Option

In Protocol Buffers (Protobuf), the `go_opt=module=<GoModule>` option is used to specify the Go module name for the generated Go code. This option is particularly useful when you want to ensure that the generated Go code is organized within a specific Go module.

Here's how you can use this option:

#+BEGIN_SRC protobuf
option go_package = "<GoModule>/<GoPackage>";
#+END_SRC

- `<GoModule>`: Replace this with the name of your Go module where you want the generated code to be placed.

- `<GoPackage>`: Replace this with the desired Go package name within your module.

For example, let's say you have a Go module named "myapp" and you want the generated Go code to be placed in the "proto" package within that module. You would define it as follows:

#+BEGIN_SRC protobuf
option go_package = "myapp/proto";
#+END_SRC

By specifying the `go_opt=module` option, you ensure that the generated Go code is placed in the appropriate Go module and package, making it easier to manage your Protobuf-generated code within your Go project.

This option is particularly useful when you have multiple Protobuf files and want to organize the generated Go code into a structured project directory. It helps maintain a clear and organized codebase in your Go application.


** Scalar Types

*** Scalars and Zero values
| Protobuf Scalar | Description                                       | Go Type | Python Type     | Java Type  | Zero Value |
|-----------------+---------------------------------------------------+---------+-----------------+------------+------------|
| bool            | Boolean value (true or false)                     | bool    | bool            | boolean    |      false |
| string          | UTF-8 encoded string                              | string  | string          | String     | "" (empty) |
| bytes           | Arbitrary byte sequence (binary data)             | []byte  | bytes           | ByteString |        nil |
| double          | Double-precision floating-point number            | float64 | float or double | double     |        0.0 |
| float           | Single-precision floating-point number            | float32 | float or double | float      |        0.0 |
| int32           | 32-bit signed integer                             | int32   | int             | int        |          0 |
| int64           | 64-bit signed integer                             | int64   | int             | long       |          0 |
| uint32          | 32-bit unsigned integer                           | uint32  | int             | int        |          0 |
| uint64          | 64-bit unsigned integer                           | uint64  | int             | long       |          0 |
| sint32          | Signed 32-bit integer with variable encoding      | int32   | int             | int        |          0 |
| sint64          | Signed 64-bit integer with variable encoding      | int64   | int             | long       |          0 |
| fixed32         | 32-bit unsigned integer with fixed-width encoding | uint32  | int             | int        |          0 |
| fixed64         | 64-bit unsigned integer with fixed-width encoding | uint64  | int             | long       |          0 |
| sfixed32        | Signed 32-bit integer with fixed-width encoding   | int32   | int             | int        |          0 |
| sfixed64        | Signed 64-bit integer with fixed-width encoding   | int64   | int             | long       |          0 |

*** Optional Fields:

Optional fields may or may not be present in a message. They are often used for values that might not always be available.

#+begin_src protobuf
message Person {
    string name = 1;
    int32 age = 2;
}
#+end_src

In Go, optional fields are represented as pointers:

#+begin_src go
person := &Person{
    name: "Alice",
    // You can omit age to make it optional
}
#+end_src

*** Repeated Fields:

Repeated fields are used for lists or arrays of values and can have zero or more elements. They are ideal for storing multiple instances of a value in a single field.

#+begin_src
message ShoppingCart {
    repeated string items = 1;
}
#+end_src

In Go, repeated fields are represented as slices:
#+begin_src go
cart := &ShoppingCart{
    items: []string{"item1", "item2", "item3"},
    // You can have an empty slice for zero or more items
}
#+end_src


** Protobuf to Json in Go
Default go json encoder doesn't work with protobuf. Need to import to have encoder.
*** Import and availale functions

"google.golang.org/protobuf/encoding/protojson"

2 important function come from protojson
#+begin_src go
func Marshal(m proto.Message) ([]byte, error)
func Unmarshal(b []byte, m proto.Message) error
#+end_src

*** Sample use case
#+begin_src go
package basic

import (
    "google.golang.org/protobuf/encoding/protojson"
    "grpc-go/protogen/basic"
    "log"
)

func ProtoToJsonUser(){
    user := &basic.User{
        Id: 99,
        Username: "Cat women",
        IsActive: true,
        Password: []byte("654321"),
        Emails: []string{"Email@email.com"},
        Gender: basic.Gender_GENDER_FEMALE,
    }

    bytesProto, _ := protojson.Marshal(user)
    jsonUser := string(bytesProto)

    log.Println(jsonUser)
}


#+end_src
** Messages and Fields
   - Defining Messages
   - Message Fields Overview
   - Data Types (int, string, bool, bytes, etc.)
   - Field Options and Annotations

** Enums and Oneof
   - Enumerations in Protobuf
   - Oneof Fields (Unions)

** Nested Messages
   - Creating Nested Messages
   - Accessing Nested Message Fields

** Extensions
   - Extending Protobuf Messages
   - Custom Extensions

** Serialization
   - Serializing Protobuf Data
   - Deserializing Protobuf Data

** Advanced Topics
   - Packages and Namespaces
   - Custom Options and Extensions
   - Compatibility and Versioning

** gRPC and Protobuf
   - Protobuf in gRPC (brief introduction)

** Resources
   - Official Protobuf Documentation
   - Community and Forums
   - Tutorials and Guides
