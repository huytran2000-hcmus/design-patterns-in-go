# Design pattern categories
- **Creational patterns**: These patterns provide various object creation mechanisms, which increase flexibility and code reuse.
- **Structural patterns**: These patterns explains how to assemble objects and classes into bigger structures while keeping this structure flexible and efficient.
- **Behavioral patterns**: These patterns are concerned with algorithms and assignments of responsibilities between objects.
# Catalog of patterns
## Creational patterns
### Singleton pattern
#### Intent
**Singleton** is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this pattern
#### Applicability
- Use when a class should have a single instance that get shared with all clients.
- Use when a resource is expensive to create and can be shared within the program like database connection, ssh connection...
- Use to stricter control over global variables.
### Factory Method pattern
#### Intent
**Factory Method** provides an interface for creating objects in the superclass, but allows subclasses to decide which type of object that will be created.
#### Applicability
- Use when you don't know beforehand the exact types and dependencies of the objects your code will work with.
- Use when your want to provide users of your framework a way to extend its internal component.
- Use when you want to reuse existing objects instead of rebuilding them(factory method + prototype, factory method + singleton).
### Abstract Factory Pattern
#### Intent
**Abstract Factory** lets you produce objects of related families without specifying their concrete classes.
#### Applicability
- Use when your code need to work with various families of related products, but you may not know them beforehand or you want allow future extensibility.
### Builder pattern
#### Intent
**Builder** allows you to produce different types and representations of the object step by step using the same construction code.
#### Applicability
- Use when you to construct complex objects with many configuration.
- Use when you want your code to be able to create different representations of some products.
### Prototype pattern
#### Intent
**Prototype** lets you copy existing objects without making your code depend on their classes.
#### Applicability
- Use when your code shouldn't depend on concrete classes of objects that you need to copy.
- Use when you want to reduce the number of subclasses that only differ in their initialization.