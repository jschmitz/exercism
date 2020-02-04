defmodule NucleotideCount do
  @nucleotides [?A, ?C, ?G, ?T]

  @doc """
  Counts individual nucleotides in a DNA strand.

  ## Examples

  iex> NucleotideCount.count('AATAA', ?A)
  4

  iex> NucleotideCount.count('AATAA', ?T)
  1
  """
  @spec count(charlist(), char()) :: non_neg_integer()

  def count('', nucleotide) do
    0
  end

  def count(strand, nucleotide) do
    Enum.reduce(strand, 0, fn(s, acc) ->
      if s == nucleotide do
        acc + 1
      else
        acc
      end
    end)
  end

  @doc """
  Returns a summary of counts by nucleotide.

  ## Examples

  iex> NucleotideCount.histogram('AATAA')
  %{?A => 4, ?T => 1, ?C => 0, ?G => 0}
  """
  @spec histogram(charlist()) :: map()
  def histogram(strand) do
    Map.new(@nucleotides, fn(n) -> {n, count(strand,n)} end)
  end
end
