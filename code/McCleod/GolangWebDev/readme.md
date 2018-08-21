# Golang Web Dev

## Links
- [Course Outline](https://docs.google.com/document/d/1QKWp1VYd26uiQZWIR05pahSa0HnbD1qqj9dtIQiVVjU/edit)
- [Course Description](https://docs.google.com/document/d/1e7r0H_3VmJw9wf3dRMWfp8_7CsLvDUpsLQ5_WacZjvw/edit)
- [Course Resources](https://docs.google.com/document/d/1jfU8-3qxrWWP9KVxrNPA77KDzRStE6bakqKUzFDJArQ/edit)

## Outline

- Templates
  - ~~primitive Method (string to html)~~
  - Parse & Execute
    - ParseFiles
    - ParseGlob
    - `init()` & `template.Must()`
  - Data 
    - Variables
  - Data Structures
    - slices
    - maps
    - structs
    - slice struct
    - struct slice struct
    - *anonymous types*
  - Functions
    - `template.FuncMap`
    - pipelines
    - predefined global functions
  - Nested Templates
  - Composition
- TCP
  - TCP (Transmission Control Protocol)
    - Listen
    - Accept
    - Read/Write
- HTTP
  - `net/http`
  - `Handler` - (`ServeHTTP(ResponseWriter, *Request)`)
  - Forms
    - POST -> body
    - GET -> URL
- Routing
  - `http.NewServeMux()`
  - Default ServeMux
  - `http.HandleFunc()`
  - `http.HandlerFunc()`
  - third-party-serveMux
- Serving Files
  - ~~io.Copy()~~
  - ServeContent
  - ServeFile
  - http.FileServer



## Terms

|   Term    |                                                                                                                                                             Definition                                                                                                                                                              |                         Link                         |
| --------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------- |
| OSI model | Open Systems Interconnection model                                                                                                                                                                                                                                                                                                  | [wikipedia](https://en.wikipedia.org/wiki/OSI_model) |
| tcp       | Transmission Control Protocol                                                                                                                                                                                                                                                                                                       |                                                      |
| http      | Hypertext Transfer Protocol                                                                                                                                                                                                                                                                                                         |                                                      |
| mux       | In electronics, a multiplexer is a device that selects one of several analog or digital input signals and forwards the selected input into a single line. <br> In telecommunications and computer networks, multiplexing is a method by which multiple analog or digital signals are combined into one signal over a shared medium. |                                                      |
| rfc7230  |                                                                                                                                                                                                                                                                                                                                     | [ietf.org](https://tools.ietf.org/html/rfc7230)      |
