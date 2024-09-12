#!/usr/bin/python3

import csv
import sys
from grafo import grafo
from biblioteca import *

FIRST_VERTEX = 0
SECOND_VERTEX = 1
CANTIDAD_PARAMETROS_VALIDA = 4
POSICION_ARCHIVO = 1
POSICION_COMANDO = 0
POSICION_ORIGEN = 1
POSICION_DESTINO = 2
POSICION_CANTIDAD = 1
POSICION_DELINCUENTES = 1
POSICION_K_MAS_IMPORTANTES = 2
POSICION_CANT_SALTOS = 2

def cargar_grafo(file_name):
    grafo_instancia = grafo(es_dirigido=True)
    try:
        with open(file_name) as file:
            reader = csv.reader(file, delimiter="\t")
            for line in reader:
                grafo_instancia.agregar_arista(int(line[FIRST_VERTEX]), int(line[SECOND_VERTEX]))
    except Exception as e:
        return f"No se pudo abrir {file_name}: {e}"
    return grafo_instancia


def main(argv):
    if len(argv) < 2:
        print("cant parametros inválida")
        return
    archivo = argv[POSICION_ARCHIVO]
    grafo_instancia = cargar_grafo(archivo)
    pagerank_scores = None
    while True:
        comando = input().strip()
        partes = comando.split()
        cmd = partes[POSICION_COMANDO]
        if cmd == "min_seguimientos":
            origen = int(partes[POSICION_ORIGEN])
            destino = int(partes[POSICION_DESTINO])
            resultado = distancia_minima(grafo_instancia, origen, destino)
            if resultado:
                print(" -> ".join(map(str, resultado)))
            else:
                print("Seguimiento imposible")
        elif cmd == "mas_imp": # pageRank(ver enunciado tp)
            cant = int(partes[POSICION_CANTIDAD])
            resultado, pagerank_scores = mas_imp(grafo_instancia, cant, pagerank_scores)
            print(", ".join(map(str, resultado)))
        elif cmd == "persecucion": # dijkstra por cada elemento de delincuentes a cada elemento de importantes(va a ser cuadratico)
            delincuentes = list(map(int, partes[POSICION_DELINCUENTES].split(",")))
            k = int(partes[POSICION_K_MAS_IMPORTANTES])
            importantes, pagerank_scores = mas_imp(grafo_instancia, k, pagerank_scores)
            resultado = min_recorrido(grafo_instancia, delincuentes, importantes)
            if resultado:
                print(" -> ".join(map(str, resultado)))
            else:
                print("Seguimientos imposibles")
        elif cmd == "comunidades": # label propagation(ver enunciado tp)
            n = int(partes[POSICION_CANTIDAD])
            resultado = comunidades(grafo_instancia, n)
            for i, comunidad in enumerate(resultado, start=1):
                print(f"Comunidad {i}: {', '.join(map(str, comunidad))}")
        elif cmd == "divulgar": # Encontrar ciclos de n elementos
            delincuente = int(partes[POSICION_DELINCUENTES])
            n = int(partes[POSICION_CANT_SALTOS])
            resultado = bfs_n_distancia(grafo_instancia, delincuente, n)
            if resultado:
                print(", ".join(map(str, resultado)))
        elif cmd == "divulgar_ciclo": # encontrar el menor ciclo que involucre a dicho delincuente
            delincuente = int(partes[POSICION_DELINCUENTES])
            resultado = ciclo_mas_corto(grafo_instancia, delincuente)
            if resultado:
                print(" -> ".join(map(str, resultado)))
            else:
                print("No se encontro recorrido")
        elif cmd == "cfc":#tarjan para componentes fuertemente conexas
            resultado = cfc(grafo_instancia)
            for i, componente in enumerate(resultado, start=1):
                print(f"CFC {i}: {', '.join(map(str, componente))}")

if __name__ == "__main__":
    main(sys.argv)
