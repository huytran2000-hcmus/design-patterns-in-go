# Design pattern categories
- **Creational patterns**: These patterns provide various object creation mechanisms, which increase flexibility and code reuse.
- **Structural patterns**: These patterns explains how to assemble objects and classes into bigger structures while keeping this structure flexible and efficient.
- **Behavioral patterns**: These patterns are concerned with algorithms and assignments of responsibilities between objects.
# Catalog of patterns
## Creational patterns
### Singleton Pattern
#### Intent
**Singleton** is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.  
![pattern image](/images/single_pattern.png 'singleton pattern')
#### Applicability
- Use when a class should have a single instance that get shared with all clients.
- Use when a resource is expensive to create and can be shared within the program like database connection, ssh connection,...
- Use when need stricter control over global variables.
### Factory Method Pattern
#### Intent
**Factory Method** provides an interface for creating objects in the superclass, but allows subclasses to decide which type of object that will be created.  
![pattern image](/images/factory_method_pattern.png 'factory method pattern')
#### Applicability
- Use when you don't know beforehand the exact types and dependencies of the objects your code will work with.
- Use when your want to provide users of your framework a way to extend its internal component.
- Use when you want to reuse existing objects instead of rebuilding them(factory method + prototype, factory method + singleton).
### Abstract Factory Pattern
#### Intent
**Abstract Factory** lets you produce objects of related families without specifying their concrete classes.  
![pattern image](/images/abstract_factory_pattern.png 'abstract factory pattern')
#### Applicability
- Use when your code need to work with various families of related products, but you may not know them beforehand ,or you want to allow future extensibility.
### Builder Pattern
#### Intent
**Builder** allows you to produce different types and representations of the object step by step using the same construction code.  
![pattern image](/images/builder_pattern.png 'builder pattern')
#### Applicability
- Use when you to construct complex objects with many configuration.
- Use when you want your code to be able to create different representations of some products.
### Prototype Pattern
#### Intent
**Prototype** lets you copy existing objects without making your code depend on their classes.  
![pattern image](/images/protype_pattern.png 'prototype pattern')
#### Applicability
- Use when your code shouldn't depend on concrete classes of objects that you need to copy.
- Use when you want to reduce the number of subclasses that only differ in their initialization.
## Structural Design Patterns
### Adapter Pattern
#### Intent
**Adapter** allows objects with incompatible interfaces to collaborate.
#### Applicability
- Use when you want to use some existing class, but its interface isn't compatible with your code.
- Use when you want to reuse existing subclasses that lack some common missing functionality that can't be added to the superclass.
