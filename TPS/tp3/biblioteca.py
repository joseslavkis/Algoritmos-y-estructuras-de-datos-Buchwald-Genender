from collections import deque, defaultdict, Counter
import random

def distancia_minima(grafo, origen, destino):
    if origen not in grafo.obtener_vertices() or destino not in grafo.obtener_vertices():
        return None
    cola = deque()
    cola.append((origen, [origen]))
    visitados = set()
    visitados.add(origen)
    while cola:
        nodo_actual, camino_actual = cola.popleft()
        if nodo_actual == destino:
            return camino_actual
        for vecino in grafo.adyacentes(nodo_actual):
            if vecino not in visitados:
                visitados.add(vecino)
                cola.append((vecino, camino_actual + [vecino]))
    return None


def min_recorrido(grafo, vectorOrigen, vectorDestino):
    min_camino = None
    min_longitud = float('inf')

    for v in vectorOrigen:
        for w in vectorDestino[::-1]:
            camino_actual = distancia_minima(grafo, v, w)
            if camino_actual is not None:
                longitud_actual = len(camino_actual)
                if longitud_actual < min_longitud:
                    min_longitud = longitud_actual
                    min_camino = camino_actual
                elif longitud_actual == min_longitud:
                    min_camino = camino_actual
    return min_camino

def calcular_pagerank_personalizado(grafo, num_walks=100, walk_length=500):
    pagerank_scores = Counter()

    for _ in range(num_walks):
        vertice_inicial = grafo.vertice_aleatorio()
        for _ in range(walk_length):
            if not grafo.adyacentes(vertice_inicial):
                break
            valor_transmitido = 1 / len(grafo.adyacentes(vertice_inicial))
            vertice_siguiente = random.choice(grafo.adyacentes(vertice_inicial))
            pagerank_scores[vertice_siguiente] += valor_transmitido
            vertice_inicial = vertice_siguiente

    return pagerank_scores

def mas_imp(grafo, cant, pagerank_scores):
    if pagerank_scores is None:
        pagerank_scores = calcular_pagerank_personalizado(grafo)
    importantes = pagerank_scores.most_common(cant)
    return [vertice for vertice, score in importantes], pagerank_scores

def comunidades(grafo, n):
    etiquetas = {v: v for v in grafo.obtener_vertices()}
    cambio = True

    while cambio:
        cambio = False
        vertices = list(grafo.obtener_vertices())
        random.shuffle(vertices)
        for v in vertices:
            etiquetas_vecinos = [etiquetas[vecino] for vecino in grafo.adyacentes(v)]
            if etiquetas_vecinos:
                etiqueta_mas_comun = max(set(etiquetas_vecinos), key=etiquetas_vecinos.count)
                if etiquetas[v] != etiqueta_mas_comun:
                    etiquetas[v] = etiqueta_mas_comun
                    cambio = True
    comunidades = defaultdict(list)
    for v, etiqueta in etiquetas.items():
        comunidades[etiqueta].append(v)

    return [comunidad for comunidad in comunidades.values() if len(comunidad) >= n]

def bfs_n_distancia(grafo, vertice, n):
    if  vertice not in grafo.obtener_vertices():
        return None
    visitados = set()
    q = deque([( vertice, 0)])
    visitados.add( vertice)
    resultado = []

    while q:
        v, dist = q.popleft()
        
        if dist <= n and v !=  vertice:
            resultado.append(v)
        
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                q.append((w, dist + 1))
    return resultado


def ciclo_mas_corto(grafo, origen):
    if origen not in grafo.obtener_vertices():
        return None
    if origen in grafo.adyacentes(origen):
        return [origen, origen]
    cola = deque([(origen, [origen])])
    visitados = set()

    menor_ciclo = None
    
    while cola:
        nodo_actual, camino_actual = cola.popleft()

        for vecino in grafo.adyacentes(nodo_actual):
            if vecino == origen and len(camino_actual) > 1:
                ciclo_actual = camino_actual + [origen]
                if menor_ciclo is None or len(ciclo_actual) < len(menor_ciclo):
                    menor_ciclo = ciclo_actual
                    break
            elif vecino not in visitados:
                visitados.add(vecino)
                cola.append((vecino, camino_actual + [vecino]))

    return menor_ciclo


def dfs_cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global):
    orden[v] = mas_bajo[v] = contador_global[0]
    contador_global[0] += 1
    visitados.add(v)
    pila.append(v)
    apilados.add(v)

    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs_cfc(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])

    if orden[v] == mas_bajo[v]:
        nueva_cfc = []
        while True:
            w = pila.pop()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == v:
                break
        cfcs.append(nueva_cfc)

def cfc(grafo):
    visitados = set()
    orden = {}
    mas_bajo = {}
    pila = []
    apilados = set()
    cfcs = []
    contador_global = [0]

    for v in grafo.obtener_vertices():
        if v not in visitados:
            dfs_cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global)

    return cfcs
