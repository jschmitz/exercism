defmodule RobotSimulator do
  defstruct direction: :north, position: {0, 0}

  @doc """
  Create a Robot Simulator given an initial direction and position.

  Valid directions are: `:north`, `:east`, `:south`, `:west`
  """
  @spec create(direction :: atom, position :: {integer, integer}) :: any
  def create(direction \\ :north, position \\ {0, 0}) do
    with {:ok} <- valid_direction(direction),
         {:ok} <- valid_position(position) do
      %RobotSimulator{direction: direction, position: position}
    else
      {:error, "invalid direction"} -> {:error, "invalid direction"}
      {:error, "invalid position"} -> {:error, "invalid position"}
    end
  end

  @doc """
  Simulate the robot's movement given a string of instructions.

  Valid instructions are: "R" (turn right), "L", (turn left), and "A" (advance)
  """
  @spec simulate(robot :: any, instructions :: String.t()) :: any
  def simulate(robot, instructions) do
    if String.match?(instructions, ~r/^[LRA]*$/u) do
      Enum.reduce(String.codepoints(instructions), robot, fn instruction, acc ->
        process(acc, instruction)
      end)
    else
      {:error, "invalid instruction"}
    end
  end

  defp process(robot, instruction) when instruction == "L" do
    %{robot | direction: turn_left(robot.direction)}
  end

  defp process(robot, instruction) when instruction == "A" do
    %{robot | position: advance(robot.position, robot.direction)}
  end

  defp process(robot, instruction) when instruction == "R" do
    %{robot | direction: turn_right(robot.direction)}
  end

  defp turn_left(direction) do
    cond do
      direction == :north -> :west
      direction == :east -> :north
      direction == :south -> :east
      direction == :west -> :south
    end
  end

  defp turn_right(direction) do
    cond do
      direction == :north -> :east
      direction == :east -> :south
      direction == :south -> :west
      direction == :west -> :north
    end
  end

  defp advance(p, direction) do
    cond do
      direction == :north -> put_elem(p, 1, adv_north(p))
      direction == :east -> put_elem(p, 0, adv_east(p))
      direction == :south -> put_elem(p, 1, adv_south(p))
      direction == :west -> put_elem(p, 0, adv_west(p))
    end
  end

  defp adv_north(pos) do
    elem(pos, 1) + 1
  end

  defp adv_east(pos) do
    elem(pos, 0) + 1
  end

  defp adv_south(pos) do
    elem(pos, 1) - 1
  end

  defp adv_west(pos) do
    elem(pos, 0) - 1
  end

  @doc """
  Return the robot's direction.

  Valid directions are: `:north`, `:east`, `:south`, `:west`
  """
  @spec direction(robot :: any) :: atom
  def direction(robot) do
    robot.direction
  end

  @doc """
  Return the robot's position.
  """
  @spec position(robot :: any) :: {integer, integer}
  def position(robot) do
    robot.position
  end

  def valid_direction(direction) do
    if Enum.member?(Tuple.to_list({:north, :east, :south, :west}), direction) do
      {:ok}
    else
      {:error, "invalid direction"}
    end
  end

  def valid_position(position) do
    if is_tuple(position) && tuple_size(position) == 2 && is_integer(elem(position, 0)) &&
         is_integer(elem(position, 1)) do
      {:ok}
    else
      {:error, "invalid position"}
    end
  end
end
