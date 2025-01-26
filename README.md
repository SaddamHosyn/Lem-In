# Ant Farm Pathfinding

This project is a Go implementation of an ant farm pathfinding algorithm. It reads a file describing an ant farm, including the number of ants, rooms, and connections between them. The program then calculates the shortest path for the ants to travel from the `##start` room to the `##end` room, ensuring that no two ants occupy the same room at the same time.

## Features

- **Pathfinding**: Uses BFS (Breadth-First Search) and DFS (Depth-First Search) to find the shortest path for ants.
- **Validation**: Ensures the input file is valid and adheres to the required format.
- **Error Handling**: Provides clear error messages for invalid inputs or unsolvable scenarios.

---

## Installation

1. Ensure you have Go installed on your system. If not, download and install it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/SaddamHosyn/Lem-In.git

   ```

---

## Usage

To run the program, use the following command:

```bash
go run . <filename>
```

Replace `<filename>` with the path to your input file (e.g., `example00.txt`).

### Example

```bash
go run . example00.txt
```

---

## Input Format

The input file must follow this structure:

- The first line specifies the number of ants.
- Rooms are defined as `name x y`, where `x` and `y` are coordinates.
- Connections between rooms are defined as `room1-room2`.
- Special rooms are marked with `##start` and `##end`.

### Example Input

```
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1
```

---

## Output Format

The program outputs the movements of the ants in the following format:

- Each line represents a turn.
- Ant movements are represented as `Lx-y`, where `x` is the ant's ID and `y` is the room it moves to.

### Example Output

```
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
```

---

### 1. **Allowed Packages**

- The program only uses standard Go packages.

### 2. **Reading Input**

- The program can read the ant farm description from a file (e.g., `example00.txt`).

### 3. **Command Validation**

- The program only accepts `##start` and `##end` as valid commands.

### 4. **Output Format**

- The output follows the required format:
  ```
  Lx-y
  Lx-y Lx-y
  Lx-y Lx-y Lx-y
  ```

---

## Error Handling

The program handles the following errors:

- Invalid file format.
- Duplicate room names or coordinates.
- Missing `##start` or `##end` rooms.
- Unconnected rooms.
- Invalid ant numbers (e.g., zero or negative).

---

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

---
