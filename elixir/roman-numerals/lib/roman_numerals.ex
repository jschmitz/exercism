defmodule RomanNumerals do
  @doc """
  Convert the number to a roman number.
  """
  @spec numeral(pos_integer) :: String.t()

  def numeral(number) when number < 4 do
    String.duplicate("I", number)
  end

  def numeral(number) when number > 3 and number < 5 do
    "IV"
  end

  def numeral(number) when number > 4 and number < 6 do
    "V"
  end

  def numeral(number) when number > 5 and number < 9 do
    "V" <> numeral(number - 5)
  end

  def numeral(number) when number > 8 and number < 10 do
    "IX"
  end

  def numeral(number) when number > 9 and number < 41 do
    "X" <> numeral(number - 10)
  end

  def numeral(number) when number > 40 and number < 50 do
    "XL" <> numeral(number - 40)
  end

  def numeral(number) when number > 49 and number < 90 do
    "L" <> numeral(number - 50)
  end

  def numeral(number) when number > 90 and number < 100 do
    "XC" <> numeral(number - 90)
  end

  def numeral(number) when number > 99 and number < 400 do
    "C" <> numeral(number - 100)
  end

  def numeral(number) when number > 400 and number < 500 do
    "CD" <> numeral(number - 400)
  end

  def numeral(number) when number > 499 and number < 900 do
    "D" <> numeral(number - 500)
  end

  def numeral(number) when number > 900 and number < 1000 do
    "CM" <> numeral(number - 900)
  end

  def numeral(number) when number > 999 do
    "M" <> numeral(number - 1000)
  end
end

# I = 1
# V = 5
# X = 10
# L = 50
# C = 100
# D = %00
# M = 1000
