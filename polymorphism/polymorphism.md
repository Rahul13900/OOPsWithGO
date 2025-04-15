# Polymorphism in Golang

Polymorphism in **Golang** is achieved through **interfaces**. Unlike traditional object-oriented languages like Java or C++, Go does not use class-based inheritance. Instead, it relies on interfaces to achieve polymorphism.

---

## 🔁 What is Polymorphism?

Polymorphism allows writing code that can work with different types through a common interface. It lets you treat different types uniformly if they implement the same behavior.

---
Interfaces in Go are implicit – types implement interfaces just by having the required methods.

The empty interface (interface{}) can hold values of any type, but it's not type-safe.

Interface-based polymorphism promotes decoupled, testable, and clean code design.