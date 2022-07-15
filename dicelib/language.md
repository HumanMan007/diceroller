# Language definition

```ebnf
program = probability|boolean EOF;

probability = NUMBER ("d" NUMBER)?
            | ("adv" | "dis") paren<probability>
            | "reroll" paren<probability ("," probability)+>
            | probability ( "*" | "/" | "+" | "-" ) probability
            | "-" probability
            | paren<probability>
            | "if" paren<boolean>
              "then" paren<probability> 
              ("else" paren<probability>)?;

boolean = probability (">" | "<" | "=" | "!=") probability
        | boolean ("and" | "or") boolean
        | paren<boolean>;

(* Using <> to represent an argument in a rule *)
paren<member> = "(" member ")"
              | "[" member "]"
              | "{" member "}"

NUMBER -> DIGIT+;
DIGIT -> "0" .. "9";
```

# Usage examples
[//]: <> (TODO - add good way to implement crits)
[//]: <> (TODO - add good way to implement equals element from list)
- Simple die throw 
`1d10 + 5`
- Advantage 
`adv(1d20)`
- If hits then roll damage 
`if [1d20 + 5 > 15] then {1d10}` 
- If hits then roll damage as a halfing 
`if [reroll(1d20, 1) + 5 > 15] then {1d10}`