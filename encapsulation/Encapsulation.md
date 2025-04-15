# Encapsulation in Go

Encapsulation in Go is achieved through the use of **exported** and **unexported identifiers**. 

Go does not have the concept of **classes and objects** in the same way as other object-oriented programming languages, but it provides mechanisms to achieve encapsulation at the **package level**.

In Go:

- An **identifier starting with an uppercase letter** is **exported** and accessible from other packages.
- An **identifier starting with a lowercase letter** is **unexported** and accessible **only within the same package**.

This mechanism allows for:
- Bundling data and related methods
- Restricting direct access from outside the package
- Achieving encapsulation

---

## 🧠 Difference between Encapsulation and Abstraction

| Concept        | Description |
|----------------|-------------|
| **Encapsulation** | "I'm keeping my variables safe and secure inside this box." |
| **Abstraction**   | "You don't need to know how this box works inside, just press the button." |

- **Encapsulation** focuses on **data protection** and **access restriction**.
- **Abstraction** focuses on **hiding complexity** and **showing only essential details**.

