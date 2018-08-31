# Chapter 12: Reflection

<!-- TOC -->

- [12.1. Why Reflection?](#121-why-reflection)
- [12.2. reflect.Type and reflect.Value](#122-reflecttype-and-reflectvalue)
- [12.3. Display, a Recursive Value Printer](#123-display-a-recursive-value-printer)
- [12.4. Example: Encoding S-Expressions](#124-example-encoding-s-expressions)
- [12.5. Setting Variables with reflect.Value](#125-setting-variables-with-reflectvalue)
- [12.6. Example: Decoding S-Expressions](#126-example-decoding-s-expressions)
- [12.7. Accessing Struct Field Tags](#127-accessing-struct-field-tags)
- [12.8. Displaying the Methods of a Type](#128-displaying-the-methods-of-a-type)
- [12.9. A Word of Caution](#129-a-word-of-caution)

<!-- /TOC -->


Go provides a mechanism to update variables and inspect their values at run time, to call their methods, and to apply the operations intrinsic to their representation, all without knowing their types at compile time. This mechanism is called *reflection*. Reflection also lets us treat types themselves as first-class values.

In this chapter, we’ll explore Go’s reflection features to see how they increase the expressiveness of the language, and in particular how they are crucial to the implementation of two important APIs: string formatting provided by `fmt`, and protocol encoding provided by packages like `encoding/json` and `encoding/xml`. Reflection is also essential to the template mechanism provided by the `text/template` and `html/template` packages we saw in Section 4.6. However, reflection is complex to reason about and not for casual use, so although these packages are implemented using reflection, they do not expose reflection in their own APIs.

## 12.1. Why Reflection? 
## 12.2. reflect.Type and reflect.Value 
## 12.3. Display, a Recursive Value Printer 
## 12.4. Example: Encoding S-Expressions 
## 12.5. Setting Variables with reflect.Value 
## 12.6. Example: Decoding S-Expressions 
## 12.7. Accessing Struct Field Tags 
## 12.8. Displaying the Methods of a Type 
## 12.9. A Word of Caution 
