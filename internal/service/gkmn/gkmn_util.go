package gkmn

import gkmnv1 "github.com/jchapman63/neo-gkmn/internal/service/gkmn/v1"

func (h *GameHandler) MapBattleMonsters(battleID string) []*gkmnv1.BattleMonster {

	for _, b := range h.activeBattles {
		if b.ID == battleID {
			respMons := []*gkmnv1.BattleMonster{}
			for _, m := range b.Monsters {
				respMoves := []*gkmnv1.Move{}
				for _, mo := range m.Moves {
					respMove := &gkmnv1.Move{
						Name:  mo.Name,
						Id:    mo.ID,
						Type:  mo.Type,
						Power: mo.Power,
					}
					respMoves = append(respMoves, respMove)
				}
				respMon := &gkmnv1.BattleMonster{
					Monster: &gkmnv1.Monster{
						Name:   m.Monster.Name,
						Id:     m.Monster.ID,
						Type:   m.Monster.Type,
						BaseHp: m.Monster.Basehp,
					},
					LiveHp: m.LiveHp,
					Speed:  m.Speed,
					Moves:  respMoves,
				}
				respMons = append(respMons, respMon)
			}
			return respMons
		}
	}
	return []*gkmnv1.BattleMonster{}
}
