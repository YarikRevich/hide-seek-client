
package main

import (
	"github.com/bazelbuild/bazel-gazelle/language"

	proto_ "github.com/bazelbuild/bazel-gazelle/language/proto"
	go_ "github.com/bazelbuild/bazel-gazelle/language/go"
)

var languages = []language.Language{
	proto_.NewLanguage(),
	go_.NewLanguage(),
}
