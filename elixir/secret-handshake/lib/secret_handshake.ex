defmodule SecretHandshake do
  @actions ["jump", "close your eyes", "double blink", "wink"]

  @doc """
  Determine the actions of a secret handshake based on the binary
  representation of the given `code`.

  If the following bits are set, include the corresponding action in your list
  of commands, in order from lowest to highest.

  1 = wink
  10 = double blink
  100 = close your eyes
  1000 = jump

  10000 = Reverse the order of the operations in the secret handshake
  """
  @spec commands(code :: integer) :: list(String.t())

  def commands(code) when code > 31 do
    []
  end

  def commands(code) when code < 1 do
    []
  end

  def commands(code) when code > 16 and code < 32 do
    code - 16
    |> Integer.to_string(2)
    |> String.pad_leading(4, "0")
    |> String.codepoints
    |> Enum.zip(@actions)
    |> Enum.filter( fn(x) -> elem(x, 0) == "1" end )
    |> Enum.map( fn(x) -> elem(x, 1) end )
  end

  def commands(code) do
    do_commands(code)
    |> Enum.reverse
  end

  def do_commands(code) do
    |> Integer.to_string(2)
    |> String.pad_leading(4, "0")
    |> String.codepoints
    |> Enum.zip(@actions)
    |> Enum.filter( fn(x) -> elem(x, 0) == "1" end )
    |> Enum.map( fn(x) -> elem(x, 1) end )
  end


end
