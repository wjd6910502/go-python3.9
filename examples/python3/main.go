/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package main

import (
    "log"
 "fmt"
 "os"
    "github.com/DataDog/go-python3"
)

func main() {
  python3.Py_Initialize()
   if !python3.Py_IsInitialized() {
      fmt.Println("Error initializing the python interpreter")
      os.Exit(1)
   }
   sysModule := python3.PyImport_ImportModule("sys")
   path := sysModule.GetAttrString("path")
   //pathStr, _ := pythonRepr2(path)
   //log.Println("before add path is " + pathStr)
   python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(""))
   python3.PyList_Insert(path, 0, python3.PyUnicode_FromString("./py3"))

    // 初始化 Python 解释器
    //python3.Py_Initialize()
    defer python3.Py_Finalize()

    // 加载 Python 模块
    module := python3.PyImport_ImportModule("script")
    if module == nil {
        log.Fatal("Error: failed to import module")
    }
    defer module.DecRef()

    // 获取 Python 函数对象
    function := module.GetAttrString("add")
    if function == nil {
        log.Fatal("Error: failed to get function")
    }
    defer function.DecRef()

    // 调用 Python 函数
    args := python3.PyTuple_New(2)
    if args == nil {
        log.Fatal("Error: failed to create tuple")
    }
    defer args.DecRef()

    arg1 := python3.PyLong_FromLong(1)
    if arg1 == nil {
        log.Fatal("Error: failed to create argument 1")
    }
    defer arg1.DecRef()

    arg2 := python3.PyLong_FromLong(2)
    if arg2 == nil {
        log.Fatal("Error: failed to create argument 2")
    }
    defer arg2.DecRef()

    python3.PyTuple_SetItem(args, 0, arg1)
    python3.PyTuple_SetItem(args, 1, arg2)

    result := function.Call(args, python3.PyDict_New())
    if result == nil {
        log.Fatal("Error: failed to call function")
    }
    defer result.DecRef()

  log.Println("begin to compute ")
    // 处理 Python 函数返回值
    sum := python3.PyLong_AsLong(result)
    log.Println("The sum of 1 and 2 is", sum)
}

