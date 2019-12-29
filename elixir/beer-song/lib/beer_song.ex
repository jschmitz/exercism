defmodule BeerSong do
  @doc """
  Get a single verse of the beer song
  """
  @spec verse(integer) :: String.t()
  def verse(number) do

    if number == 0 do
      "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
    else

      plural = fn
        n when n > 1 -> "s"
        n when n == 1 -> ""
        n when n <= 0 -> "s"
      end

      one_or_it = if number > 1, do: "one", else: "it"
      i_or_nomore = if number > 1, do: "#{number - 1}", else: "no more"

      "#{number} bottle#{plural.(number)} of beer on the wall, #{number} bottle#{plural.(number)} of beer.\nTake #{one_or_it} down and pass it around, #{i_or_nomore} bottle#{plural.(number - 1)} of beer on the wall.\n"
    end
  end

  @doc """
  Get the entire beer song for a given range of numbers of bottles.
  """
  @spec lyrics(Range.t()) :: String.t()
  def lyrics(range) do
    s = Enum.reduce(range, "", fn i, acc -> acc <> verse(i) <> "\n" end)
    String.slice(s, 0..-2)
  end

  def lyrics() do
    lyrics(99..0)
  end
end
