defmodule Animal do
  def animal(name) do
    receive do
      :getName ->
        IO.puts(name)
        animal(name)
      :speak ->
        IO.puts("")
        animal(name)
      end
  end
end

defmodule Animal.Dog do
  def dog(name) do
    receive do
      :getName ->
        IO.puts(name)
        dog(name)
      :speak ->
        IO.puts("#{name} says: Woof!")
        dog(name)
    end
  end
end

defmodule Animal.Cat do
  def cat(name) do
    receive do
      :getName ->
        IO.puts(name)
        cat(name)
      :speak ->
        IO.puts("#{name} says: Meow!")
        cat(name)
    end
  end
end

defmodule Main do
  def start do
    a1 = spawn(Animal.Dog, :dog, ["Rex"])
    a2 = spawn(Animal.Cat, :cat, ["Whiskers"])
    send(a1, :speak)
    send(a2, :speak)
  end
end

Main.start()

# defprotocol Animal do
#   def getName(name)
#   def speak(name)
# end

# defimpl Animal, for: Dog do
#   def getName()

# end
