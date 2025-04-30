 Program in Elixir that implements this:
 
```
abstract class Animal {
     private String name;              // encapsulation

     public Animal(String name) { this.name = name; }

     public String getName() { return name; }   // encapsulation

     public abstract void speak();   // abstraction + polymorphism
   }

   class Dog extends Animal {
     public Dog(String name) { super(name); }

     @Override
     public void speak() { System.out.println(getName() + " says: Woof!"); }
   }

   class Cat extends Animal {
     public Cat(String name) { super(name); }

     @Override
     public void speak() { System.out.println(getName() + " says: Meow!"); }
   }

   public class Main {
     public static void main(String[] args) {
       Animal a1 = new Dog("Rex");
       Animal a2 = new Cat("Whiskers");
  
       a1.speak(); // polymorphism
       a2.speak();
     }
   }
```

Your goal is to design and write Elixir code that provides the same OO "pillars" (encapsulation, abstraction, inheritance, and polymorphism) that Java does, but in a functional, actor-based language with no objects, no classes, and no class-based inheritance.

In your Elixir, for processes use the basic spawn_link approach, and do not use higher level packages like GenServer, etc. Just use the basic Elixir language features.

Questions for Discussion (readme file)
Explain what features in Elixir you are using to provide the objects and OO pillars you are implementing. You may or may not have all of them... so explain the ones you do have in your code solution. Explain what parts of your solution maps to what part of the Java example you are implementing.

