# Seminar "Deadlock and data race prediction"
This repository contains the code and documentation of the seminar at the HKA.

## T1: Chronos and RELAY — Scaling Static Data Race Detection to Go

Dynamic race detectors (like Go's built-in -race flag) only find bugs if your test suites happen to trigger the exact racy execution path at runtime. Static analysis aims to solve this by scanning all possible execution paths at compile-time without ever running the code. The classic [RELAY](https://cseweb.ucsd.edu/~lerner/papers/relay.pdf) paper laid the foundational architecture for how to scale static race detection to millions of lines of code. [Chronos](https://github.com/amit-davidson/Chronos) is an open-source tool that attempts to implement these heavy static analysis principles specifically for the Go programming language.

### The Mission

Act as an independent tool reviewer. You will deploy the open-source Go static analyzer Chronos, build a test suite to discover its boundaries, and use the structural concepts from the RELAY paper to analyze the core challenges of static code verification.

### Your Step-by-Step Task List

1. <b>Deploy the Engine:</b> Clone the open-source repository for [Chronos - A static race detector for the go language](https://github.com/amit-davidson/Chronos). Follow its setup guide and run the analyzer against a simple multi-threaded Go file that uses basic mutex locking.

2. <b>Stress-Test the Tool:</b> Chronos relies on tracking guarded memory accesses to find races. Write 3 minimal Go test programs to locate the architectural limits of the tool:

*  1 program using sync.Mutex that contains a clear data race (verify if Chronos catches it).

* 1 program using Go chan (channels) or sync.WaitGroup that contains a data race. *Hint: Look at Chronos's documentation regarding its known limitations with channel synchronization.*

3. <b>Deconstruct the RELAY Paper:</b> Read the RELAY paper. Skip the dense formal notation and focus heavily on how it builds "function summaries" to pass lock and memory access information up the call graph.

4. <b>The Evaluation Report:</b> Summarize your findings. Why do static tools like Chronos struggle with complex modern language features like Go channels, while a framework like RELAY could scale across massive legacy codebases? What are the inherent trade-offs between false positives and false negatives in static analysis?
