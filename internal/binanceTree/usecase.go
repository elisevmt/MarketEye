package binanceTree

import (
	"MarketEye/pkg/graphZero"
	"github.com/gorilla/websocket"
)

type UseCase interface {
	TreeLoad()                                                      // TreeLoad() - ключевая функция в бинанс арбитраже - инициализирует всё валютное дерево биржи и соединяет между собой валютные узлы
	BatchTreeLoad()                                                 // BatchTreeLoad - улучшенная версия TreeLoad(), использующая один соккет для получения инфомрации по всемв валютам
	BranchUpdate(branchRaw, branchReverse *graphZero.Branch)        // BranchUpdate - функция для обновления данных в ветке между узлами
	BatchBranchUpdate(branchRaw, branchReverse []*graphZero.Branch) // BatchBranchUpdate - функция, принимающая список веток и обновляющая их в БАТЧ WSS подключении
	BatchConnectionResolver(conn *websocket.Conn)                   // BatchConnectionResolver - Резолвит соединение с бинансом, записывая все данные с потоков в дерево
}
