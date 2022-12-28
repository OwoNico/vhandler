package vhandler

import (
	"errors"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/venitynetwork/vhandler/priority"
)

type Player struct {
	handlers map[handlerId]*subHandler
	h        *PlayerHandler
}

func NewPlayer() *Player {
	v := &Player{handlers: map[handlerId]*subHandler{}}
	v.h = newPlayerHandler(v)

	v.handlers[PlayerItemDropId] = newSubHandler()
	v.handlers[PlayerMoveId] = newSubHandler()
	v.handlers[PlayerJumpId] = newSubHandler()
	v.handlers[PlayerTeleportId] = newSubHandler()
	v.handlers[PlayerChangeWorldId] = newSubHandler()
	v.handlers[PlayerToggleSprintId] = newSubHandler()
	v.handlers[PlayerToggleSneakId] = newSubHandler()
	v.handlers[PlayerCommandExecutionId] = newSubHandler()
	v.handlers[PlayerTransferId] = newSubHandler()
	v.handlers[PlayerChatId] = newSubHandler()
	v.handlers[PlayerSkinChangeId] = newSubHandler()
	v.handlers[PlayerStartBreakId] = newSubHandler()
	v.handlers[PlayerBlockBreakId] = newSubHandler()
	v.handlers[PlayerBlockPlaceId] = newSubHandler()
	v.handlers[PlayerBlockPickId] = newSubHandler()
	v.handlers[PlayerSignEditId] = newSubHandler()
	v.handlers[PlayerItemPickupId] = newSubHandler()
	v.handlers[PlayerItemUseId] = newSubHandler()
	v.handlers[PlayerItemUseOnBlockId] = newSubHandler()
	v.handlers[PlayerItemUseOnEntityId] = newSubHandler()
	v.handlers[PlayerItemConsumeId] = newSubHandler()
	v.handlers[PlayerItemDamageId] = newSubHandler()
	v.handlers[PlayerAttackEntityId] = newSubHandler()
	v.handlers[PlayerExperienceGainId] = newSubHandler()
	v.handlers[PlayerPunchAirId] = newSubHandler()
	v.handlers[PlayerHurtId] = newSubHandler()
	v.handlers[PlayerHealId] = newSubHandler()
	v.handlers[PlayerFoodLossId] = newSubHandler()
	v.handlers[PlayerDeathId] = newSubHandler()
	v.handlers[PlayerRespawnId] = newSubHandler()
	v.handlers[PlayerQuitId] = newSubHandler()

	return v
}

func (v *Player) OnItemDrop(p priority.Priority, h PlayerItemDropHandler) {
	v.handlers[PlayerItemDropId].add(p, h)
}

func (v *Player) OnMove(p priority.Priority, h PlayerMoveHandler) {
	v.handlers[PlayerMoveId].add(p, h)
}

func (v *Player) OnJump(p priority.Priority, h PlayerJumpHandler) {
	v.handlers[PlayerJumpId].add(p, h)
}

func (v *Player) OnTeleport(p priority.Priority, h PlayerTeleportHandler) {
	v.handlers[PlayerTeleportId].add(p, h)
}

func (v *Player) OnChangeWorld(p priority.Priority, h PlayerChangeWorldHandler) {
	v.handlers[PlayerChangeWorldId].add(p, h)
}

func (v *Player) OnToggleSprint(p priority.Priority, h PlayerToggleSprintHandler) {
	v.handlers[PlayerToggleSprintId].add(p, h)
}

func (v *Player) OnToggleSneak(p priority.Priority, h PlayerToggleSneakHandler) {
	v.handlers[PlayerToggleSneakId].add(p, h)
}

func (v *Player) OnCommandExecution(p priority.Priority, h PlayerCommandExecutionHandler) {
	v.handlers[PlayerCommandExecutionId].add(p, h)
}

func (v *Player) OnTransfer(p priority.Priority, h PlayerTransferHandler) {
	v.handlers[PlayerTransferId].add(p, h)
}

func (v *Player) OnChat(p priority.Priority, h PlayerChatHandler) {
	v.handlers[PlayerChatId].add(p, h)
}

func (v *Player) OnSkinChange(p priority.Priority, h PlayerSkinChangeHandler) {
	v.handlers[PlayerSkinChangeId].add(p, h)
}

func (v *Player) OnStartBreak(p priority.Priority, h PlayerStartBreakHandler) {
	v.handlers[PlayerStartBreakId].add(p, h)
}

func (v *Player) OnBlockBreak(p priority.Priority, h PlayerBlockBreakHandler) {
	v.handlers[PlayerBlockBreakId].add(p, h)
}

func (v *Player) OnBlockPlace(p priority.Priority, h PlayerBlockPlaceHandler) {
	v.handlers[PlayerBlockPlaceId].add(p, h)
}

func (v *Player) OnBlockPick(p priority.Priority, h PlayerBlockPickHandler) {
	v.handlers[PlayerBlockPickId].add(p, h)
}

func (v *Player) OnSignEdit(p priority.Priority, h PlayerSignEditHandler) {
	v.handlers[PlayerSignEditId].add(p, h)
}

func (v *Player) OnItemPickup(p priority.Priority, h PlayerItemPickupHandler) {
	v.handlers[PlayerItemPickupId].add(p, h)
}

func (v *Player) OnItemUse(p priority.Priority, h PlayerItemUseHandler) {
	v.handlers[PlayerItemUseId].add(p, h)
}

func (v *Player) OnItemUseOnBlock(p priority.Priority, h PlayerItemUseOnBlockHandler) {
	v.handlers[PlayerItemUseOnBlockId].add(p, h)
}

func (v *Player) OnItemUseOnEntity(p priority.Priority, h PlayerItemUseOnEntityHandler) {
	v.handlers[PlayerItemUseOnEntityId].add(p, h)
}

func (v *Player) OnItemConsume(p priority.Priority, h PlayerItemConsumeHandler) {
	v.handlers[PlayerItemConsumeId].add(p, h)
}

func (v *Player) OnItemDamage(p priority.Priority, h PlayerItemDamageHandler) {
	v.handlers[PlayerItemDamageId].add(p, h)
}

func (v *Player) OnAttackEntity(p priority.Priority, h PlayerAttackEntityHandler) {
	v.handlers[PlayerAttackEntityId].add(p, h)
}

func (v *Player) OnExperienceGain(p priority.Priority, h PlayerExperienceGainHandler) {
	v.handlers[PlayerExperienceGainId].add(p, h)
}

func (v *Player) OnPunchAir(p priority.Priority, h PlayerPunchAirHandler) {
	v.handlers[PlayerPunchAirId].add(p, h)
}

func (v *Player) OnHurt(p priority.Priority, h PlayerHurtHandler) {
	v.handlers[PlayerHurtId].add(p, h)
}

func (v *Player) OnHeal(p priority.Priority, h PlayerHealHandler) {
	v.handlers[PlayerHealId].add(p, h)
}

func (v *Player) OnFoodLoss(p priority.Priority, h PlayerFoodLossHandler) {
	v.handlers[PlayerFoodLossId].add(p, h)
}

func (v *Player) OnDeath(p priority.Priority, h PlayerDeathHandler) {
	v.handlers[PlayerDeathId].add(p, h)
}

func (v *Player) OnRespawn(p priority.Priority, h PlayerRespawnHandler) {
	v.handlers[PlayerRespawnId].add(p, h)
}

func (v *Player) OnQuit(p priority.Priority, h PlayerQuitHandler) {
	v.handlers[PlayerQuitId].add(p, h)
}

func (v *Player) Set(p *player.Player) {
	p.Handle(v.h)
}

func (v *Player) Remove(h Handler) error {
	switch h.(type) {
	case PlayerItemDropHandler:
		return v.handlers[PlayerItemDropId].remove(h)
	case PlayerMoveHandler:
		return v.handlers[PlayerMoveId].remove(h)
	case PlayerJumpHandler:
		return v.handlers[PlayerJumpId].remove(h)
	case PlayerTeleportHandler:
		return v.handlers[PlayerTeleportId].remove(h)
	case PlayerChangeWorldHandler:
		return v.handlers[PlayerChangeWorldId].remove(h)
	case PlayerToggleSprintHandler:
		return v.handlers[PlayerToggleSprintId].remove(h)
	case PlayerToggleSneakHandler:
		return v.handlers[PlayerToggleSneakId].remove(h)
	case PlayerCommandExecutionHandler:
		return v.handlers[PlayerCommandExecutionId].remove(h)
	case PlayerTransferHandler:
		return v.handlers[PlayerTransferId].remove(h)
	case PlayerChatHandler:
		return v.handlers[PlayerChatId].remove(h)
	case PlayerSkinChangeHandler:
		return v.handlers[PlayerSkinChangeId].remove(h)
	case PlayerStartBreakHandler:
		return v.handlers[PlayerStartBreakId].remove(h)
	case PlayerBlockBreakHandler:
		return v.handlers[PlayerBlockBreakId].remove(h)
	case PlayerBlockPlaceHandler:
		return v.handlers[PlayerBlockPlaceId].remove(h)
	case PlayerBlockPickHandler:
		return v.handlers[PlayerBlockPickId].remove(h)
	case PlayerSignEditHandler:
		return v.handlers[PlayerSignEditId].remove(h)
	case PlayerItemPickupHandler:
		return v.handlers[PlayerItemPickupId].remove(h)
	case PlayerItemUseHandler:
		return v.handlers[PlayerItemUseId].remove(h)
	case PlayerItemUseOnBlockHandler:
		return v.handlers[PlayerItemUseOnBlockId].remove(h)
	case PlayerItemUseOnEntityHandler:
		return v.handlers[PlayerItemUseOnEntityId].remove(h)
	case PlayerItemConsumeHandler:
		return v.handlers[PlayerItemConsumeId].remove(h)
	case PlayerItemDamageHandler:
		return v.handlers[PlayerItemDamageId].remove(h)
	case PlayerAttackEntityHandler:
		return v.handlers[PlayerAttackEntityId].remove(h)
	case PlayerExperienceGainHandler:
		return v.handlers[PlayerExperienceGainId].remove(h)
	case PlayerPunchAirHandler:
		return v.handlers[PlayerPunchAirId].remove(h)
	case PlayerHurtHandler:
		return v.handlers[PlayerHurtId].remove(h)
	case PlayerHealHandler:
		return v.handlers[PlayerHealId].remove(h)
	case PlayerFoodLossHandler:
		return v.handlers[PlayerFoodLossId].remove(h)
	case PlayerDeathHandler:
		return v.handlers[PlayerDeathId].remove(h)
	case PlayerRespawnHandler:
		return v.handlers[PlayerRespawnId].remove(h)
	case PlayerQuitHandler:
		return v.handlers[PlayerQuitId].remove(h)
	default:
		return errors.New("invalid handler")
	}
}
