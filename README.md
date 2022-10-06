Just playing around with golang AST API, following along this post:

https://eli.thegreenplace.net/2021/rewriting-go-source-code-with-ast-tooling/

This should render as math $\sqrt{3}$
Let's require $\require{bussproofs}$

$$
\begin{prooftree}
\AxiomC{}
\RightLabel{Hyp$^{1}$}
\UnaryInfC{$P$}
\AXC{$P\to Q$}
\RL{$\to_E$}
\BIC{$Q^2$}
\AXC{$Q\to R$}
\RL{$\to_E$}
\BIC{$R$}
\AXC{$Q$}
\RL{Rit$^2$}
\UIC{$Q$}
\RL{$\wedge_I$}
\BIC{$Q\wedge R$}
\RL{$\to_I$$^1$}
\UIC{$P\to Q\wedge R$}
$$
```

