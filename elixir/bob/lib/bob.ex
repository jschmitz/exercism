defmodule Bob do
  def hey(input) do
    cond do
      nothing?(input) -> "Fine. Be that way!"
      question?(input) && shouting?(input) -> "Calm down, I know what I'm doing!"
      shouting?(input) ->  "Whoa, chill out!"
      question?(input) ->  "Sure."
      true -> "Whatever."
    end
  end

  def nothing?(input) do
    String.trim(input) == ""
  end

  def question?(input) do
    String.last(input) == "?"
  end

  def shouting?(input) do
    input == String.upcase(input) && String.match?(input, ~r/[\p{L}]/)
  end 
end
