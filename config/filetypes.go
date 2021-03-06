package config

import (
	"codexray/cxdig/types"
	"encoding/json"
)

// GetDefaultFileTypes returns the default definition list used to identify the file types
func GetDefaultFileTypes() []types.FileTypeInfo {
	var result []types.FileTypeInfo

	err := json.Unmarshal([]byte(defaultFileTypes), &result)
	if err != nil {
		panic(err) // not supposed to happen
	}

	return result
}

// embedded JSON definition list for file types
const defaultFileTypes = `
[
    {
        "language":"conan",
        "type":"BuildSystem",
        "fileNames":["conanfile.txt", "conanfile.py", "conanenv.txt"]
    },
    {
        "language":"scons",
        "type":"BuildSystem",
        "fileNames":["sconstruct"]
    },
    {
        "language":"premake",
        "type":"BuildSystem",
        "fileNames":["premake4.lua", "premake5.lua"]
    },
    {
        "language":"gulp",
        "type":"BuildSystem",
        "fileNames":["gulp.js"]
    },
    {
        "language":"zeus",
        "type":"BuildSystem",
        "fileNames":["zeusfile.yml"]
    },
    {
        "language":"bam",
        "type":"BuildSystem",
        "fileNames":["bam.lua"]
     },
     {
        "language":"meson",
        "type":"BuildSystem",
        "fileNames":["meson.build"]
    },
    {
        "language":"hunter",
        "type":"BuildSystem",
        "fileNames":["huntergate.cmake", "hunter.cmake"]
    },
    {
        "language":"cget",
        "type":"BuildSystem",
        "fileNames":["requirements.txt"]
    },
    {
        "language":"conda",
        "type":"BuildSystem",
        "fileNames":["meta.yaml"]
    },
    {
        "language":"shake",
        "type":"BuildSystem",
        "fileNames":["build.hs"]
    },
    {
        "language":"gemfile",
        "type":"BuildSystem",
        "fileNames":["gemfile"]
    },
    {
        "language":"npm",
        "type":"BuildSystem",
        "fileNames":["package.json"]
    },
    {
        "language":"webpack",
        "type":"BuildSystem",
        "fileNames":["webpack.config.js"]
    },
    {
        "language":"bower",
        "type":"BuildSystem",
        "fileNames":["bower.json"]
    },
    {
        "language":"maven",
        "type":"BuildSystem",
        "fileNames":["pom.xml"]
    },
    {
        "language":"cmake",
        "type":"BuildSystem",
        "fileNames":["cmakelists.txt"],
        "fileSuffixes": [".cmake"]
    },
    {
        "language":"makefile",
        "type":"BuildSystem",
        "fileNames":["makefile"],
        "fileSuffixes": [".make", ".mkfile", ".mak", ".mk"]
    },
    {
        "language":"qmake",
        "type":"BuildSystem",
        "fileSuffixes": [".pro", ".pri"]
    },
    {
        "language":"visual studio",
        "type":"BuildSystem",
        "fileSuffixes": [".sln", ".vcxproj", ".vcproj", ".props"]
    },
    {
        "language":"xcode",
        "type":"BuildSystem",
        "fileSuffixes": [".xcconfig", ".pbxproj", ".xcworkspacedata"]
    },
    {
        "language":"automake",
        "type":"BuildSystem",
        "fileSuffixes": [".am"]
    },
    {
        "language":"qbs",
        "type":"BuildSystem",
        "fileSuffixes": [".qbs"]
    },
    {
        "language":"ninja",
        "type":"BuildSystem",
        "fileSuffixes": [".ninja"]
    },
    {
        "language":"vcpkg",
        "type":"BuildSystem",
        "fileSuffixes": [".vcpkg"]
    },
    {
        "language":"boost.jam",
        "type":"BuildSystem",
        "fileSuffixes": [".jam"]
    },
    {
        "language":"gradle",
        "type":"BuildSystem",
        "fileSuffixes": [".gradle"]
    },
    {
        "language":"tup",
        "type":"BuildSystem",
        "fileSuffixes": [".tup"]
    },
    {
        "language":"bazel",
        "type":"BuildSystem",
        "fileSuffixes": [".bzl"]
    },
    {
        "language":"gyp",
        "type":"BuildSystem",
        "fileSuffixes": [".gyp", "gypi"]
    },
    {
        "language":"inno setup",
        "type":"BuildSystem",
        "fileSuffixes": [".iss"]
    },
    {
        "language":"wix",
        "type":"BuildSystem",
        "fileSuffixes": [".wixproj", ".wxs"]
    },
    {
        "language":"nsis",
        "type":"BuildSystem",
        "fileSuffixes": [".nsh", ".nsi"]
    },
    {
        "language":"eslint",
        "type":"EnvConfig",
        "filePrefixes": [".eslintrc."]
    },
    {
        "language":"travis",
        "type":"EnvConfig",
        "fileNames":[".travis.yml"]
    },
    {
        "language":"appveyor",
        "type":"EnvConfig",
        "fileNames":["appveyor.yml"]
    },
    {
        "language":"gitlab",
        "type":"EnvConfig",
        "fileNames":[".gitlab-ci.yml"]
    },
    {
        "language":"circleci",
        "type":"EnvConfig",
        "fileNames":["circle.yml"]
    },
    {
        "language":"clangformat",
        "type":"EnvConfig",
        "fileNames":[".clang-format"]
    },
    {
        "language":"clang_complete",
        "type":"EnvConfig",
        "fileNames":[".clang_complete"]
    },
    {
        "language":"editorconfig",
        "type":"EnvConfig",
        "fileNames":[".editorconfig"]
    },
    {
        "language":"gdbinit",
        "type":"EnvConfig",
        "fileNames":[".gdbinit"]
    },
    {
        "language":"yard",
        "type":"EnvConfig",
        "fileNames":[".yardopts"]
    },
    {
        "language":"istanbul",
        "type":"EnvConfig",
        "fileNames":[".istanbul.yml"]
    },
    {
        "language":"codecov.io",
        "type":"EnvConfig",
        "fileNames":[".codecov.yml"]
    },
    {
        "language":"pylint",
        "type":"EnvConfig",
        "fileNames":[".pylintrc"]
    },
    {
        "language":"flake8",
        "type":"EnvConfig",
        "fileNames":[".flake8"]
    },
    {
        "language":"emacs.dir-locals",
        "type":"EnvConfig",
        "fileNames":[".dir-locals.el"]
    },
    {
        "language":"doxygen",
        "type":"EnvConfig",
        "fileNames":["doxygen.config"]
    },
    {
        "language":"mention-bot",
        "type":"EnvConfig",
        "fileNames":[".mention-bot"]
    },
    {
        "language":"apache-2.0",
        "type":"License",
        "fileNames":["apache-2.0.txt"]
    },
    {
        "language":"agpl-3.0",
        "type":"License",
        "fileNames":["gnu-agpl-3.0.txt"]
    },
  
    {
        "language":"assembly",
        "type":"Source",
        "fileSuffixes": [".asm", ".s"]
    },
    {
        "language":"c",
        "type":"Source",
        "fileSuffixes": [".c"]
    },
    {
        "language":"c++",
        "type":"Source",
        "fileSuffixes": [".cpp", ".cc", "cxx", ".c++", ".h", ".hpp", ".hh", ".hxx", ".h++", ".inc", ".inl", ".ipp", ".tcc", ".tpp", ".txx", ".moc"]
    },
    {
        "language":"c#",
        "type":"Source",
        "fileSuffixes": [".cs", ".csx"]
    },
    {
        "language":"d",
        "type":"Source",
        "fileSuffixes": [".d", ".dd", ".di"]
    },
    {
        "language":"fortran",
        "type":"Source",
        "fileSuffixes": [".f", ".f03", ".f08", ".f77", ".f90", ".f95", ".for", ".fpp"]
    },
    {
        "language":"go",
        "type":"Source",
        "fileSuffixes": [".go"]
    },
    {
        "language":"java",
        "type":"Source",
        "fileSuffixes": [".java"]
    },
    {
        "language":"pascal",
        "type":"Source",
        "fileSuffixes": [".dfm", ".dpr", ".lpr", ".pas", ".pascal"]
    },
    {
        "language":"perl",
        "type":"Source",
        "fileSuffixes": [".al", ".perl", ".ph", ".pl", ".plx", ".pm", ".psgi"]
    },
    {
        "language":"php",
        "type":"Source",
        "fileSuffixes": [".aw", ".ctp", ".php", ".php3", ".php4", ".php5", ".phps", ".phpt", ".phtml"]
    },
    {
        "language":"python",
        "type":"Source",
        "fileSuffixes": [".lmi", ".py", ".py3", ".pyde", ".pyp", ".pyt", ".pyw", ".xpy"]
    },
    {
        "language":"cython",
        "type":"Source",
        "fileSuffixes": [".pyx"]
    },
    {
        "language":"groovy",
        "type":"Source",
        "fileSuffixes": [".groovy", ".grt", ".gsp", ".gtpl", ".gvy"]
    },
    {
        "language":"qml",
        "type":"Source",
        "fileSuffixes": [".qml"]
    },
    {
        "language":"r",
        "type":"Source",
        "fileSuffixes": [".r", ".rd", ".rsx"]
    },
    {
        "language":"ruby",
        "type":"Source",
        "fileSuffixes": [".eye", ".gemspec", ".god", ".irbc", ".rabl", ".rake", ".rb", ".rbuild", ".rbw", ".rbx", ".ru", ".ruby"]
    },
    {
        "language":"rust",
        "type":"Source",
        "fileSuffixes": [".rs"]
    },
    {
        "language":"sql",
        "type":"Source",
        "fileSuffixes": [".cql", ".ddl", ".mysql", ".prc", ".sql", ".tab", ".udf", ".viw"]
    },
    {
        "language":"visual basic",
        "type":"Source",
        "fileSuffixes": [".vb", ".vba", ".vbs"]
    },
    {
        "language":"swig",
        "type":"Source",
        "fileSuffixes": [".i"]
    },
    {
        "language":"idl",
        "type":"Source",
        "fileSuffixes": [".idl"]
    },
    {
        "language":"protocol buffer",
        "type":"Source",
        "fileSuffixes": [".proto"]
    },
    {
        "language":"thrift",
        "type":"Source",
        "fileSuffixes": [".thrift"]
    },
    {
        "language":"flatbuffers",
        "type":"Generator",
        "fileSuffixes": [".fbs"]
    },
    {
        "language":"cap'n proto",
        "type":"Generator",
        "fileSuffixes": [".capnp"]
    },
    {
        "language":"lex",
        "type":"Generator",
        "fileSuffixes": [".l", ".lex", ".ll"]
    },
    {
        "language":"yacc",
        "type":"Generator",
        "fileSuffixes": [".y", ".yacc", ".yxx"]
    },
    {
        "language":"m4",
        "type":"Generator",
        "fileSuffixes": [".m4"]
    }
]
`
