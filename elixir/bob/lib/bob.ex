defmodule Bob do
  def hey(input) do
    if String.trim(input) == "" do
      "Fine. Be that way!"
    else

      last_char = String.at(input, String.length(input) - 1)
      punc? = String.match?(last_char, ~r/[?!.]/)
      all_upper = input != String.upcase(input)
      has_letters = String.match?(input, ~r/[\p{L}]/)

      if punc? do
        if has_letters do
          result(last_char, all_upper)
        else
          "Sure."
        end
      else
        if has_letters do
          "Whoa, chill out!"
        else
          "Whatever."
        end
      end
    end
  end

  def result(p, caps?) when p == "." and caps?, do: "Whatever."
  def result(p, caps?) when p == "!" and not caps?, do: "Whoa, chill out!"
  def result(p, caps?) when p == "!" and caps?, do: "Whatever."
  def result(p, caps?) when p == "?" and caps?, do: "Sure."
  def result(p, caps?) when p == "?" and not caps?, do: "Calm down, I know what I'm doing!"
end
