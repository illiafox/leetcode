from collections import defaultdict

def dijkstra(num_nodes: int,
             edges: list[tuple[int, int, int]],
             start: int = 1,
             inf: int = 99) -> list[int]:
    # build adjacency list
    adj: dict[int, list[tuple[int, int]]] = defaultdict(list)
    for u, v, weight in edges:
        adj[u].append((v, weight))
        adj[v].append((u, weight))

    # init
    dist = [inf] * (num_nodes + 1)  # 1-based; index 0 is dummy
    seen = [False] * (num_nodes + 1)
    dist[start] = 0

    print(f"step 0: {dist[1:]}")

    # --- main loop ----------------------------------------------------------
    for step in range(num_nodes):
        # find the unseen vertex with the smallest tentative distance
        u = None
        best = inf + 1
        for v in range(1, num_nodes + 1):
            if not seen[v] and dist[v] < best:
                best, u = dist[v], v


        if u is None:  # remaining vertices unreachable
            break

        # print(f'checking node {u}')

        seen[u] = True

        # relax its outgoing edges
        for v, weight in adj[u]:
            if not seen[v] and dist[u] + weight < dist[v]:
                dist[v] = dist[u] + weight

        print(f"step {step+1:2d}: {dist[1:]}")
    return dist[1:]


edges = [
    (1, 3, 5), (1, 2, 6),
    (2, 4, 22),
    (3, 6, 2),
    (4, 7, 16),
    (5, 6, 14), (5, 8, 5), (5, 7, 2),
    (6, 8, 13),
    (7, 9, 14), (7, 10, 1), (7, 8, 19),
    (9, 10, 3)
]

if __name__ == "__main__":
    print('res', dijkstra(10, edges, start=1))
