# Language definition

```ebnf
program = probability ("\n" probability)* EOF;

probability = NUMBER ("d" NUMBER)?
            | ("adv" | "dis" | "crit") "(" probability ")"
            | "reroll(" probability ("," probability)+ ")"
            | probability ( "*" | "/" | "+" | "-" ) probability
            | "-" probability
            | "(" probability ")"
            | "if" boolean "then" probability;

boolean = probability (">" | "<" | "=" | "!=") probability
        | probability "in(" (NUMBER)+ ")"
        | boolean ("and" | "or") boolean
        | "(" boolean ")";

NUMBER -> DIGIT+;
DIGIT -> "0" .. "9";
```

# Usage examples
[//]: <> (TODO - add good way to implement crits)

- Simple die throw 
`1d10 + 5`
- Advantage 
`adv(1d20)`
- If hits then roll damage 
`if (1d20 + 5 > 15) then (1d10)` 
- If hits then roll damage as a halfing 
`if (reroll(1d20 + 5, 1) > 15) then (1d10)`
- Roll all the die twice `crit(1d10 + 7)` -> `2d10 + 7`