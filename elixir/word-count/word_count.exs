defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    sentence
      |> String.downcase
      |> String.replace(~r/[^\w\s\-\_]/u, "")
      |> String.replace(~r/[_]/u, " ")
      |> String.split
      |> Enum.reduce(%{}, fn (el, acc) -> Map.update(acc, el, 1, &((&1 +1))) end)
  end
end
