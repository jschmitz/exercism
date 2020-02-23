defmodule RotationalCipher do

  def rotate(text, shift) when text == " " do
    " "
  end

  @doc """
  Given a plaintext and amount to shift by, return a rotated string.

  Example:
  iex> RotationalCipher.rotate("Attack at dawn", 13)
  "Nggnpx ng qnja"
  """
  @spec rotate(text :: String.t(), shift :: integer) :: String.t()
  def rotate(text, shift) do
    if String.length(text) > 1 do
      IO.puts("text is: #{text}")
      text
      |> String.codepoints
      |> Enum.map(fn(c) -> rotate(c, shift) end)
      |> List.to_string
    else
      IO.puts("text is #{text}")
      IO.puts("shift is: #{shift}")

      cta = charToAscii(text)
      IO.puts("text ascii is: #{cta}")

      if (cta > 96 and cta < 123) or (cta > 64 and cta < 91) do
        shift_rem = rem(shift, 26)

        r = rotate(cta + shift_rem)
        IO.puts("rotated #{r}")

        asciiToString(r)
      else
        text
      end

    end
  end

  def rotate(text, shift) when text == " " do
    " "
  end

  defp rotateM(text, shift) do
    text
    |> String.codepoints
    |> Enum.reduce("", fn c -> rotate(c, shift) end)
  end


  defp remainder(shift) do
    rem(shift, 26)
  end

  defp charToAscii(text) do
    text
    |> String.to_charlist
    |> hd
  end

  defp rotate(a) do
    if a > 122 do
      a - 26
    else
      a
    end
  end

  defp asciiToString(ascii) do
    List.to_string([ascii])
  end
end

