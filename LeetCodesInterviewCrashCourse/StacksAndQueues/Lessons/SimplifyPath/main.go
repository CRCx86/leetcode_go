package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
	fmt.Println(simplifyPath("/..."))

	fmt.Println(simplifyPath2("/home/"))
	fmt.Println(simplifyPath2("/../"))
	fmt.Println(simplifyPath2("/home//foo/"))
	fmt.Println(simplifyPath2("/a/./b/../../c/"))
	fmt.Println(simplifyPath2("/a//b////c/d//././/.."))
	fmt.Println(simplifyPath2("/..."))
}

func simplifyPath(path string) string {

	stack := make([]string, 0)

	buff := make([]uint8, 0)
	for i := 1; i < len(path); i++ {
		s := string(buff)
		if path[i] == '/' && s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			buff = make([]uint8, 0)
			continue
		} else if path[i] == '/' && s == "." {
			buff = make([]uint8, 0)
			continue
		} else if path[i] == '/' && len(s) > 0 {
			stack = append(stack, s)
			buff = make([]uint8, 0)
			continue
		} else if path[i] != '/' {
			buff = append(buff, path[i])
		}
	}

	s := string(buff)
	if s == ".." {
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	} else if len(s) > 0 && s != "." && s != "/" {
		stack = append(stack, s)
	}

	return "/" + strings.Join(stack, "/")

}

func simplifyPath2(path string) string {

	subs := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, s := range subs {

		if s == "." {
			continue
		}

		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		if len(s) > 0 {
			stack = append(stack, s)
		}

	}

	if len(stack) > 0 {
		return "/" + strings.Join(stack, "/")
	} else {
		return "/"
	}
}
