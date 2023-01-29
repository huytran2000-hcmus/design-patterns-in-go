# Design pattern categories
- **Creational patterns**: These patterns provide various object creation mechanisms, which increases flexibility and code reuse.
- **Structural patterns**: These patterns explains how to assemble objects and classes into bigger structures while keeping this structure flexible and efficient.
- **Behavioral patterns**: These patterns are concerned with algorithms and assignments of responsibilities between objects.
# Catalog of patterns
## Creational patterns
### Intent
Provide an interface for creating objects in the superclass, but allows subclass to decide which type of object that will be created.
### Applicability
- Use when you don't know beforehand the exact types and dependencies of the objects your code should work with.
- Use when your want to provide users of your framework a way to extend its internal component.
- Use when you want to reuse existing objects instead of rebuilding them(factory method + prototype, factory method + singleton).