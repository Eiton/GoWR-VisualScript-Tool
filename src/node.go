package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

const MAX_NODE_LEVEL = 50

type Node interface {
	To_String() string
}

// 787d5f8620820850, 033390A1708C4402
type Node_Header struct {
	name    string
	node    Node
	nodeMap map[uint16]uint32
}

type Node_Variable struct {
	str string
}

type Node_Dummy struct {
}

// 339749F41CDA6CF3
type Node_SetEffectsSubtitle struct{}

// 93B1E21ECE378ACB
type Node_ShowUIItemCardMessage struct{}

// DAE225FB7D050062
type Node_ShowUIMessage struct{}

// C041E37D26A4C555
type Node_ShowUISidebarMessage struct{}

// EC90133CE7D29A2E
type Node_ShowUISplashScreenMessage struct{}

// 52C0B11E31001F7E
type Node_ContextAction struct{}

// A815ADC12F455A3F
type Node_Ceil struct{}

// 6415879C6D60BE1F
type Node_HasInteractZoneTag struct{}

// 1FCDE7BD7CAB4A12
type Node_AsLootDistributor struct{}

// 416FCD56CBF9CB08
type Node_IsPinType struct{}

// 73D5138B1D10A806
type Node_GetCameraTargetPosition struct{}

// F99BA6CE08A2E503
type Node_GetActiveMovePercent struct{}

// CBE7681D2D081500
type Node_AsBanter struct{}

// 2A6DB82686190E01
type Node_InteractZoneTemplate struct{}

// BB587847BD5C2401
type Node_GetScale struct{}

// 20083C263AAB8E02
type Node_GetBlackboardVariable struct{}

// 178838A79AFC1504
type Node_Icon struct{}

// 751884C159741E04
type Node_GetFlawlessAresAvailable struct{}

// 979C59704EF99404
type Node_IsInNewGamePlus struct{}

// EDC78D0A7699A205
type Node_GetSkillTreeTokensOfTypeOn struct{}

// E87A372A7C499707
type Node_CanCraftRecipe struct{}

// FBD721EE0B2BDA06
type Node_IsVolumeEnabled struct{}

// BB362CCC34741507
type Node_WalletHasEquipment struct{}

// 7F4BDEBE43771907
type Node_IsSideQuestInAnotherRealm struct{}

// 4C25C63B98998607
type Node_CheckProgressionFact struct{}

// 2E1B42EF540D9807
type Node_WwiseRTPC struct{}

// 328187E7E557E707
type Node_VectorDot struct{}

// 172E5CE9A06E6008
type Node_GetVectorFromBlackboard struct{}

// 0266DCBCB18CCB08
type Node_InstancedTraversePath struct{}

// 55507B2CFDF28A0D
type Node_GetLastAttacked struct{}

// 1401E22515CB8109
type Node_InstancedContextAction struct{}

// F60CC1933C77D408
type Node_GetClonedWeapons struct{}

// B4401A111A7EE308
type Node_GetContextActionApproachPoint struct{}

// 8A517611D6CB4009
type Node_IsAnySoundPlaying struct{}

// 599B0A794FE05309
type Node_IsFalling struct{}

// 9CFAF828B7AB9A0A
type Node_GetPlayersCurrentRegion struct{}

// DA7EB07E8079DA0A
type Node_IsControlDown struct{}

// 945A99A6DFFC5C0B
type Node_GetSeekTargetRemainingCount struct{}

// 579FEEFB38B0D60B
type Node_IsInsideFocalZone struct{}

// B106A6654350800F
type Node_QueryEquipmentTrait struct{}

// FB33A2532545D70D
type Node_GetAnimJoint struct{}

// 7E6943A6AFEDC40E
type Node_IsHostile struct{}

// 7013A1BADB00C80E
type Node_VectorCross struct{}

// 3424C43128A3EF0E
type Node_GetMeterName struct{}

// 3CA708DA86781810
type Node_FindRandomPositionOnNavMesh struct{}

// 6521D1F12CF09B10
type Node_AsMeter struct{}

// C3C6BA077B56C710
type Node_VectorAdd struct{}

// 3A50F9AE60CF1C11
type Node_QuestHasAnyFlags struct{}

// 21BBBA62C3816017
type Node_CheckFeatureData struct{}

// 368A6F5DB229CE14
type Node_GetStickInput struct{}

// 04E1E41094515D13
type Node_IsInteractZonePromptEnabled struct{}

// 1885624E8D82B412
type Node_CreatureBehaviorSetOptionData struct{}

// 0EFE86250CACCF12
type Node_Effect struct{}

// 6D4996EE826AEF12
type Node_Subtract struct {
	params [2]Node
}

// A16AD5E6F6402613
type Node_PlayerEitherHasOrCanCraftEquipmentRightNow struct{}

// A0A140DEDA0B6313
type Node_AsPlayFX struct{}

// C8A240163665F613
type Node_GetShieldValue struct{}

// 965120C0AFB71814
type Node_Equipment struct{}

// A4D12B9B6FD28714
type Node_IsUnoccupied struct{}

// A12BF9D5B5BA6C15
type Node_GetStringHashValueFromElement struct{}

// 854A33EA70E6EA14
type Node_IsCollisionEnabled struct{}

// BD53422167DE0B15
type Node_IsAttackableArrow struct{}

// 836EA18B55D51715
type Node_AsBranch struct{}

// FE4A56C4F82F5415
type Node_GetAnimationLength struct{}

// BACD3F39DCD98E15
type Node_IsMovementLimitedByCurrentAOO struct{}

// 5F812DF89D2B2E16
type Node_GetFloatValueFromElement struct{}

// 4610D0BAC8516A16
type Node_GetAvailabilityStateBySource struct{}

// CE06F1CE3BC99416
type Node_GetAllInteractsDisabled struct{}

// C141BC765096711D
type Node_ACos struct{}

// 42D173CFEBD58B1A
type Node_AsPickupSlot struct{}

// 99766FC467773A19
type Node_IsContainedInArray struct{}

// A93F370F32B35519
type Node_MakeVector struct{}

// 8D6670BB9249C119
type Node_FullScreenEffect struct{}

// 6F14F08AB5B9EC19
type Node_AsWwiseEvent struct {
	param Node
}

// FE2E85C681F4721B
type Node_ASin struct{}

// C659BC3FC5998A1B
type Node_GetActiveWeapons struct{}

// DE24398CE73F781C
type Node_AsWallet struct{}

// A62DE0C1AEAF3E1D
type Node_ArrowEmitter struct{}

// 54D0297AABD2E71E
type Node_GetOptionControlUpForSetting struct{}

// C9B10F8369EA261E
type Node_WasWeaponReflected struct{}

// 8DD1EC5A6C75541E
type Node_GetAnimDriverValue struct{}

// 036E13BEA402A71E
type Node_GetDockTargetType struct{}

// 74C487E17719C31E
type Node_IsDynamicCharacterSpa struct{}

// 62FA573E079D0E1F
type Node_EncounterElementData struct{}

// 768BDC18FCF91D1F
type Node_BeamAttack struct{}

// B08AB75C16F69E1F
type Node_Rand struct{}

// 4C9EAB4ACBE1CE2F
type Node_IsMentalStateEqual struct{}

// E1438E33E5989F28
type Node_IsSequenceSkipping struct{}

// 50C0D28308EF2124
type Node_GetEquipmentInWallet struct{}

// B304B794456B0322
type Node_GetRegionForMapMarker struct{}

// 27EEAAE00585A120
type Node_GetTyr struct{}

// 3E12DE27B934C920
type Node_HasResurrectionStoneBeenUsed struct{}

// 2858D742F7EC5E21
type Node_MakeVector4 struct{}

// 2508E084C4C1A621
type Node_GetCameraOrbitLeft struct{}

// FBA91B0E60677C22
type Node_GetVFSInt struct{}

// 53EF795BE8954523
type Node_IsValidEquipment struct{}

// 92D8229CA2E56823
type Node_IsMarkerClaimed struct{}

// AF7B822BADB36D23
type Node_GetTouchpadSwipeMapped struct{}

// F546D53574C7D824
type Node_RecipeHasFlag struct{}

// 2747BC2296105E24
type Node_GetBreakableStage struct{}

// 4F01A73671379024
type Node_IsTouchpadSwipeMapped struct{}

// 33E132C52E35C624
type Node_TraverseLink struct{}

// D7D95157B103D824
type Node_HasHitFlag struct{}

// 67CF4AEA90A29126
type Node_CanEquipToEquipmentSlot struct{}

// EA9343A22E819927
type Node_ChooseByCondition struct{}

// 3733EAFEDB939828
type Node_Enum struct{}

// B3ECE0E6AA649E28
type Node_IsSvrContextConnected struct{}

// 8B06A6B7BB50B12B
type Node_IsTraverseLinkEnabled struct{}

// 0492A60321D4502A
type Node_GetEquipmentInCharacterSlot struct{}

// E1EB3646F8125129
type Node_GetListenerPositionInfo struct{}

// D5729238CE147829
type Node_LootCondition struct{}

// 3019B268A85CDD29
type Node_MapMarkerHasAllFlags struct{}

// C8D2110D854D172A
type Node_IsWeatherCategoryActive struct{}

// 60D02BF79E7A562A
type Node_Max struct{}

// 64E850E126946F2A
type Node_GetContextActionFromBlackboard struct{}

// F365BC694CD33C2B
type Node_GetWolfSledInertia struct{}

// E2A49D74E9AF732B
type Node_AsConcussion struct{}

// 7A2E56CF1416602D
type Node_IsPlatform struct{}

// D24FAD9AC019C32B
type Node_AsNavCurve struct{}

// 737E8ACFD9D8FB2B
type Node_IsRealmLocked struct{}

// FDA150728F35972C
type Node_GetCameraOrbit struct{}

// B8C25393BE0B212D
type Node_AsWwiseSwitchGroup struct{}

// 4FDD3EF36DD8402E
type Node_AsSoundEmitter struct{}

// 0FEE9E195929AF2E
type Node_GetPlayersCurrentArea struct{}

// EE8370865260642F
type Node_GetScriptOwningWad struct{}

// F9A854988056AE37
type Node_Crank_GetCurrentC struct{}

// 1C8ED91C2E90C233
type Node_GetAvailabilityState struct{}

// DA6EFAEFA5BA8B32
type Node_IsRouteToOtherRealmsViaFastTravel struct{}

// 590600725C75FB2F
type Node_SortIterator struct{}

// 0E2A52DDE8961E30
type Node_AsInteractZoneTemplate struct{}

// 2D5B2216780C8D30
type Node_GetTimerFromBlackboard struct{}

// 430644D86593D131
type Node_GatewayMarker struct{}

// 090EEFC2F2FE1633
type Node_HasPositionalFlag struct{}

// 10B8A1D632732533
type Node_VectorLength struct {
	vector Node
}

// 96AC2853E2FF5F33
type Node_AreNavCurvesEnabledForObjects struct{}

// C88C515AC2B99433
type Node_GetCompassTargetPositionLerpSpeed struct{}

// 8B7837ECFBDE4B35
type Node_HitFlagsTo_String struct{}

// 359F14D86ADB4834
type Node_Add struct {
	params [2]Node
}

// 0E91834533D78A34
type Node_MemoryAvailableToSpawnCreature struct{}

// 5E5476DACFA62735
type Node_HasAnimationComponent struct{}

// F5F3D48E34812835
type Node_GetVelocity struct{}

// F81FEE9567519E35
type Node_Branch struct{}

// 7D9B9F4AC84B2A36
type Node_IsPlayGoDownloaded struct{}

// 629DD5298A06A036
type Node_ChooseByInt struct{}

// B951B756C2B29437
type Node_GetKratos struct{}

// 16CDEBE8FBC38F3C
type Node_IsPickupSlotOnCooldown struct{}

// 62F04BFDB01FAB3A
type Node_GetCompanionReticleHitData struct{}

// 832B6E7F4A3BCC37
type Node_GetCurrentOptionIndexForSetting struct{}

// F926409DB8A42838
type Node_AreBehaviorsEnabled struct{}

// 5F7A24CB8F4A8738
type Node_GetEquipmentInWalletWithFlags struct{}

// B240B78A528DBD38
type Node_Sin struct{}

// 296D5FA5EDCFD43A
type Node_GetBreakableHitPoints struct{}

// 389529A0B866E73A
type Node_GetGameObjectFromBlackboard struct{}

// 4B05B1DF22693A3B
type Node_GetGroundLevel struct{}

// EA26BA1CBCCB583C
type Node_GetVisibility struct{}

// 72226F5F3F62AA3D
type Node_RecipeHasFlags struct{}

// B0808B8786A0F13C
type Node_GetPathDirection struct{}

// A6B7FF47FDAC773D
type Node_GetBestNavmeshPosInAOO struct{}

// F88EF926B1E7993D
type Node_GetFightPosition struct{}

// 82F5BE7ED96C9F3D
type Node_GetTimeSinceBanterPlayed struct{}

// 5F06BFA64EA0E23D
type Node_ShuffleIterator struct{}

// 53C1A8302A92AE3E
type Node_IsButtonDown struct{}

// 5F47207ACCE6C23E
type Node_AreContextActionsEnabled struct{}

// 6F360CF22AD4A461
type Node_CreatureHasAOO struct{}

// CD17C3E5A0EEDA50
type Node_AsCreature struct{}

// 02CEEBE3D95E5A46
type Node_GetStringLength struct{}

// 06559D1FE9AA1E41
type Node_LootConditionSet struct{}

// 57D2EFC714904540
type Node_SoundPortal struct{}

// 47758251DEC08E3F
type Node_GetClosestFastTravelMarkerPos struct{}

// 7F27A7F5A1AEEE3F
type Node_GetQuestName struct{}

// 1AA9354D22941A40
type Node_Weapon struct{}

// 4E70A52DF80D1F40
type Node_AsEquipment struct{}

// B83A779AA5CEC540
type Node_GetWadName struct{}

// D039DA99EC66EC40
type Node_GetSetting struct{}

// E7BC749F62140F41
type Node_GetSpawnFacing struct{}

// 90F135F0B3731241
type Node_Or struct {
	params [2]Node
}

// A496A11065C04B44
type Node_GetPickupStageCount struct{}

// F316CC9E82359341
type Node_ScreenShake struct{}

// 4F5A9564A3073E43
type Node_IsPickupSlotActive struct{}

// E3DA123C99CC3F43
type Node_GetChildren struct{}

// DD8BEDFA5C28A043
type Node_MapRealm struct{}

// 2524992DB41E8F44
type Node_CompassPathDistance struct{}

// 909250A5DBAFA444
type Node_InteractZoneLocked struct{}

// 01DE231A03FF3246
type Node_WwiseSwitchState struct{}

// 54A2056ABB5E5146
type Node_AsMapArea struct{}

// D6760CAB6E13744C
type Node_IsPositionOnScreen struct{}

// 8730F1E200545C49
type Node_EquipmentHasFlags struct{}

// 3B6175BB1E297F46
type Node_GetPlayerBoat struct{}

// E0A740A87F7F2C47
type Node_GetQuestProgressAndGoal struct{}

// 7DF7E88CA9F83F47
type Node_GetObjectsWithFlagsInRadius struct{}

// 8FDD6E92AD8F2F49
type Node_GetMotionDebuggerRecordLocation struct{}

// 3F19FD355521B249
type Node_GetVariable struct {
	name string
}

// 29AA1F002C21BD49
type Node_AsAnimation struct{}

// 7A22EB2D80162B4A
type Node_GetVFSFloat struct{}

// 6EB6D1128093194B
type Node_GetCreatureGroun struct{}

// 5DC16FB7F65C0750
type Node_Numeric struct {
	val float32
}

// 41B277201D11AA4C
type Node_IterateFixedSizeArray struct{}

// B912D0ADB49B364D
type Node_GetCurrentOptionJoysticForSetting struct{}

// 285649A3EF25BD4E
type Node_GetExistingResourcesInWallet struct{}

// CEEB0BD5841CE64F
type Node_GetSpawnPosition struct{}

// 1E533D3CC3573550
type Node_GetBoolValueFromElement struct{}

// FA79944964216550
type Node_GetTargetCreature struct{}

// 3188431D706FCF50
type Node_AsQuest struct{}

// 25725C0309FA4C57
type Node_GetAngleBetweenVectors struct{}

// DFDBD89284F95D53
type Node_IsNavAssistCompassMarker struct{}

// F04679388D8E5B52
type Node_IsPlayerDebugMovement struct{}

// 9F51B67FA6E3DF50
type Node_DidCreatureRecentlyHaveFlag struct{}

// 64F5CCB6B5E23351
type Node_GetPendulumAngle struct{}

// 7EB1931EF4B98151
type Node_Prop struct{}

// 38FF1281E4B0B151
type Node_AsSoundProximityTrigger struct{}

// 296D624319FE9752
type Node_GetMaxSpeedOverride struct{}

// 817F4A9FA9E9B252
type Node_IsInValhalla struct{}

// 81613FB023BFCA52
type Node_GetWeaponGameObject struct{}

// 8ACA12277C431753
type Node_GameObject struct{}

// 58C352A97FFA1C55
type Node_GetCurrentSpline struct{}

// 53370E89D3090C54
type Node_GetAnimationTime struct{}

// F5516FD8F533A654
type Node_AsMapRealm struct{}

// E68599F5C63BC254
type Node_IsSummaryCategoryInMapRegion struct{}

// 67A9E17EF62DFE54
type Node_GetMapMarkersInArea struct{}

// 3A7FDFDF1BE57355
type Node_GetCollisionRadius struct{}

// EEA72C87452D6156
type Node_GetMPIconObjectByName struct{}

// 0D9F978FC8B76B56
type Node_GetName struct{}

// E207026532A1FC56
type Node_InstancedNavCurve struct{}

// 424A385421D7B85B
type Node_AsCamera struct{}

// 7BEED5AF1D6C6458
type Node_GetWolfSledHarnessVisibility struct{}

// 68D7787552A5C957
type Node_HasBlackboardVariable struct {
	params [2]Node
	param  uint16
}

// 48B6B23A1686D157
type Node_IsPickupSlotFree struct{}

// D7A644434EDCDE57
type Node_GetAllSoundEmittersOnObject struct{}

// 6A0FDC973A484F58
type Node_AsTweakConstant struct{}

// BD75554DAE3D155A
type Node_InstancedInteractZone struct{}

// 52753D5851D03E5A
type Node_IsInsideVolume struct{}

// 0E946842DE6BEB5A
type Node_GetCompanion struct{}

// 7CCFE6321089EB5A
type Node_GetCreatureUnderReticle struct{}

// 6B2C1C2EBD44DA5D
type Node_GetQuestState struct{}

// DA57C4E5FC9EFA5C
type Node_VectorScale struct{}

// 66F2867B59960D5D
type Node_VectorSubtract struct{}

// 0C46EAEC68BF6D5D
type Node_ArePhysicsEnabled struct{}

// 9B9D492BD58C865D
type Node_GetArrowOwner struct{}

// E0DDB57F9A80035E
type Node_GetCurrentOptionControlDownForSetting struct{}

// 5F946D06C162CD5E
type Node_ContextActionSequencerStage struct{}

// 548DB112F4929960
type Node_GetPendulumSpeed struct{}

// 6E7A5956333D8172
type Node_IsValidReference struct{}

// ECEC3A6FFED31B6A
type Node_GetBranchValueFromElement struct{}

// D0B482CC18289467
type Node_IsTargetObstructed struct{}

// 31EBD975352DE763
type Node_AsRecipe struct{}

// CA5D71C1AEF52862
type Node_IsPickupActive struct{}

// 3FA7C4E1990ED462
type Node_GetPlayerDeathsOnCurrentCheckpoint struct{}

// DB2571EF4F5E3663
type Node_CastToInt32 struct {
	param Node
}

// 73FEE5BB22D38E63
type Node_GetWolfSledRopeVisibility struct{}

// 5172760831106664
type Node_CanAim struct{}

// DF47AA8D47943565
type Node_Floor struct {
	param Node
}

// 371BB06DBB9F1466
type Node_BitwiseAnd struct{}

// 05340FB9C978D166
type Node_GetAllRecipesInWallet struct{}

// D89AD93F11DC6068
type Node_SoundEmitter struct{}

// 91112EB0D18C9967
type Node_ContextActionSequencerStageData struct{}

// 4C3B8555AF109E67
type Node_GetWeaponThrowStatus struct{}

// F82D881709B01368
type Node_AsWwiseSwitchState struct{}

// E8FB5449C6845868
type Node_IsPlayingMove struct {
	params [2]Node
}

// 53A03EE076770569
type Node_GetMappedButton struct{}

// D3C9C0F0B6665A69
type Node_IsWeatherSystemActive struct{}

// DC94522861C46769
type Node_MapMarkerHasAnyFlags struct{}

// A167F4F521BACD69
type Node_HasBeenUsed struct{}

// BA395A6FAE49776C
type Node_IsOnNavMesh struct{}

// F511929C12B9606B
type Node_GetUniqueEquipmentByName struct{}

// B4A03C1D27573B6A
type Node_Sqrt struct{}

// 87F2A2ED0F07936A
type Node_IsAnyBanterPlaying struct{}

// D71817507057166B
type Node_GetAnimationFrame struct{}

// 48691295C08A3E6B
type Node_GetRandomPositionInVolume struct{}

// 2AB7D87BF8CA786B
type Node_GetSplinePosition struct{}

// 4DDEDF5ECAE7F36B
type Node_EncounterWaveData struct{}

// 2A0829F19FCC146C
type Node_GetCameraOrbitForward struct{}

// 01F0BD04F14F196C
type Node_Bool struct {
	val bool
}

// BF9FC5F418875D70
type Node_Abs struct{}

// E0E82EC16CF5B56D
type Node_GetSpawnedObjectVariable struct{}

// FEBD20D79F2CB86E
type Node_GetTweakFloatConstant struct {
	params [2]Node
}

// 5D01376A7A7F866F
type Node_GetRegionForWad struct{}

// 3261F73DA77F4270
type Node_AsWildlife struct{}

// 52757504C8607C70
type Node_AsPickup struct{}

// 17E1652A4CAE2071
type Node_AsAnimJoint struct{}

// 66AF364938109171
type Node_IsPlayerOnWolfSled struct{}

// F3DD8834B3F29777
type Node_Wallet struct{}

// 2F5543725EFB2F75
type Node_QuestHasAllFlags struct{}

// 0B0F5B516621CB73
type Node_GetDepthOfField struct{}

// DF387335D5B2A072
type Node_GetFriendlyCreatures struct{}

// D87B24DEAFE7D772
type Node_AsColor struct{}

// 600A509FDEFD2473
type Node_IsMarkerLocked struct{}

// 62C6D8B2E4486973
type Node_Divide struct {
	params [2]Node
}

// B183BC7D15D37074
type Node_PointOnPath struct{}

// 866D92FD752FDE74
type Node_GetBitFromVariable struct{}

// F8DA23BAFF35F074
type Node_IsHighlightDisableOverrideActive struct{}

// DB6C47BB54310C75
type Node_AsProp struct{}

// 379986296E5A0776
type Node_Pickup struct {
	name          string
	nameHash      uint64
	namespace     string
	namespaceHash uint64
}

// 65941400E52C4E75
type Node_Vector struct {
	x float32
	y float32
	z float32
}

// 713BB0F002E84E75
type Node_AsRumble struct{}

// 3B7CACFA2AB58E75
type Node_GetIntFromBlackboard struct{}

// BC9F575C8C84B175
type Node_IsAreaOfOperationEnabled struct{}

// 71688B18653EB276
type Node_IsPickupAvailable struct{}

// 99733C6F0EC3CC76
type Node_AsEnum struct{}

// E2E68DA846182377
type Node_GetLootObjectResult struct{}

// 399B421B678F8877
type Node_IsWeaponEmbedded struct{}

// FB29C935F9425E7A
type Node_GetLootConditionByName struct{}

// 6C2E3D80AF897379
type Node_GetAngrboda struct{}

// 1B9B09F8AD74B777
type Node_InstancedSoundEmitter struct{}

// BCC4F6E9DCD1C678
type Node_AsGatewayMarker struct{}

// E693829E4EB7FF78
type Node_IsAOOAssignmentTypeEqualTo struct{}

// 3B9ED4FD51822A79
type Node_GetCurrentContextAction struct{}

// DF980BF9BB3D9179
type Node_IsSpawnEnemiesEnabled struct{}

// 49E297A263D60C7A
type Node_ATan struct{}

// C81FFF2816BB227A
type Node_CastToFloat struct {
	param Node
}

// 01C63D2F87242E7A
type Node_GetClosestPositionInVolume struct{}

// 161D1ED9AC60C67C
type Node_IsDoingSyncMove struct{}

// FA326E83B480777A
type Node_ShouldPointToFastTravelMarker struct{}

// 7F107D84E6CA837A
type Node_AreContextActionsEnabledForObjects struct{}

// 502257241BE8C17A
type Node_FindPath struct{}

// D3AD0C02C6BD677B
type Node_GetAtreus struct{}

// 9151FC57AA7FDB7C
type Node_CastToUInt64 struct {
	param Node
}

// 0CFBFEBAF7DEE47C
type Node_GetPlayerLockOnTarget struct{}

// 88752874E4D1D07D
type Node_IsPlayerInBoat struct{}

// 7F61F9A8EFB262C1
type Node_Min struct{}

// 9E31788953E2759F
type Node_WasEncounterRunning struct{}

// 229D87CD9FECB28F
type Node_CanEquipToCharacterSlot struct{}

// 7456B1DC7113F583
type Node_GetCreaturesInRadius struct {
	params [2]Node
}

// 7B7EA83A406A8781
type Node_GetAllSkillsInTree struct{}

// F8C8E79ACA29BB7F
type Node_IsWeaponActive struct {
	params [2]Node
}

// 9624AED0F6312C7E
type Node_GetCurrentVehicleCreature struct{}

// 42146239F457CE7E
type Node_AsFullScreenEffect struct{}

// 688A0263E7A5267F
type Node_GetMapMarkersInRegion struct{}

// D565F65DB8F09C7F
type Node_GetClosestFastTravelMarker struct{}

// 1E0C7BA7D5562380
type Node_GetGIBlend struct{}

// 18ACEFF4DF05B080
type Node_GetSideQuestRealm struct{}

// 43044349DCC42F81
type Node_IsJumping struct{}

// D17577A116198381
type Node_IsWolfSledSpawned struct{}

// 886AC326F68F9082
type Node_AsVector struct{}

// B7BEE12B6609C681
type Node_GetFirstEquippableEquipmentSlot struct{}

// 4DA747B96E332782
type Node_GetEquipmentGeneratorFromHandle struct{}

// 0752554762683C82
type Node_GetLookAtTarget struct{}

// 6A7B1B6023C45282
type Node_IsPlayerInsideLeashZone struct{}

// 73252C8B0585B682
type Node_GetNavCurveLength struct{}

// 6E7476D289B8DC82
type Node_AsGameObject struct{}

// 3CD62D69A2EC7383
type Node_Rumble struct{}

// A2153E07B938BF83
type Node_IsAnyOtherSoundPlaying struct{}

// EA49A8F3CFF5F78A
type Node_CastToBool struct {
	param Node
}

// F45CEC5CEC084788
type Node_IsAvailabilitySourceRequesting struct{}

// 85001198E0113284
type Node_WalletHasRecipe struct{}

// 2665688A6083BD85
type Node_BreakVector struct{}

// 0C9DFAA307379986
type Node_GetMainQuestRealm struct{}

// EA52A31ED9223087
type Node_GetVFSTokenFromPath struct{}

// D4B0603E13F26388
type Node_GetButtonIndex struct{}

// 86154A1358F01389
type Node_IsModelEnabled struct{}

// EC7E4E1F46C54289
type Node_GetOptionJoysticForSetting struct{}

// 313283CDD133948A
type Node_GetPickupStage struct{}

// 0D82A05038A2148E
type Node_GetWolfSledIsInDrift struct{}

// 6871ED3EF297578B
type Node_GetOptionControlDownForSetting struct{}

// 0C818064BBE55C8C
type Node_GetBooleanFromBlackboard struct{}

// D808964ED96BF48C
type Node_AsSoundPortal struct{}

// 55055BAA233BF98D
type Node_CameraRecenter struct{}

// EE053C4B3B89168E
type Node_And struct {
	params [2]Node
}

// FDF5A13C60A8D08E
type Node_CreatureOptionKeyValuePairData struct{}

// 18D62C0E1139368F
type Node_AreLightsEnabled struct{}

// 55B391042B8F9199
type Node_StringHash struct {
	str  string
	hash uint64
}

// 7496C50263153294
type Node_Round struct{}

// 6F95E7B322A76492
type Node_ResourceHasFlags struct{}

// E6618B4E12E2FA8F
type Node_GetCameraRender struct{}

// 2A50EEC9FC722C90
type Node_AsLootCondition struct{}

// 426F8EEDDB208790
type Node_GetVFSBool struct{}

// 1ED5AA8BD2EA2B91
type Node_AsMapRegion struct{}

// C4DF3851C8388992
type Node_IsInteracting struct{}

// 92A9AA86DD4AF292
type Node_IsPositionInCreatureAOO struct{}

// ADEC9B8C3BCF7593
type Node_IsInContextAction struct{}

// 7D9AB7B48DA3E593
type Node_MapAreaHasAnyFlags struct{}

// 832146796071F695
type Node_GetWolfSummoningPoints struct{}

// A70EBA6B08AEC094
type Node_BitwiseOr struct{}

// A1FD87CAE6B41695
type Node_IsSideQuestPathfindingNonContiguous struct{}

// DC56F77D99994895
type Node_GetMapRegionsInRealm struct{}

// C83F0A500667AF95
type Node_GetElement struct{}

// 49B4DE5874EE4496
type Node_Clamp struct{}

// 5DD2C814B4B10E98
type Node_IsDying struct{}

// 1C6EAE6464D5FB98
type Node_Float struct {
	val float32
}

// CCAE45D57C4D3B99
type Node_GetSkillTreeTokenIn struct{}

// 717948816C00289D
type Node_GetMusicIntensity struct{}

// 0E07073B5B47379B
type Node_CanAcquireSkill struct{}

// BF4420A11121B999
type Node_GetCreatureByName struct{}

// BDE036C40307D699
type Node_GetEnemiesInRadius struct{}

// AA1CEBFAF632069A
type Node_NavCurve struct{}

// DF0A4CA6DED6019B
type Node_AreNavCurvesEnabled struct{}

// 0E722736994A389B
type Node_GetViewport struct{}

// B396F09C2F24479B
type Node_PartFlagsTo_String struct{}

// C63D04A31890649C
type Node_Resource struct{}

// D3CCB41E23B5119D
type Node_GetSplineProgression struct{}

// BD9E097279CC6C9E
type Node_CreatureHasBehaviorTree struct{}

// DA7DC9B7184C5D9D
type Node_GetVariableInfo struct{}

// 63DF4CBFE1C47C9D
type Node_Wildlife struct{}

// 5E8E6105EC5D889D
type Node_IsWadLoaded struct{}

// A9244278D9E9AC9D
type Node_HashString struct{}

// 281472EF290AB89E
type Node_CreatureOptionData struct{}

// BF74532C78553B9F
type Node_IsInsideLeashZone struct{}

// C405E1E70CEB629F
type Node_GameMap_IsItemStateV2 struct{}

// F2609AA50A0D81AF
type Node_GetMaterialEntryValue struct{}

// 8358A69A1957F3A4
type Node_GetAllInteractZones struct{}

// B086AE0D1C053AA3
type Node_IsInPlaytest struct{}

// 477EE3F0251656A1
type Node_GetYakWaterHeight struct{}

// C37E656C704E42A0
type Node_GetFirstChild struct{}

// 330620E6553D5AA0
type Node_GetIntForContextActionEnum struct{}

// AFEFF0AEE067D5A0
type Node_Haptic struct{}

// DF830E2FF3BA38A1
type Node_MapRegion struct{}

// AD105E0C2CCE0BA2
type Node_GetBanterDuration struct{}

// 620B22C1F08347A2
type Node_GetControlIndex struct{}

// D7DD89E55389A0A2
type Node_Recipe struct{}

// 34C3C285D8C5DAA2
type Node_ATan2 struct{}

// AD17799B143419A4
type Node_IsBreakableBroken struct{}

// 8772407649B990A3
type Node_MapRegionHasAnyFlags struct{}

// 34438E5351AFB3A3
type Node_VectorDistance struct{}

// 2E0C717DF51309A4
type Node_InstancedSoundProximityTrigger struct{}

// D9E36B38CB4D0BA4
type Node_GetCurrentOptionControlUpForSetting struct{}

// DDF15EBFE9101EA4
type Node_GetAssociatedGameObject struct{}

// 1D501527749595A4
type Node_GreaterThan struct {
	params [2]Node
}

// 702750DD62D3BFA4
type Node_GetAllResourcesInWallet struct{}

// 55C6E968B78EE3A4
type Node_GetAngleToTarget struct{}

// 6159ABFC63F817A9
type Node_Multiply struct {
	params [2]Node
}

// DFE3C7BB430545A7
type Node_GetHitPoints struct{}

// 54DBD189813355A6
type Node_IsInteractCandidateSetActive struct{}

// 33475C6802A3CCA6
type Node_DebugLams struct{}

// 505178C3079000A7
type Node_AsDecision struct{}

// BF77EC68816803A7
type Node_GetCurrentOptionEventForSetting struct{}

// 1615318E67B54DA7
type Node_GetCurrentOptionEventModForSetting struct{}

// C6AFEDA51C3CB0A7
type Node_GetOptionEventForSetting struct{}

// C369570D1C63E2A7
type Node_String struct {
	str string
}

// 86C0E290F1AD4DA8
type Node_InteractZone struct{}

// 5FBABEBFC1A2E2AC
type Node_HasCombatSetFlag struct{}

// DFA4510CDE5355A9
type Node_HasMeter struct{}

// 6BB256AC5B70C9AA
type Node_AsLams struct{}

// E31CAACDEE14DBAB
type Node_GetObjectWithUniqueFlag struct{}

// AAFDC116F94E44AC
type Node_BehaviorTreeSubtree struct{}

// E5770991F2D626AD
type Node_Lams struct{}

// 9B8F2BB4456648AE
type Node_GetProgressionFact struct{}

// 1AC1E51081DF9FAE
type Node_GetPlayer struct{}

// 787A99E56BB03CB8
type Node_GetStringHashFromBlackboard struct{}

// 578B34D8B6E724B4
type Node_AsContextAction struct{}

// 7E2056F668B1A4B1
type Node_IsMonitorLookingAt struct{}

// F72F16207A6215B0
type Node_BitwiseNot struct{}

// D226FE538301C5B0
type Node_IsTraversePathEnabled struct{}

// 81462280D132D6B0
type Node_HasAnySaveGames struct{}

// F100FA5C837567B1
type Node_IsJointVisible struct{}

// E961E431F4ED29B3
type Node_AsFixedSizeArray struct{}

// 17E33942C0EE2EB3
type Node_GetSoundEmitterSplineTV struct{}

// C41E8A29CCFFB6B3
type Node_Meter struct{}

// 76D9EE85BEAE10B4
type Node_HasHitOrPartFlags struct{}

// B8F0233145DD7CB6
type Node_Not struct {
	param Node
}

// FC9C3531599C31B4
type Node_IsOnGround struct{}

// D0CBBCA6788843B4
type Node_GetRecipesInWalletWithFlags struct{}

// AD349BD0B7555FB5
type Node_AsBehaviorTreeRoot struct{}

// 93E78A7B43A101B6
type Node_GetArrayLength struct{}

// 4032A55A4D21EEB6
type Node_DoesCreatureHaveLookAtTarget struct{}

// 92A05E33F2D3B0B7
type Node_AsHaptic struct{}

// 2E8918D499F5B3B7
type Node_IsSoundPlaying struct{}

// AB65E4688479DEB7
type Node_GetPosition struct {
	param Node
}

// 0226303867D3F9BB
type Node_GetMeterMin struct{}

// EE4D32DA6322E8B9
type Node_GetGlobalVariable struct{}

// EEECB26DF1A255B8
type Node_IsGoldBuild struct{}

// 6456B4DFFFBF64B8
type Node_AsStringHash struct{}

// 3342095F3FD484B8
type Node_GetDistanceToTargetCapsule struct{}

// 770741C5BE8201B9
type Node_Animation struct{}

// 93BEC81F6320F9BA
type Node_AsArrow struct{}

// 2F87F985625358BB
type Node_IsEncounterRunning struct{}

// 87C06047A5C4AFBB
type Node_AsCameraRecenter struct{}

// 4911B8AB8B0DECBB
type Node_AreCreaturesOnSameTeam struct{}

// 63C122CC9AF800C0
type Node_GetResourceAmountInWallet struct{}

// 053FEDFE0AE084BC
type Node_AreTraverseGraphsEnabledForObjects struct{}

// F4B78004ECD8E3BC
type Node_HasFlag struct {
	params [2]Node
}

// A76B7D032A8E32BE
type Node_BreakVector4 struct{}

// 7880289BA5E9C8BF
type Node_Vector4 struct{}

// D189CE56A4830EC0
type Node_GetCallAndResponseObjectsInRadius struct{}

// F432A630B5C843C0
type Node_IsOnTraverseLink struct{}

// 46659A3A806247C0
type Node_MapArea struct{}

// 9B0CDE5C5E5799E3
type Node_GetMapAreasInRegion struct{}

// 39635B54512FEED1
type Node_EquipmentGeneratorHasFlags struct{}

// 495E1EB3063A82C9
type Node_GetPickupInSlot struct{}

// 95FCA6571219A9C5
type Node_RotationFromEuler struct{}

// 8BB6CF4ADF19FFC3
type Node_AsTraverseLink struct{}

// D17C2FA6806A87C1
type Node_GetSelectedCreature struct{}

// 4C55E233FC9124C2
type Node_PlayFX struct{}

// 7B15884DA37D2CC3
type Node_GetCreatureStatsValue struct{}

// FD68FEF59C136AC3
type Node_IsDead struct{}

// 99FF9FA945EE7EC4
type Node_Modulo struct{}

// B7CD252DC758D4C4
type Node_IsUnderSystemCon struct{}

// 76315154735FEEC4
type Node_LootDistributor struct{}

// 10EA2556B0F501C5
type Node_CreatureSpawnOptions struct{}

// 672B8FA3E239F8C6
type Node_IsMusicPlayingAnything struct{}

// 37EC670CE44437C6
type Node_GetPlayersCurrentRealm struct{}

// 04349256F7BF5CC6
type Node_IsBanterPlayingOnCharacter struct{}

// 37EEBE2B4FF27FC6
type Node_Array struct {
	length Node
	elems  []Node
}

// 2704A419565AD7C6
type Node_GetMeterMax struct{}

// D29B600A6F9187C7
type Node_AreTraverseLinksEnabledForObjects struct{}

// C37CA2B248ADC5C7
type Node_CanEquipTokenOfTypeTo struct{}

// D4A31B29754CEEC8
type Node_InstancedSoundPortal struct{}

// 1B3CABBE455981C9
type Node_GetFirstEquippableCharacterSlot struct{}

// 368558C4ADBC16CE
type Node_GetContextActionTraversalData struct{}

// 285B0213E0F582CB
type Node_Camera struct{}

// FB02F8C5E190EAC9
type Node_IsAvailableForFollow struct{}

// 69872491E73AF4C9
type Node_GetLootConditionSetByName struct{}

// B6A56E5C21D55CCA
type Node_IsCreature struct{}

// 9C5D660BC9DB7CCB
type Node_GetCreaturesSpeedSettingFor struct{}

// BDC90EF314FDA9CB
type Node_MapMarker struct{}

// CB3886EF25E415CC
type Node_GetNumAvailableGameObjects struct{}

// 3F1C46148B738BCC
type Node_IsAvailableForSplines struct{}

// 249CFF2245FDA1CD
type Node_AsIcon struct{}

// 3508216B54C220D1
type Node_IsMainQuestInAnotherRealm struct{}

// DAAD760A7E4AE3CE
type Node_GetClosestPositionOnNavMesh struct{}

// 178C304720B937CF
type Node_GetPointOnFX struct{}

// D2B6FB0F5611CACF
type Node_Color struct{}

// 4A9129D8D9082BD0
type Node_AsArrowEmitter struct{}

// 647E815E224B39D1
type Node_GetEnemyFromEncounter struct{}

// B49B90D88BEA54D1
type Node_GetTimeScale struct{}

// 1E12DB985CACC9D1
type Node_AsBitfield struct{}

// 4DEB6AFF92B2C4D9
type Node_IsNavObstacleEnabled struct{}

// 81B24747F7AC40D7
type Node_IsEncounterCompleted struct{}

// 2F4F2CDF82D589D5
type Node_GetCenterOfScreenEnemy struct{}

// AF0FD36574803CD2
type Node_InstancedTraverseLink struct{}

// C6B7F5822F0B9BD2
type Node_IsAnyWeatherSystemActive struct{}

// 97560C0AF399B4D3
type Node_GetParticleMonsterAliveTotemCount struct{}

// F40E9C2B8C4429D4
type Node_LootResultHasEntry struct{}

// B01451D181A4A6D5
type Node_TraversePath struct{}

// 6F75AB11FCBBB7D5
type Node_AsScreenShake struct{}

// EB250729C71019D6
type Node_IsFocalZoneLockInEnabled struct{}

// B79D93C8B88E19D6
type Node_SimpleStateMachine_GetState struct{}

// 03F9F3FDCA190FD8
type Node_LessThan struct {
	params [2]Node
}

// 046CD45CAE0644D7
type Node_GetPlayingBanter struct{}

// A93B70BB950485D7
type Node_IsFightKnowledgeEnabled struct{}

// 96D0FA1D7BAFBED7
type Node_MapRegionHasAllFlags struct{}

// 64BA34F93ACFFDD7
type Node_IsContextActionTraversalBehaviorInitialized struct{}

// A4CF9664A7D11ED8
type Node_WwiseEvent struct {
	name string
}

// 7D379EE135DF9AD9
type Node_IsSettingEnabled struct{}

// E33553F71FA59FD9
type Node_IsInNavigationMove struct{}

// 9FF0429CF84CBDD9
type Node_IsAvailableForCombat struct{}

// 907BA069DD6033E1
type Node_IsControlDisabled struct{}

// 5B9A8608258C96DC
type Node_GetLastAttacker struct{}

// 11641DAC004780DA
type Node_CanBanterPlay struct{}

// FDE9D25003E0E3DA
type Node_IsQuestInState struct{}

// B17A3F4EF1DF05DB
type Node_TweakConstant struct {
	name string
	val  uint64
}

// 8A72A8C41BB8A1DB
type Node_IsMainQuestPathfindingNonContiguous struct{}

// B21F68D00245B7DC
type Node_GetTweakIntegerConstant struct {
	params [2]Node
}

// 89F027CD83D021DE
type Node_AsWeapon struct{}

// 6BB77C3F1D0313E0
type Node_IsAvailableForBanter struct{}

// 237B2246A8DAFDE0
type Node_MapAreaHasAllFlags struct{}

// F9BFE3551C6692E2
type Node_GetFixedSizeArrayCapacity struct{}

// 74C2F79FF80CA1E1
type Node_InteractZoneEnabled struct{}

// 8DA19E6D6441D9E1
type Node_IsWithinAngle struct{}

// 1367F30432F642E2
type Node_PickupSlot struct{}

// BD65B3FCDD2C53E2
type Node_IsOnActiveTraversePath struct{}

// FFB95FFE2187D9E2
type Node_IsPaused struct{}

// 736B88FCE725EEE2
type Node_IsPickupAcquired struct {
	params  [3]Node
	params2 []Node
}

// 8C39C23A821601E3
type Node_AsWad struct{}

// 84F7EDDA93CF63F2
type Node_AnimJoint struct{}

// 57BDA2D4B0C09DE9
type Node_GetActiveWeatherSystems struct{}

// 5538B58B256474E7
type Node_GetQuest struct{}

// 4AC136F672A591E6
type Node_GetSpawnedObjectOwningObject struct{}

// D0321ADD1D0AFAE4
type Node_GetCurrentAnimation struct{}

// A3732CA01D2403E5
type Node_GetCombatStatus struct{}

// C4B873BA32340CE5
type Node_Tan struct{}

// 8F5C142799B7B2E5
type Node_GetWolfSledIsDriftAllowed struct{}

// 33676F753F6CC9E6
type Node_GetFloatFromBlackboard struct{}

// AA3CE495FE1BD8E6
type Node_GetResourcesInWalletWithFlags struct{}

// DC600FCEB385E1E6
type Node_GetCreatureAttributeValue struct{}

// D88F5C7C1F760EE7
type Node_Banter struct{}

// 4DDDF19C603A6AE8
type Node_GetNavCurveInformation struct{}

// E6B13578701776E7
type Node_FindRandomPositionOnNavMeshInRectangle struct{}

// 0E34D69ACD4DA5E7
type Node_GetEquipmentInEquipmentSlot struct{}

// CCF200B1F6E5F9E7
type Node_GetAnimationPlayRate struct{}

// C67859E53DA50CE8
type Node_GetValhallaComplete struct{}

// 36792104A9317EE8
type Node_GetMiniGameplaySkipped struct{}

// 0B2CC0C2456EB9E8
type Node_IsWaveRunning struct{}

// AC14AC78C1C85EE9
type Node_AsBeamAttack struct{}

// 020912693E7C73E9
type Node_IsAvailableForSync struct{}

// BECDF5899D7A22EF
type Node_GetNumPointsOnFX struct{}

// 1FA2A974C1A752EB
type Node_GetPickupSlotName struct{}

// 9356985B357FF9E9
type Node_IsParticleEmitterEnabled struct{}

// 8708331F910095EA
type Node_GetValuesFromStageData struct{}

// 4DF5397C12DCA3EA
type Node_Arrow struct{}

// A0CED481D2A9D6EA
type Node_AsVector4 struct{}

// F6CF4842280431EC
type Node_GetFlawlessZeusAvailable struct{}

// 9AB2DAEC074AD8EC
type Node_GetNearestSoundPortal struct{}

// B47E0FF794B39EED
type Node_IsMusicPlaying struct{}

// 0D4826EF61A9B4EE
type Node_BehaviorTreeRoot struct{}

// 5506D83EAD1E4DF0
type Node_GetIntValueFromElement struct{}

// 5DA75980D97930EF
type Node_AsWwiseRTPC struct{}

// 02C33FBD7D6A63EF
type Node_IsVisualScriptEnabled struct{}

// 3874D2177E6CE0EF
type Node_GetGameObjectValueFromElement struct{}

// 5EB00DEA065407F0
type Node_GetGameTime struct{}

// 303ABE3F2C0074F0
type Node_AsBehaviorTreeSubtree struct{}

// 60602ECBB26BC0F0
type Node_GetParent struct{}

// 9BB436CDE2E232F2
type Node_GetCreatureContext struct{}

// 27CBE85B9C93DEFA
type Node_Int32 struct {
	val int32
}

// C433C5E8EBB7CEF7
type Node_GetGameObjectWad struct{}

// 161F6BECFDA7EFF5
type Node_WwiseSwitchGroup struct{}

// 46395421ED4DEAF3
type Node_AsTraversePath struct{}

// 89E8E463C819CDF4
type Node_SoundProximityTrigger struct{}

// D59DCCEF0AE031F5
type Node_ChooseByValue struct{}

// 3AA614F52D3A53F5
type Node_GetNumBreakableStages struct{}

// 0899E3885F1A5EF6
type Node_GetSpawnElementUID struct{}

// 970195BB56316BF6
type Node_Wad struct{}

// 1CD81E378151C6F7
type Node_ConcatenateAndHashStrings struct{}

// 202C4ADD3B39CBF7
type Node_GetBitFromGlobalVariable struct{}

// 0F94C4CD796462F9
type Node_AsResource struct{}

// 0D3A2D374D0FF6F7
type Node_Creature struct{}

// C16C77586BDF35F8
type Node_GetOptionEventModForSetting struct{}

// 7059642CA7FB9FF8
type Node_AreParticlesEnabled struct{}

// 0E0B6301D1621BF9
type Node_Decision struct{}

// 2A48C16366F9DEF9
type Node_IsQuestTracked struct{}

// 1D75EDF5FBE304FA
type Node_GetSpawnedObjectHierarchyRoot struct{}

// 9638C8EE21BE49FA
type Node_IsSummaryCategoryInMapArea struct{}

// D8264A83A27FD2FA
type Node_VectorNormalize struct{}

// 5CFE5D60A01723FE
type Node_AreInteractZonesEnabledForObjects struct{}

// 13C70F875EAA77FC
type Node_UInt64 struct {
	val uint64
}

// EBFF6DA57D36E0FA
type Node_Cos struct{}

// 6929CF4286D1FCFA
type Node_GetMeterValue struct{}

// 72DF2EFD1FD1D2FB
type Node_AsLootConditionSet struct{}

// 63A1BB2AE5B8DFFB
type Node_Concussion struct{}

// 630D325A8182DFFC
type Node_GetQuestPrimaryCompletionFact struct{}

// 1057AC44F526E0FC
type Node_Equals struct {
	params [2]Node
}

// E27D8DA7F380F0FC
type Node_Quest struct{}

// 07F8BD2057B179FD
type Node_IsSkillAcquired struct{}

// 8EC3504D30015AFF
type Node_CheckProgressionFactGreaterEquals struct{}

// 6FD0959F311324FE
type Node_GetNameHash struct{}

// DF83E783F97B52FE
type Node_RotateVector struct{}

// 115D2C6DFA0D6AFE
type Node_GetRotation struct{}

// 157ADF5C41BB4BFF
type Node_AsMapMarker struct{}

// B843D9F7DFB461FF
type Node_IsInAir struct{}

// E63B6D13A85A6AFF
type Node_ResourceHasFlag struct{}

// 5540C13E096EE5FF
type Node_AsInteractZone struct{}

// DA5B10BB7B14EF07
type Node_ControlClothSim struct{}

// 7727A8CE67F30804
type Node_DrawPoint struct{}

// 3E251066BD213901
type Node_MarkResurrectionStoneUsed struct{}

// 71DDE16CD8DF3600
type Node_ClearWeatherCategory struct{}

// C3534134634A4F00
type Node_SetAnimationPlayRate struct{}

// CABC7C93B62E5200
type Node_SetDisableAllInteracts struct{}

// 78E02FAFE4467D00
type Node_AddExcludeTraverseLinkFilterToCreature struct{}

// C96987A11B93A400
type Node_PlayCombatConcussion struct{}

// DE783CC0462E3901
type Node_SetHapticInstanceParameter struct{}

// 62B858BE1916AC01
type Node_RequestCinematicModeEnabled struct{}

// BEE4A5313024E102
type Node_Sluice_StartWater struct{}

// 4F74958321C7F602
type Node_StopRotatingObjectTowardsTarget struct{}

// 7857067582334203
type Node_RegisterOnEncounterFinished struct{}

// 10C4F2AD364BB005
type Node_SetTraversePathFlagEnabled struct{}

// 387D17C39FDDA304
type Node_SetSplineLeashing struct{}

// E49568386456F104
type Node_DrawLine2D struct{}

// 92ED96A77591FD04
type Node_SetWolfSledAutoRunSplineAngle struct{}

// B4C390B0403B4705
type Node_Crank_Enable struct{}

// 0CBDFB84C3719505
type Node_CacheValuesOnStack struct{}

// 73AE4293D8116406
type Node_SetAggressivePriorityOverride struct{}

// 8F07AB3AFF3F6806
type Node_Crank_UpdateDrivenObjectAnim struct{}

// C52D14E2556EA606
type Node_LootPot_Reroll struct{}

// 288B37F314A28C07
type Node_DisableBoatForceTurnAroundControlMode struct{}

// B9CBE07D9FA2BA07
type Node_AddRecipeToWallet struct{}

// 80CE4CB46A3DCB0D
type Node_RemoveTraverseLinkFilter struct{}

// 4929DDCD848B420B
type Node_StopAllSounds struct{}

// 2989409736FB5709
type Node_SimpleStateMachine_ClearStateMac struct{}

// 55FAF3E7AF7AFC09
type Node_SendTelemetryEvent struct{}

// 1C7C9A2235650D0A
type Node_SetSplineFollowDistances struct{}

// A8228CF5B40F5B0A
type Node_ClearFinishedEncounterData struct{}

// A042A1D17C3D910A
type Node_ClearPlayerNavCurve struct{}

// 368617853E18670B
type Node_SetNavObstacleEnabled struct{}

// 122C2762943FEF0B
type Node_SetPreventSoftSave struct{}

// 4CE027774001FF0B
type Node_FindInteractZoneWithName struct{}

// 8FC95FA334416B0D
type Node_SubtractFromGlobalVariable struct{}

// BBF4D2C15527900D
type Node_CheckCineModeIntegrity struct{}

// E718AEFB3E001A11
type Node_InterruptBanterOnCharacter struct{}

// F819B981F8233A0E
type Node_SetContextActionApproachData struct{}

// 4A87EDA43D69F50E
type Node_ForceCompanionTraverseLink struct{}

// F303A75F7ECB150F
type Node_SendImmediateEvent struct{}

// A4EA6BD9694FB30F
type Node_SpawnNPC struct{}

// 268F34910A62FC0F
type Node_SetBehaviorTreeUpdatePriority struct{}

// B26F48EDDBC01A11
type Node_CreateVFSExec struct{}

// A4386EE237F14211
type Node_RemoveRecipeFromWallet struct{}

// 9A97E399977CAB11
type Node_EndCreatureInteract struct{}

// CF0B837AF11DE711
type Node_SendTelemetryEventDLC struct{}

// 9D267AB056B1B318
type Node_SetNavCurveReversal struct{}

// 75862A3D2C034015
type Node_SetJointHighlightCategory struct{}

// 5DFCC8477BED2713
type Node_AlertWave struct{}

// B62F4B594F105212
type Node_DrawArc struct{}

// BBA840868A768D12
type Node_DismissEffectsSubtitle struct{}

// B0C45756AD00F212
type Node_ModifyFocalZoneCameraAngleDeactivationThreshold struct{}

// F769E83282D0F712
type Node_SetChainVisible struct{}

// F45FA65471CC1D13
type Node_AddFlag struct{}

// DBECE7DA9D4F5C13
type Node_DistributeLootResult struct{}

// 5139FD4DB7F0B013
type Node_ClearFocalZones struct{}

// 65F39DF8E8A5FC13
type Node_SetPhysicsEnabled struct{}

// FF5247A36B39C014
type Node_UnmapTouchpadSwipe struct{}

// 1A9B2EC813932A15
type Node_StartYakEnter struct{}

// 6738DA20AFE3C716
type Node_GetClosestElementToPlayer struct{}

// 4ABD4E55056A8015
type Node_TransferWalletContents struct{}

// 569008FA65B19E15
type Node_SetVFSBool struct{}

// 538FCED6126B4816
type Node_Switch struct{}

// 5C715BBAFD107416
type Node_DebugHighlightObject struct{}

// CC66989424F8A216
type Node_SendBehaviorTreeEvent struct{}

// 4CF5F1E0086D5617
type Node_RemoveByValue struct{}

// 48A65FA02D407E17
type Node_SetEnableCounterWind struct{}

// EAEBD054CA33F617
type Node_GrantLootCondition struct{}

// 5E4E9F67AACB1C18
type Node_SetCombatStatus struct{}

// F9230FF35FA39B18
type Node_ClearAllAnimationCallbacks struct{}

// F0BE424E415C451C
type Node_MarkCurrentWeaponModeForPlayer struct{}

// CB1C0A56AB52201A
type Node_SetBehaviorsEnabled struct{}

// 043473E2CF40F118
type Node_ResetListenerPosition struct{}

// 67546100F1C53719
type Node_SetAOOAssignment struct{}

// 9028FA4F28D7DE19
type Node_CreateEquipment struct{}

// 97171F3FCF41021A
type Node_MapTouchpadSwipe struct{}

// E55AE1D8A55B051A
type Node_ActivateWeatherCategory struct{}

// 4D99CD1C4078861A
type Node_TriggerBlenderEffect struct{}

// 884336645708881A
type Node_SetCompassTargetPositionLerpSpeed struct{}

// 9593DC7E02A2001B
type Node_SetUDSActivityAvailable struct{}

// 55CE4F200039191B
type Node_SetJointVisible struct{}

// 7A343491407F3A1B
type Node_Crank_Unlock struct{}

// 2DE0C45B3EE6611E
type Node_Sluice_WaterPourStopped struct{}

// 4952916514C9881C
type Node_SetContextActionMentalStateTag struct{}

// 8500B62F5ECA041D
type Node_SetCombatIndicator struct{}

// E4C12FBA9F8A611D
type Node_ForceCompanionTraversePath struct{}

// EA5893450B92B81D
type Node_SetOverrideSpeedControlData struct{}

// 1E0380F5E66D041E
type Node_ActivateBoatNavCurveFunctionalityFunneling struct{}

// 25DCB8EE9DF86C1E
type Node_ClearVariable struct {
	name string
	next Node
}

// CE6E69C96A68A01E
type Node_ToggleBitOnVariable struct{}

// 4C78F7F902B0F21E
type Node_StartCheckpointedMusic struct{}

// D26E0E945DDF3C1F
type Node_ClearPreventSoftSave struct{}

// 7EDE4757E12E332E
type Node_SetNavAssistCompassMarker struct{}

// 6A2DC71F4AA63E28
type Node_ClearAllDisabledControls struct{}

// F344DF491B074225
type Node_SetMusicCaption struct{}

// 33B54CB0351C8622
type Node_StartTimer struct {
	params [2]Node
	next   Node
}

// E733EC043D008A1F
type Node_SetSplineAvoidDistances struct{}

// A08EAD9CE2E4E31F
type Node_IncrementMeterValue struct{}

// 7DE0ABAD54201420
type Node_OverridePowerLevelForEncounterElement struct{}

// 684437F2DE3E3220
type Node_ClearWeather struct{}

// 735C2E7B7435F221
type Node_StartMusicSting struct{}

// 92410A8B10829B22
type Node_GrantLootConditionSet struct{}

// 2F77F227C0A8B622
type Node_SetInteractZoneLocked struct{}

// 3B35636149B0DE22
type Node_CreatePendulum struct{}

// E2A7F695910FE922
type Node_SetBoatIntoTriggeredLaunchState struct{}

// 55E3F9F4AF766823
type Node_LootPot_SetLootType struct{}

// D11C88B1D544B926
type Node_StopRestartableSoundLoop struct{}

// CF00A48430315825
type Node_SetOverridePositionOnDock struct{}

// 0B01ABD4C7AC5F25
type Node_SetAnimDriverValue struct{}

// F9B23D1E7CA38525
type Node_ClearForceLookAt struct{}

// 6B15C03CC49D0A26
type Node_SetDynamicNavmeshEnabled struct{}

// 7E1BFAB8D15D2A26
type Node_SetWolfSledAutoRunAcceleration struct{}

// 2A2A172FBBA43327
type Node_SetAIUseTraverseLinksWithoutPath struct{}

// 308C019DBE0F4C27
type Node_StartCreatureInteract struct{}

// 2447D8E93F628E27
type Node_SetContextActionFactionTag struct{}

// BA9B425DF41E2628
type Node_AcquireSkill struct{}

// C0D85C334AB42A28
type Node_RemoveCreatureAvailability struct{}

// 9744665A8284182C
type Node_SetTraversePathsEnabled struct{}

// 039F1CF39FA7312A
type Node_RecenterToFacing struct{}

// CC0FB878A20C5B28
type Node_StopTimer struct {
	param Node
}

// EB5418804497DB28
type Node_ResetCompassPathfinding struct{}

// F423BAFE6A03F528
type Node_ActivatePickup struct{}

// E656CE862BC50029
type Node_KillPlayer struct{}

// DFC961A37572E529
type Node_EndPaperboatSync struct{}

// 0D6EBBAC1B3E322A
type Node_WolfSledSplineSectionEnter struct{}

// 4C371668E474AB2A
type Node_SetTraverseLinkOverridePrompt struct{}

// 029B09FE5129F92A
type Node_RevertAggressivePriorityOverride struct{}

// 1C386E87024A052B
type Node_ClearVisualVariantOnDock struct{}

// F5F6FCD63EBD6E2B
type Node_PrepareDynamicSyncedSequence struct{}

// FBED9C71BD25382D
type Node_OverrideCreatureLootDistributor struct{}

// 018577AA8FB94C2C
type Node_Sequence struct {
	nodes []Node
}

// 2EAA9A11A072652C
type Node_SetBreakableEnabled struct{}

// 5C09202FCF23712C
type Node_DrawCircle struct{}

// FF3593492281F42C
type Node_RequestDespawnDynamicCharacter struct{}

// CD34986ACC68252D
type Node_AddCharacterFX struct{}

// 8EE2A416D4FB762D
type Node_ClearOverrideSpeedControlData struct{}

// 9BB4603768D6922D
type Node_SetWaitGateEnabledState struct{}

// 42C0CCCFEB51BF2D
type Node_SimpleStateMachine_SwitchState struct{}

// 658FCCA94838FB2D
type Node_SetContextActionTraversalCompanionLead struct{}

// 6C8A7E2A067B9838
type Node_SetClonePositionRotation struct{}

// 51E44AA5E9634133
type Node_SetPaperBoatOffsetSpringValues struct{}

// 86C3A625D03C7D2F
type Node_EndSyncedSequenceByActor struct{}

// 6E6E6D6D6D1A4A2E
type Node_ProfileBegin struct{}

// 93544193F56E882E
type Node_SwitchOnPinType struct{}

// 72F45AB36D00D82E
type Node_StopAndSettlePendulum struct{}

// 52971ED46F9B5B2F
type Node_SetWolfSledAutoRunStrafeMultiplier struct{}

// B2C4A0B134191B30
type Node_DrawZone struct{}

// AF485758559C5C30
type Node_DeactivateBoatNavCurveFunctionality struct{}

// 2F863749050CCA30
type Node_SetSoundProximityTriggerEnabled struct{}

// 1102C7F05B954632
type Node_DebugSpawner struct{}

// 3E4593F64110FA32
type Node_SetContextActionTraversalSegmentDiffThreshold struct{}

// 156F11A23DD01236
type Node_NotifySluiceIceChanged struct{}

// E93B1DEA0E7BF933
type Node_HideUIMessage struct{}

// D10679CE47A61934
type Node_AddAreaOfOperation struct{}

// 76C490AC76EC8934
type Node_SendContainerRewardAcquiredTelemetry struct{}

// 0A90C8256932AF35
type Node_ShowUIMessageWithTokensBase struct{}

// DDED51C91AB6C735
type Node_SetInfiniteSpawning struct{}

// 0CF8E56F4BC06436
type Node_FixBreakable struct{}

// F13CFC8409357536
type Node_SetMusicWwiseSwitch struct{}

// 717FC3314922FE36
type Node_OverridePowerLevelForEncounter struct{}

// 6FA5583940A59937
type Node_CreateVFSInt struct{}

// 06BAE7F60494E83B
type Node_WildlifeDestroy struct{}

// C29F860E66F9813A
type Node_StartCamera struct{}

// 2A00FD10ED97AB38
type Node_SetPendulumPosition struct{}

// 161B8E733316BA38
type Node_ClearParticleMonsterCheckpointing struct{}

// 9F053B4780A20639
type Node_SetEnableAreaOfOperation struct{}

// 720FFE691266A639
type Node_Sluice_RemoveManualBlocker struct{}

// D582C0B1D425D439
type Node_RegisterSpawnedObjectForFrameUpdate struct{}

// 2EC26095F10CDF3A
type Node_RemovePropToSyncWithCAOwner struct{}

// E05AAABFD680183B
type Node_RegisterForDistanceEvent struct{}

// E9344B811A781C3B
type Node_AddInteractZoneTag struct{}

// 1B08D02274C3A23B
type Node_RemoveIncludeTraverseLinkFilterFromCreature struct{}

// 7AEC7BA79A00CF3B
type Node_SetContextActionInitialized struct{}

// 6855CD0B8F6E323E
type Node_TeleportWolfSled struct{}

// 62EC3086F7EE0B3D
type Node_ModifyNavCurveProgressionLimit struct{}

// 4692F93A886E523D
type Node_SetWolfSledAutoRunSpeedModifier struct{}

// ED30CA2022C5803D
type Node_GrantCreatureLoot struct{}

// 09AD7B015DB9D83D
type Node_QuitAndJumpToWad struct{}

// 3E0CB4D9C9232B3E
type Node_HideAllUIMessages struct{}

// 12DC70E7837B773E
type Node_AddCreatureToCustomSplineGroup struct{}

// 465A37C8D78DB63E
type Node_SetContextActionTraversalDistanceThreshold struct{}

// FB20B46C7EB91E3F
type Node_SetPendulumRotation struct{}

// 16630BEEEC20213F
type Node_ClearOverridePositionOnDock struct{}

// D53108F82C22A75C
type Node_RemoveContextActionCreatureTag struct{}

// 219FE8B841DF424D
type Node_RegisterBoatDock struct{}

// EF58D3E8FD14CE45
type Node_SetWolfSledHarnessVisibility struct{}

// DD6BD2434A0F1742
type Node_CreateVFSBool struct{}

// 301C88564F331841
type Node_DEBUG_ONLY_SetQuestStateAndBackpropagate struct{}

// E31E6CFA256BB13F
type Node_For struct{}

// 943DEDD1D9244240
type Node_ModifyFocalZoneLockInEnabled struct{}

// B584358FD7A14840
type Node_SendEnvironment struct{}

// DCAEF4A1AA828940
type Node_ResetEncounter struct{}

// 3D601DE6260A0F41
type Node_CrankSetup struct{}

// CB73C182D3DD7D41
type Node_SetPosition struct{}

// 635521765585AD41
type Node_Crank_ForcePlayerDetach struct{}

// 5232D8F045C1BE41
type Node_SetCheckpointedMusic struct{}

// B2D67099EB54E341
type Node_AttachZiplineToSpear struct{}

// F205B0231C85F041
type Node_SetSplineAttributes struct{}

// D961E23297EF5244
type Node_PauseEncounter struct{}

// 51611A1B530F2A42
type Node_SetEncounterLODRange struct{}

// 7BEBB046B6249242
type Node_RegisterWaitGateSetup struct{}

// 0A15143B381CD742
type Node_SummonWolfSled struct{}

// 1C8E4548473BE942
type Node_SetPickupStage struct {
	params [4]Node
	next   Node
}

// C4B96DDB9A25B743
type Node_RestoreMarkedWeaponModeForPlayer struct{}

// A095CC228C1F7A44
type Node_SpawnVehicle struct{}

// 1F11942C6A599544
type Node_DestroyBeam struct{}

// 5939F09851630B45
type Node_SetMaterialEntryValue struct{}

// 4B734DF45781A045
type Node_AnimationSync struct{}

// 5D4E56346388B645
type Node_RequestSpawnDynamicCharacter struct{}

// F7EAC2A1C1938949
type Node_SetControlEnabledWithType struct{}

// DFB194EAB199CD47
type Node_SetControlEnabled struct{}

// FE00B48403CEFE45
type Node_PauseCurrentAnimation struct{}

// DBB12B2F19DF3946
type Node_GetTriggerIndex struct{}

// 9340C63E01F6D946
type Node_ClearForcedSpline struct{}

// 31DE27603A865647
type Node_UpdateGameObjectSpline struct{}

// A8117AF7EE3F8347
type Node_SetDynamicCharacterSlot struct{}

// 938CA81FE7CEE447
type Node_ModifyFocalZonePriority struct{}

// 2D6DD310504C2448
type Node_SetInteractZoneHintSubPrompt struct{}

// 587984230F023248
type Node_StopSound struct {
	params [6]Node
}

// D74EA71FFF1C4249
type Node_RaycastCollision struct{}

// C604727C4A8C7049
type Node_ClearOverrideEnterBranchesOnDock struct{}

// C5D0358F1CDA234C
type Node_UnpauseEncounter struct{}

// 5C12FA50FABBF249
type Node_Sluice_StopWater struct{}

// A7ACED2793D0244B
type Node_Crank_Lock struct{}

// 9C4CBC2538FB444B
type Node_EndSyncedSequenceByName struct{}

// F0292FE82CF7814B
type Node_EndCinematicSkip struct{}

// CFE7E5B82F9BDC4B
type Node_StartHapticVibration struct{}

// C53764E25ACD244C
type Node_EnableSetting struct{}

// 4BA5514C4D86484C
type Node_DiscoverSummaryCategoryInAllRegions struct{}

// 314915088211AB4C
type Node_AddEmbedWeaponToGameObject struct{}

// 38179EC010AEE14C
type Node_DisableBoatForceDirectionControlMode struct{}

// D0F5BE15EDFF5756
type Node_ShowUIMessageBase struct{}

// AE3FDD64913A4F51
type Node_ModifyResourceInWallet struct{}

// 5B086678509EC34E
type Node_Crank_API struct{}

// 699B5FEBED41A84D
type Node_DrawSphere struct{}

// 129D7D8D2E42AF4D
type Node_TransitionBetweenFocalZones struct{}

// 15B065324E27CE4D
type Node_MarkUsed struct{}

// D058FED0CF54E84D
type Node_RemoveFocalZone struct{}

// 151970BED78DAA4E
type Node_ModifyFocalZoneMovementPriorityOverrideWhenLockedIn struct{}

// 1D6725E75DE77F50
type Node_SetOverrideVisualVariantOnDock struct{}

// 3856FB635BA39E50
type Node_EnableLoadGate struct{}

// 6B29F2BBAD06FF50
type Node_SetCompassGatewayMarkerIsOpen struct{}

// 28EA6ED49B1B2251
type Node_OverrideCreatureLootConditionSet struct{}

// 9292A0CB50C04151
type Node_ActivateBoatNavCurveFunctionality struct{}

// 8948BE000FA91954
type Node_ResetHeightField struct{}

// 5213DE1DA03AA651
type Node_GetInteractLerpObject struct{}

// EFFA2D420310BA51
type Node_PushPendulum struct{}

// 5995AFF0B0F17252
type Node_UnregisterTraverseGraphEntryEvent struct{}

// 43E658AE7D8DCB53
type Node_SetWolfSledIsDriftAllowed struct{}

// E183E938A952DC53
type Node_RemoveThrowableTarget struct{}

// EA58B431F33A4454
type Node_ApplyStatusMeterDamage struct{}

// 5669F3B305DF4454
type Node_FreezePendulum struct{}

// 7B2C555E4DA90D55
type Node_DiscoverSingleRegionSummaryCategory struct{}

// 07CC6B26748F8E55
type Node_SetRouteToOtherRealmsViaFastTravel struct{}

// 563C3B673122C558
type Node_CreateBeam struct{}

// 397DB182DDC6AF57
type Node_RemoveAreaOfOperation struct{}

// A85EE96210E98156
type Node_BreakBreakable struct{}

// D1EC526ADCB5A156
type Node_PlayBanter struct{}

// 03E167B884B5BC56
type Node_ModifyFocalZoneEnablePitchOnHeadTracking struct{}

// D2F9C5FED111E556
type Node_UnregisterMinibossCheckpoint struct{}

// C389F20F91B02757
type Node_TriggerPendulumSlowdown struct{}

// AB5D4ADF86D41A58
type Node_StartNewDLCTelemetryRun struct{}

// 4701198492FF3B58
type Node_SetMeterValue struct{}

// 31D075C18EA34058
type Node_SetManualSaveEnabled struct{}

// 3DD30475EA985E58
type Node_MarkSkipAIUpdate struct{}

// 04ECF96E34EF7D58
type Node_SetTraverseLinkPromptOffset struct{}

// B576B7E88896745A
type Node_AddCombatCollision struct{}

// 386125DA5A6F1759
type Node_EquipToEquipmentSlot struct{}

// AEA860A240B9AD59
type Node_SetupPaperboatSync struct{}

// 096D3084EE48D059
type Node_UsePlayerHeightForFloating struct{}

// 69610626A2D8D259
type Node_SendNamedEventBase struct{}

// 71E735FBCB73EB59
type Node_CreateEmptyLootResult struct{}

// F5B6C5F0DA93025B
type Node_RegisterForBreakableBrokenEvent struct{}

// E0A1F327FF70095C
type Node_ActivateContextAction struct{}

// BE22F8C8B47E905C
type Node_SetScale struct{}

// BED58696F973945C
type Node_Crank_JamCrank struct{}

// 314C06281E34876D
type Node_SetPreventCinematicSkip struct{}

// 43AD68C88FB8B663
type Node_SetContextActionMultiCreatureData struct{}

// A8CB958D0EECBF5F
type Node_SetGlobalVariable struct{}

// 3FBE712E3A79DB5D
type Node_DestroyVFSEntry struct{}

// 73DA311E272BD25C
type Node_SetTraversePathPromptEnabled struct{}

// 8A1284A602512F5D
type Node_OverrideCreatureCullingFadeDistance struct{}

// 229A3402AAFA765D
type Node_SetFightKnowledgeEnabled struct{}

// 6898571C88FCBC5D
type Node_WolfSledForceEnter struct{}

// 642D85E5B06FDA5D
type Node_ClearWolfSledSpeedControlOverride struct{}

// 68F3CC91971D845E
type Node_SetWolfSledRopeVisibility struct{}

// 4B444193D8F1D35E
type Node_SendNamedEvent struct{}

// B335C58CB151125F
type Node_PositionObjectsForSyncedSequence struct{}

// A4C1A9659B9C285F
type Node_UnequipFromEquipmentSlot struct{}

// 4416805AC540515F
type Node_DespawnVehicle struct{}

// 9A667F95ADAC2B62
type Node_SimpleStateMachine_Clea struct{}

// F1D8FEFFE1926560
type Node_RegisterTraverseGraphEntryEvent struct{}

// 4226E02CC656E960
type Node_RemoveChild struct{}

// 05FC42C1A0EB7261
type Node_DrawDebugTable struct{}

// B8240C85F4D68861
type Node_Crank_Disable struct{}

// 11518B3408BF7862
type Node_ActivateWeather struct{}

// 5EFA7E8894FF7E62
type Node_AddGameObjectToProfileGroup struct{}

// 730E89A86343D062
type Node_SetZiplineSlack struct{}

// AF807E205C2F1963
type Node_TriggerQuestManualActiveBehaviors struct{}

// A7D0B75CCEDF8A63
type Node_SetPaperBoatFadeIn struct{}

// 3EE00202E4E1BE67
type Node_Print struct {
	param    Node
	logLevel Node
	next     Node
}

// 61DBF01B41B1B665
type Node_SetWolfSledAutoRun struct{}

// E82FBA575A5EC363
type Node_ResetAccesibilityToggleState struct{}

// 65374919F2D2E563
type Node_SpawnVehicleAtPosition struct{}

// 2C9D8FDCEF568164
type Node_OverrideCreatureMotionParams struct{}

// ED1919F454FEED64
type Node_SetMarkerLocked struct{}

// 59DA6CA9BDF5A365
type Node_PrestreamRequest struct{}

// 20EE6BA57EFB1C66
type Node_TriggerEncounterEnd struct{}

// D137C56A2AAA4667
type Node_BreakLoop struct{}

// C7C4F34846BA6367
type Node_AddSpawnToEncounter struct{}

// 0773A21B30F76D67
type Node_SetSplineSpeed struct{}

// DF05DE4B98AFB167
type Node_AddPointToFX struct{}

// E06414F529DCCC6B
type Node_ClearCallAndResponseObject struct{}

// 061CEFE38BEF3069
type Node_OpenSoundPortal struct{}

// F4D93323EBEFCF69
type Node_SetPlayerControlledCreature struct{}

// 2BE57FD32FABAF6A
type Node_SetWwiseRTPCValue struct{}

// 2D255CBD71900F6B
type Node_GetBranchStartPositionAndOrientation struct{}

// BB23CE23A3C67E6B
type Node_SetWolfSledAutoRunSlopeBoostMultiplier struct{}

// A7ADEFCC5C47286C
type Node_SetSoundPortalOpen struct{}

// B1C43C68ABD6306C
type Node_StartPaperBoatBounceFadeOut struct{}

// 2C5404F8CB91096D
type Node_SetContextActionEncounterMarkerIDTag struct{}

// D79BB5A9C1F9696D
type Node_DrawRect2D struct{}

// 66CA00E9A7A90173
type Node_SetInfluenceConeAttributes struct{}

// B073BF6B158CBE70
type Node_SetPlayerNavCurve struct{}

// 36BC76908E0A8F6E
type Node_DestroyEquipment struct{}

// 9F92B70E4931D86D
type Node_SetFlawlessZeusAvailable struct{}

// E6D6F3771F35EC6D
type Node_OverrideLookAtConversationSettings struct{}

// 3A6A03ADAE14F66D
type Node_SetTextOnTextObject struct{}

// D3865F197FEF626E
type Node_SetYakInteractOptions struct{}

// 3BC222CC06E9696E
type Node_GameMap_SetItemStateV2 struct{}

// 6E464014AC1CED6E
type Node_TriggerCallAndResponse struct{}

// 59BF1AB246EB3D6F
type Node_CheckFeature struct{}

// 434314B7011C586F
type Node_CoilChainedObject struct{}

// 420FEB9D3DBFBB6F
type Node_ModifyFocalZoneFacingAngleDeactivationThreshold struct{}

// 1B304DDD6C3D9A70
type Node_StartWolfSledExit struct{}

// 4DDA11FE18701672
type Node_SetSlowdown struct{}

// D821C10DA2C13B71
type Node_InterruptAllBanter struct{}

// 8B1762741EA69B71
type Node_Sluice_AddManualBlocker struct{}

// 775021CFA11CD471
type Node_StopCamera struct{}

// 130035D4E3EAE271
type Node_SetWolfSledAutoRunRequiresInput struct{}

// F056416533200E72
type Node_RegisterForControlMashEvent struct{}

// B30F62EEE87BA072
type Node_UnpauseCurrentAnimation struct{}

// D9DC1F71DE85A472
type Node_SetCreatureAvailabilityRequest struct{}

// 68FDE7C168CDB472
type Node_SetCreatureFocus struct{}

// 5087AC8E5960BE72
type Node_ShowHiddenWads struct{}

// 5E189DCCC2890B77
type Node_OverrideCreatureLootCategory struct{}

// A0469122B8A38F74
type Node_ClearAllWeaponsCinematicMode struct{}

// D586105632799173
type Node_ModifyFocalZoneAttachObject struct{}

// C3FDBEC4E5C09173
type Node_SetBitOnGlobalVariable struct{}

// AC4762B8E9EC0674
type Node_ClearAndHideAllNavAssistCompassMarkers struct{}

// D7E9B8996A9A5D74
type Node_ClearCreatureFocus struct{}

// E7D6946E7F566274
type Node_RemoveSecondaryActorFromSequence struct{}

// CB414DE50803DD74
type Node_RegisterSpawnedObjectForVariables struct{}

// A4333446FD611075
type Node_PrepareSyncedSequence struct{}

// 7D838CB70CA44875
type Node_SetBoatDockInteractRadius struct{}

// 52B3BAA2A5300376
type Node_RegisterAsChainedObject struct{}

// 11A796139E109076
type Node_EnableContextAction struct{}

// 48A8AF7C45DBD079
type Node_SimpleStateMachine_SwitchStateUsingEnumPin struct{}

// 740024F1E7B57377
type Node_ForEach struct {
	param     Node
	params2   []Node
	operation Node
	next      Node
}

// 0B0F251FC92A7F79
type Node_SetBanterFact struct{}

// 6D9F38641D279779
type Node_SetWeaponMode struct{}

// 346DC94BAFA2BD79
type Node_SetWolfSledAutoRunControlRange struct{}

// D5B3A3CAF4ABC179
type Node_TransferLootResultEntry struct{}

// 778C5529632BD679
type Node_SwitchIntRange struct{}

// 1E5A282AF598267A
type Node_DropThrowable struct{}

// DFFA41428F01EB7A
type Node_RegisterForMeterEvent struct{}

// 333E378B9262FE7A
type Node_SoundEmitterEnabled struct{}

// D91BE5EC6BE840BB
type Node_SetFlawlessAresAvailable struct{}

// 6E7430D3CF19CC9B
type Node_SetContextActionBasicData struct{}

// 84C244632D21FF8C
type Node_FakeInteraction struct{}

// 936167F32AA52D84
type Node_SetSpawnedObjectVariable struct{}

// E6D818053070497E
type Node_StartBoatSyncMove struct{}

// 299BC62F52B8867C
type Node_SpawnWildlife struct{}

// C401E5F40ACF787B
type Node_SetPointOnFX struct{}

// 8811A6E37059027C
type Node_GetNextArrayValue struct{}

// 0636C4F6B6F80D7C
type Node_ProcessInteractByCandidateSet struct{}

// A0A7D75092F2807C
type Node_RegisterForBlackboardTimerExpired struct{}

// 4AE5F0FD2520867C
type Node_ClearHideHealthBar struct{}

// 4D3F6C2D2363937C
type Node_SetTraversePathPromptLOSEnabled struct{}

// C0078F22522ACD7C
type Node_TriggerCamera struct{}

// D7DEAFD6478AD17C
type Node_CacheCallAndResponseObject struct{}

// F82A8C3D8E593D7D
type Node_SetPropEnabled struct{}

// F17FE57DE7713F7D
type Node_ResetStickDamping struct{}

// E32E2634F0F9D081
type Node_SetResourceDebugAddMode struct{}

// 0B35998ECD02947F
type Node_SetUDSActivityStart struct{}

// 7716D3FF11C6FC7F
type Node_AddNewLookAtTarget struct{}

// 1C8ED930A00CC280
type Node_RegisterAsZipline struct{}

// 02F42A0926C06D81
type Node_SetAccelerationOverride struct{}

// 5A0CF0D59BD7C081
type Node_CustomSortIterator struct{}

// 98A120AC62923982
type Node_SetSplineAvoidAvoidancePreferences struct{}

// A75D59680BEE4782
type Node_UnregisterForMonitorEvent struct{}

// 16206A5B4EA99482
type Node_RegisterForMovePlayEvent struct{}

// 0BC4A5B61373B982
type Node_AddContextActionCreatureTag struct{}

// 2D54924FDA4B1D84
type Node_ClearPreventCheckpoint struct{}

// 5D9BB9E76E013A89
type Node_SetValhallaComplete struct{}

// D7F610BFD75EDE87
type Node_SimpleCombatSync struct{}

// E56A267342096F84
type Node_SetBlackboardTimer struct {
	params [4]Node
	param  uint16
	next   Node
}

// 6674CE0272981885
type Node_ParticleMonsterCheckpointTotemState struct{}

// A0DC62E57BBDA485
type Node_WildlifeFlee struct{}

// 64D5602381F8D485
type Node_OverrideListenerPosition struct{}

// 1CB11C7D4FD6A986
type Node_SetRaceBackHomeEnableSledControl struct{}

// EB6F214E9473E187
type Node_SetForceOddNumTraverseLinkConnections struct{}

// E128B61ECBAE0788
type Node_StartYakExit struct{}

// 2F8EDA94CF125588
type Node_ForceEngageContextAction struct{}

// B0C8F5C04B67C388
type Node_SetCombatTarget struct{}

// 640F100B5C9F1A89
type Node_SetPaperBoatSpeed struct{}

// 668E069FF8A0048B
type Node_RegisterTraverseGraphEnteredEvent struct{}

// 865F09779951F889
type Node_ClearBoatDockInteractRadius struct{}

// B4F8CFBF24D9778A
type Node_SetVisibility struct{}

// C4A39C82A587808A
type Node_GetAdditionalGravityOnPendulum struct{}

// 391742579FF1B98A
type Node_SetObjectHighlightCategory struct{}

// D2BBE09E1C5B028B
type Node_OverrideCreatureCullingRadius struct{}

// 03D28F3C9639468B
type Node_SendUIEvent struct{}

// FC7C9BF83551D38B
type Node_ToggleBitOnGlobalVariable struct{}

// E6C4FED4DD85FA8B
type Node_RemovePositionalFlag struct{}

// 38C9CDDE5C664B8C
type Node_UntrackAllSideQuests struct{}

// 704E59DE641A4794
type Node_ClearAllDisabledControlsWithType struct{}

// 75FAC3DCC24AF38F
type Node_AddThrowableTarget struct{}

// C8A4564AF1AA628E
type Node_SetInteractZoneEnabled struct{}

// ED0F9F98867B0C8D
type Node_PrepareRegisteredSyncedSequence struct{}

// E97CEFEBEECA398D
type Node_SetBitOnVariable struct{}

// FC108195473A878D
type Node_OverridePowerLevelForEncounterWave struct{}

// 2212589FBE33C98D
type Node_SetTraversePathDirectionEnabled struct{}

// A8AD5D7DF30D398E
type Node_SetDockEnabled struct{}

// A2450476F9099F8E
type Node_SetCharacterFX struct{}

// 2D2A949C7ECD298F
type Node_PlayMaterialAnim struct{}

// 4989D6ABC788378F
type Node_SetSubtitleDistanceThresholds struct{}

// E929EAA15401A18F
type Node_WolfSledSplineSectionExit struct{}

// 4EB946151140A68F
type Node_SetFlightVolumeEnabledByType struct{}

// 46351222B9005E92
type Node_CreateEncounterData struct{}

// 48AE1D67FFD63F90
type Node_SetSubtitleDistanceThresholdsAlternate struct{}

// BD2977F2C15C5790
type Node_ClearEquipmentPreview struct{}

// 7B44093CB0869F90
type Node_ClearMaxSpeedOverride struct{}

// 6AA211B734E6DE90
type Node_SetCreatureHighlightCategoryOverride struct{}

// 1B637553DB144792
type Node_ForceZeroTime struct{}

// 932CDBB618E06A92
type Node_TriggerQuestManualCompleteBehaviors struct{}

// 58B8F950D55D7392
type Node_DrawCone struct{}

// 24B5FD371D619692
type Node_Warp struct{}

// 45980D6E32BA0093
type Node_ForceLookAt struct{}

// 49978BF425648593
type Node_SetElement struct{}

// 34EAA3B4230BE198
type Node_EnableStopToAvoid struct{}

// 38186362C8A14A95
type Node_EmitArrowWithEmitter struct{}

// B979CEBE13C2BD94
type Node_IncrementBanterFact struct{}

// 5D16F611413BEC94
type Node_ExecuteVFS struct{}

// 108FD597EDB2FD94
type Node_DespawnDynamicCharacter struct{}

// ED858B153B2F1B95
type Node_EnableStoryTime struct{}

// 699A313A4A7F3795
type Node_Crank_SetComplete struct{}

// A70317F760F04E95
type Node_EvaluateLoadZones struct{}

// E54443A46FED9B95
type Node_TriggerMoveEvent struct{}

// 8E7AFA908FD5E896
type Node_DrawTextInWorld struct{}

// FCC8A7AE429DF597
type Node_ModifyFocalZoneEnabled struct{}

// 460CE4F00675AC98
type Node_RemoveExcludeTraverseLinkFilterToCreature struct{}

// 0E74F748793AE199
type Node_Pop struct{}

// A8BD4E51ADF01499
type Node_LockPlayerOnWolfsled struct{}

// 959BA82796F32999
type Node_DrawText2D struct{}

// B213B3202D415399
type Node_SetPreventCheckpoint struct{}

// FDB50F7AA4E88599
type Node_Crank_FreeCrank struct{}

// F380AF556AE79A99
type Node_SetInteractZonesEnabled struct{}

// 30B3C9BD624CFC99
type Node_RemoveCombatCollision struct{}

// 7FD48C222279B99A
type Node_SetContextActionRangeData struct{}

// F33CF886B66AD69A
type Node_RegisterForLookAtEvent struct{}

// E54754C1C7A55C9B
type Node_SetWolfSledInteractOptions struct{}

// 3CF06B296D68F9AC
type Node_SetInteractZoneTemplate struct{}

// B58B12285ABC5DA3
type Node_PlaySound struct {
	params [9]Node
}

// 06ED874C600BFA9F
type Node_ThawPendulum struct{}

// 47C289B71BA5ED9D
type Node_DestroyAllCombatConcussions struct{}

// 45A6AA43C708039C
type Node_SetSoundEmitterSplineCap struct{}

// ED086487612CB69C
type Node_ShowCompassMarker struct{}

// 2BD3C99FC77A029D
type Node_SetPendulumRestrictionPlane struct{}

// 0B95A4D7AAD0089D
type Node_CreateVFSFloat struct{}

// 0BADEF43F0BF459D
type Node_ModifyFocalZoneObstacle struct{}

// 6E6D631004FB2A9E
type Node_SetAdditionalGravityOnPendulum struct{}

// FC8B572B55812F9E
type Node_PlayUISound struct{}

// 91E07862FAB91C9F
type Node_SetGIBlend struct{}

// EE595C442C98349F
type Node_SetSplineEnabled struct{}

// FE58556DE495709F
type Node_ClearWeaponCinematicMode struct{}

// BDB26859BD0CA7A1
type Node_SluiceSetup struct{}

// E06930B6A95E2BA0
type Node_SetSingleContextActionEnabled struct{}

// 154E1258828695A0
type Node_SetPlayerNavCurveAutorun struct{}

// 71CD4766342DA7A0
type Node_SetRotation struct{}

// AEE9054F63B253A1
type Node_RegisterOnEncounterStart struct{}

// F21D95B0478585A1
type Node_CreatePendulumWithAnchorJoint struct{}

// D1CACD0FDD7DC5A1
type Node_SetWolfSledSpeedControlOverride struct{}

// AEA7CBEDA61807A2
type Node_AddIncludeTraverseLinkFilterToCreature struct{}

// E23288D20C8CE3A2
type Node_AddPropToSyncWithCAOwner struct{}

// 814A6A0F5FF9EBA2
type Node_RegisterForControlPressEvent struct{}

// C589C70C1A5009A3
type Node_SetSplineSpeedType struct{}

// E907ED7BBD9878A7
type Node_PlayAnimation struct{}

// F99907F6D50948A5
type Node_SetCreatureExitedWaitGate struct{}

// E9B8521F3B7B01A4
type Node_RemoveEmbedWeaponFromGameObject struct{}

// 570A24C77A612DA4
type Node_ChangeBehaviorTree struct{}

// B8915EA08DBD2FA4
type Node_StopMusic struct{}

// 6BD3A840ABC9F3A4
type Node_SetWwiseState struct{}

// 8B89BD2B78DD29A5
type Node_ActivateBoatNavCurveFunctionalityAutoMovement struct{}

// 9BD48B594AFA54A5
type Node_SendSystemicBanterResponse struct{}

// 020806227BCF7CA5
type Node_PlayRestartableSoundLoop struct{}

// 7313E9BD6CF27FA5
type Node_SetTraversePathCustomEntry struct{}

// 0C4F78EA31F1F6A5
type Node_RecenterToFramePosition struct{}

// D7756AF46F618BA6
type Node_SetZiplineSpeedOverride struct{}

// 8F98E6CF544B4BAA
type Node_DetermineSplineSpeedAndInputLocks struct{}

// B8920460E2B483A7
type Node_SetVFSInt struct{}

// C1A2F2BC5950C3A7
type Node_CopyUpgradeEquipState struct{}

// 9E7338CD4D3227A8
type Node_SetMaxSpeedOverride struct{}

// CA4B548CAC3847A8
type Node_AddFocalZone struct{}

// F3FF6D257BEDB4A8
type Node_SetRealmLocked struct{}

// 2B63F1F9A07591AA
type Node_AddPositionalFlag struct{}

// 4B003F6849C0B2AB
type Node_DisableContextAction struct{}

// C003C543D34D07AC
type Node_ProfileEnd struct{}

// 0D2B060CAB7C40AC
type Node_AddToGlobalVariable struct{}

// 7C73674E0DED74B3
type Node_SetCameraTargetEnabled struct{}

// E6890CA7D51AD5B1
type Node_SendLuaHook struct{}

// A42E1444D1908DAF
type Node_GetArrayElement struct{}

// 89BE5980598A09AD
type Node_RemoveCreatureFromCustomSplineGroup struct{}

// DB0944412F63BFAD
type Node_TransferFullLootResult struct{}

// 27ED992A8A91E0AD
type Node_OverrideStickDamping struct{}

// 89CFAA38D5DF19AE
type Node_SetHighlightEnabledOverride struct{}

// CEE4429F322AA9AE
type Node_SetMaxSpeedOverrideWithType struct{}

// BBB3B7601DCDE1AF
type Node_RegisterOnEncounterReset struct{}

// 6F714798792DECAF
type Node_RegisterForDeathEvent struct{}

// 3DCE422724D3CFB0
type Node_RestoreDefaultConversationSettings struct{}

// C1800676C4A84DB1
type Node_DespawnEncounter struct{}

// C77FE43B811988B1
type Node_SetIsDeathAllowed struct{}

// 5536CDAD8E4308B3
type Node_ClearMaxSpeedOverrideWithType struct{}

// E82753E110B6D7B1
type Node_SkipWave struct{}

// A5BC7B841A46EAB1
type Node_ClearAllMarkUsed struct{}

// 93F317BBFFBA20B2
type Node_EnableBoatForceDirectionControlMode struct{}

// 73FD14DB205D77B2
type Node_EnableAllInteractTags struct{}

// BEBB357D4BE2D4B2
type Node_RegisterForHealthChangeEvent struct{}

// 586FBE26D77818B3
type Node_AcquirePickup struct {
	params  [4]Node
	params2 []Node
}

// 703FB7ADC34639B3
type Node_StoreCheckpointWithPositionOverride struct{}

// 19B9CEED4CF96AB3
type Node_HideCompassMarker struct{}

// 17E2EF454E236FB3
type Node_CheckTriggerWithI struct{}

// 9B53B26CDC5F70B7
type Node_BucketEquipmentByTrait struct{}

// CF0F8A1ABC0E2BB5
type Node_CloseSoundPortal struct{}

// 1428B5B813E68EB3
type Node_HideWads struct{}

// 8354A048658DCCB3
type Node_ApplyEquipmentJointsToObject struct{}

// A47E97F13A6278B4
type Node_HideHealthBar struct{}

// EEC20F7878088DB4
type Node_Remove struct{}

// 70D71059AA5CE3B4
type Node_SetTraverseLinkEnabled struct{}

// F6DCADF3E65E82B5
type Node_RemoveFlag struct{}

// 72D321C22620FCB5
type Node_ModifyFocalZone struct{}

// 800379829F5402B6
type Node_FunctionCall struct {
	params  []Node
	params2 []Node
	node    Node
	next    Node
}

// E620D9C5E701D9B6
type Node_SuspendCreatureCulling struct{}

// D42C7509A63A00B7
type Node_SetContextActionCustomData struct{}

// E7DBC0432EA296B9
type Node_Spawn struct{}

// 9C774F258C4097B7
type Node_SetInteractZoneMainPrompts struct{}

// 7F719B5AFF7C01B8
type Node_FilterIterator struct{}

// 9EFDE2F647C016B8
type Node_AddLookAtPriorityOverride struct{}

// FB9211B734B116B9
type Node_SetInteractZoneSubPrompt struct{}

// 50A0D02A440589B9
type Node_DeactivatePickup struct{}

// C09EEC09B6972BBA
type Node_TriggerPlayFX struct{}

// FE5ACCE5749F5ABA
type Node_SetTeam struct{}

// 1F06E7DA5EDC8CBA
type Node_ClonePose struct{}

// A32FF958880623BB
type Node_ClearSoundEmitterSplineCap struct{}

// 33DE5CD422F239E0
type Node_RemoveCombatIndicator struct{}

// 160858D1E26567CB
type Node_ConfigureLogicLoadGroup struct{}

// BD6B9D367BBBD0C3
type Node_RegisterForEnemySpawn struct{}

// 1CA6830FB3AF14C0
type Node_AssignCreatureToEncounter struct{}

// 6CB2FB587A8EF4BD
type Node_SetSkillTreeToken struct{}

// ED799B7E6D9717BC
type Node_FinishHapticVibration struct{}

// F1D367B310F491BC
type Node_SetParticlesEnabled struct{}

// 3A103AA04E69B1BC
type Node_RegisterForAnimFrameEvent struct{}

// 2AF1E3155106B5BC
type Node_SetThrowable struct{}

// 99FBF37F4C5E91BD
type Node_SetBlackboardVariable struct {
	params [4]Node
	param  uint16
	next   Node
}

// 3EA9812733C284BE
type Node_If struct {
	Condition Node
	True      Node
	False     Node
}

// CEA196DA307787BE
type Node_ResetPadLightColor struct{}

// 7D8CF66F8FA13BBF
type Node_SetWindMotorsEnabled struct{}

// 7492B9D10A2996BF
type Node_MeterAdjustPauseDelay struct{}

// C577213B3DAB0AC0
type Node_ZoomSnap struct{}

// 5686F27FBD621BC2
type Node_RelinquishPickup struct{}

// 93456DA09F6173C1
type Node_SetWolfSledAutoRunSplineSpeed struct{}

// 23CF75E3714C80C1
type Node_SetPaperBoatDistanceFromSplineClamp struct{}

// 88444EF88BAE90C1
type Node_RemoveLookAtTarget struct{}

// 4C9F1A5A7BADADC1
type Node_DrawCircle2D struct{}

// D260E07CAF060EC2
type Node_SetPendulumRestrictionAngle struct{}

// 9A2E9FED1B762DC2
type Node_RegisterForFrameUpdateEvent struct{}

// F7578140C5A266C2
type Node_ClearGlobalVariable struct{}

// 0CBDEEFA3218BAC2
type Node_RemoveLookAtPriorityOverride struct{}

// 104CBA15132F22C3
type Node_SetZiplineStartAndEnd struct{}

// 75FFBF388E3C2AC3
type Node_MusicFadeOutAndLogVolume struct{}

// 7B790CFD067064C8
type Node_StartMainMusicTrack struct{}

// D86A7DE7A38B3BC7
type Node_SetPadLightColor struct{}

// 17030A7A0E6462C4
type Node_EmitArrow struct{}

// 006615B249EA39C5
type Node_RegisterForEventQueueProcessedEvent struct{}

// 0E669863732A85C6
type Node_SetProgressionFact struct{}

// 89C2B1718383CFC6
type Node_ResetTraverseLinkFilterForCreature struct{}

// 11EEF6D23D45F7C6
type Node_CheckPersistentCreaturesEnabled struct{}

// 7DAC2138270FB2C7
type Node_DisableLoadGate struct{}

// DAE24153E3C2C1C7
type Node_RemoveFX struct{}

// C17719B48B2708C8
type Node_RemoveArrows struct{}

// B62D431D723214C8
type Node_ResetTraverseLinkFilter struct{}

// 0766B9A6D53C2BC8
type Node_FreeRequestedPrestreams struct{}

// 2715F8E04AB12ECA
type Node_RecenterPendulum struct{}

// CD8402C769416CC8
type Node_SetWolfSledAutoRunEnableSteering struct{}

// 5C3EE4310A07ACC8
type Node_SetSingleNavCurveEnabled struct{}

// 078AA15F42E331C9
type Node_SetWeaponCinematicMode struct{}

// A4594F86CDEB6BC9
type Node_SendContextActionStim struct{}

// 13EC9DD37EBF17CA
type Node_AddEquipmentPreview struct{}

// F4552D883BDD64CA
type Node_EquipToCharacterSlot struct{}

// D28B3FC779BF1DCB
type Node_StoreSoftSave struct{}

// 6890F438920F49CB
type Node_ForceInterruptContextAction struct{}

// 590A375878C159CB
type Node_TriggerEncounter struct{}

// 8F34F197FA833AD3
type Node_SetCreaturePersistent struct{}

// 58DC3020AFE4CBCF
type Node_SetSoundGeometryEnabled struct{}

// 911500A33D3E1BCD
type Node_RegisterForButtonPressEvent struct{}

// 1181913D2D39E5CB
type Node_SetHapticGlobalParameter struct{}

// 46493595ED2900CC
type Node_ForceCompanionStriderPath struct{}

// E8034B50CDC8AFCC
type Node_Push struct{}

// 3E409BF05CDE14CD
type Node_Sluice_WaterPourHit struct{}

// 4D5CD6B23FEB43CD
type Node_GetParticleMonsterTotemCheckpointState struct{}

// F5A02EDEBA664ACD
type Node_AddTraverseLinkFilter struct{}

// 0859ED9B8EA393CD
type Node_SaveYakWaterSpeedType struct{}

// 09814617B42945CE
type Node_UpdateGameObjectSplineCheap struct{}

// AAA4902BA47498CE
type Node_SetInfluenceCircleAttributes struct{}

// 829CE0E9FB2E8FD1
type Node_WildlifeEvent struct{}

// 62B1E65A1D890AD0
type Node_RegisterOnEncounterCreated struct{}

// E04618BA38E132D0
type Node_SpawnDynamicCharacter struct{}

// DDBFD135EE960ED1
type Node_SendAwarenessStimWithPa struct{}

// 306FC73442ED64D1
type Node_SetVariable struct {
	name string
	val  Node
	next Node
}

// FB7D4D63C03F86D1
type Node_SetTraversePathSegmentEnabled struct{}

// 02EA73472B150DD2
type Node_SetCreatureEnteredWaitGate struct{}

// 4A798C01AFB922D2
type Node_ClearCreatureMotionParamsOverride struct{}

// 52AE2B1E2F5384D2
type Node_RegisterForReticleTargetInvalid struct{}

// 9DD616B30C7DC0D2
type Node_SetCompanionFollowStriderPath struct{}

// 2D156651A99FF8D2
type Node_SetCompassEnabled struct{}

// 3D87A333EBE35DDA
type Node_AddChild struct{}

// EBB7567C0E0B03D6
type Node_ForceUseSpline struct{}

// C60095F026B241D3
type Node_SetModelEnabled struct{}

// C8F0D35E1DA069D3
type Node_SendAwareness struct{}

// 8AF0389F805AA5D4
type Node_SetVisualScriptEnabled struct{}

// DA0A5BCFA82514D5
type Node_UnregisterTraverseGraphEnteredEvent struct{}

// EE1AD52FED2AFAD5
type Node_CompleteWave struct{}

// 94F21B6443784FD6
type Node_SetWolfSledAutoRunResetSettings struct{}

// F5353AA0246B0ED7
type Node_SetVFSFloat struct{}

// 9C7A8F0CBE5E38D7
type Node_SetAnimFloatChannelDriverValue struct{}

// D716A9489B8ABBD8
type Node_UpdateStaticWindClientPositions struct{}

// C08A6EF4FE48C3D8
type Node_RotateObjectTowardsTarget struct{}

// F56486D47F7AC7DD
type Node_SetPaperBoatSpringValues struct{}

// 4BC17693B7199CDB
type Node_CancelLockOn struct{}

// BEEA97399149A3DC
type Node_StartMusic struct{}

// 06D2FB09AC380DDD
type Node_RegisterForUIButtonPressEvent struct{}

// 0790A5250A294DDD
type Node_StopBlenderEffect struct{}

// 83F841C3053D7EDD
type Node_SetForceCompanionTraverseLink struct{}

// 46BF5D0FA7360FDE
type Node_SetDecalsEnabled struct{}

// 387B9CAAB2DE38DE
type Node_ClearRage struct{}

// 4727B2C4817C3ADE
type Node_GetInteractObject struct{}

// 3DB70B15D52A48DF
type Node_SetSavedNavCurveAttributes struct{}

// DADE3DCF0F2D8BF2
type Node_SetSplineLeadDistances struct{}

// 514178653F5CD5E9
type Node_ModifyFocalZoneObstacleEnabled struct{}

// 545D9E1D95BD28E5
type Node_SpawnFX struct{}

// 4C3EFE4DEB0BCBE1
type Node_SpherecastCollision struct{}

// 391AF8B8AF3243E0
type Node_SimulateSpearZipline struct{}

// E1B98A6DC7E479E0
type Node_SetAllWeaponsCinematicMode struct{}

// 4115DE9E537B7BE0
type Node_IncrementProgressionFact struct{}

// 81ED7185C3790CE1
type Node_RegisterOnEnemyDeath struct{}

// 59614AE27D8F2DE1
type Node_UnlockCombatStatus struct{}

// E552C038B5FED7E1
type Node_CompleteWaveElement struct{}

// 2719CF7D476FFDE2
type Node_SetTraverseLinksEnabled struct{}

// 666D38C8B83528E3
type Node_SetCharacterConfig struct{}

// 58E222C4EC3CFCE4
type Node_EnableBoatForceTurnAroundControlMode struct{}

// 1E8B6A149CB61EE5
type Node_FrameDelay struct{}

// DCAF72DFAA47ECE7
type Node_SetInteractTagEnabled struct{}

// ABF8484B8E2CE9E5
type Node_DestroyPendulum struct{}

// 4AD10217ADD409E7
type Node_SetZiplineEndDisconnected struct{}

// 431B8AAFE6711BE7
type Node_DestroyGameObject struct{}

// F62E6CE42D9073E7
type Node_SetVolumeEnabled struct{}

// 8E50C18E35B79FE7
type Node_ClearAllInteractZoneTags struct{}

// 8ED4956C99540FE8
type Node_AddSimpleObjectActorToSequence struct{}

// 4DC272A52D3312E9
type Node_PushWeaponStateForSequence struct{}

// 95F55CD91AB52CE9
type Node_ClearCustomSplineGroup struct{}

// 08ED25FC75F798E9
type Node_AddEquipmentToWallet struct{}

// 00959812C143BCE9
type Node_CheckDecision struct{}

// 14D68E2697B0F5ED
type Node_SetContextActionLoopActionData struct{}

// 18FBC2D5EC1A86EA
type Node_SetOverrideEnterBranchesOnDock struct{}

// 3DF048F7A817DAE9
type Node_SetOverrideStatusEffectIconParent struct{}

// EA40D91F29E3ECE9
type Node_SetWwiseSwitch struct{}

// 384DCA43F6210EEA
type Node_Error struct{}

// 81F87ACD9E2570EA
type Node_RequestMaterialSwap struct{}

// 6E74FF9A5F9F7DEA
type Node_SetContextActionTraversalBehavior struct{}

// 699CD6CF59262AEB
type Node_BatchQueryEquipmentTraits struct{}

// 6C34C065D76515EC
type Node_ApplyContextIdle struct{}

// 44141A7DF026CBEC
type Node_SetContextActionBasicBtreeData struct{}

// 14A7C3903DF5D0EC
type Node_SubtractFromVariable struct{}

// 828474D4FC561AED
type Node_SetCompanion struct{}

// 552AEA13CB7C14F0
type Node_MusicFadeIn struct{}

// F8DB059D6EF36EEE
type Node_SpawnGameObject struct{}

// 780702AA5482B9EE
type Node_SendPreStartNamedEvent struct{}

// 882DEF46A5A7DDEE
type Node_FindChildGameObjectWithName struct{}

// 91CA30F7678B15EF
type Node_UnequipFromCharacterSlot struct{}

// 67F09185C02292EF
type Node_LoadFastTravelWads struct{}

// 912D41C1D8C16CF0
type Node_SetComponentsEnabled struct{}

// 115DF0963B92E2F0
type Node_AddLootResultToWallet struct{}

// 4B798F5D4EEE23F2
type Node_SetAttachmentConfig struct{}

// 6D29E03DF73D41F2
type Node_SetPlayGoRequired struct{}

// 81E90E4EAB478AF9
type Node_SetMarkerClaimed struct{}

// 871F022830B708F5
type Node_StoreCheckpoint struct{}

// B871A4D9BE9FEFF3
type Node_SetPaperBoatPitchRollValues struct{}

// 52CF3FF5AFD416F3
type Node_RelinquishAllPickups struct{}

// 419022E1061F2DF3
type Node_ModifyFocalZoneFacingAngleActivationTolerance struct{}

// 31A9AE98E5FA34F3
type Node_Insert struct{}

// CF50D29EA21DEAF3
type Node_SetCollisionEnabled struct{}

// 002B6D0560A10BF4
type Node_IncrementPickupStage struct{}

// B3FA7525ED7588F4
type Node_RemoveInteractZoneTag struct{}

// 89BC7EAAEBC2CFF4
type Node_SetUDSActivityEnd struct{}

// 246174CD6CB7D2F4
type Node_SetLightsEnabled struct{}

// 73A1AA47ABC3D6F4
type Node_SetTraversePathEnabled struct{}

// 71B639DB54AF51F6
type Node_RegisterForControlSwipeEvent struct{}

// 502EC1D347864BF5
type Node_AbortCreatureInteract struct{}

// 9830AF3A232673F5
type Node_ClearBlackboardVariable struct {
	params [2]Node
	param  uint16
	next   Node
}

// ACA2F86C9167CCF5
type Node_SetParticleEmitterEnabled struct{}

// 64EFB050714EE0F5
type Node_AddToVariable struct{}

// A04DDE040CD640F6
type Node_RecenterToAngle struct{}

// ED7172791C4A52F6
type Node_RelinquishPickupBySlot struct{}

// 9200ECA91D4AE1F6
type Node_AddCreatureActorToSequence struct{}

// 4CB79DDE2B22FBF7
type Node_TriggerWave struct{}

// ADBD64976106C7F8
type Node_SetUDSActivityUnavailable struct{}

// 55090118944CE2FD
type Node_RegisterForWeaponThrown struct{}

// 47E12B02002160FB
type Node_DestroyArrows struct{}

// 8AB1851B268709FA
type Node_DecrementPickupStage struct{}

// 8AD3EF9F220936FA
type Node_SetContextActionsEnabled struct{}

// 4810D4406CCE17FB
type Node_ScentTracker_SetUpdateFrequency struct{}

// 4B674C79E5A41DFB
type Node_SetContextActionTraversalExitSegment struct{}

// 119F4CD3CE2928FB
type Node_DropChainedObject struct{}

// 84CF01EEDEEC8EFB
type Node_DrawNavCurve struct{}

// 7A4232127681BFFB
type Node_SetTimeScale struct{}

// 5432F51B14AD4CFC
type Node_SetContextActionBranchData struct{}

// EAA1FC9B92648BFC
type Node_DrawLine struct{}

// DA2EA6B42CA9F9FC
type Node_SetInteractZonePromptEnabled struct{}

// 2C9625B907D114FF
type Node_RegisterMinibossCheckpoint struct{}

// ABE9385ED79254FE
type Node_RegisterObjectAsProp struct{}

// B8FCC6DD299C97FE
type Node_AlertWaveElement struct{}

// 5882181CE75BD2FE
type Node_GetCreature struct {
	param Node
}

// 1132BFDDC3CFD6FE
type Node_LockOn struct{}

// F5B303D764C3DCFE
type Node_SetShowPlaceholderBoatOnDockDisable struct{}

// C0B46C3F28F02DFF
type Node_SetCompanionWaitOnTraversePath struct{}

// 1D32CC53BBAC71FF
type Node_CineOnly_SwapCreatureToExistingObject struct{}

// 5C0F3DE76B8BBEFF
type Node_LoadCheck struct{}

// E345CC797D57D5FF
type Node_ResetCompass struct{}

// cc60de7cd3b6e663
type Node_Event struct{}

// 269b61b515610c3b
type Node_RegisteredEvent struct{}

// 7cc678873d79f660
type Node_Data struct{}

// e94fc2743690e481
type Node_Triggerable struct{}

// 89f5f4d75ff26610
type Node_ConstantDisplay struct{}

// fea23d13757d05c5
type Node_Constant struct{}

// 856bc88dce81cfd1
type Node_EditorAssetReference struct{}

// c5e296be1dd6ced9
type Node_Struct struct{}

// b5e4f22a23062ff8
type Node_EnumReference struct{}

// 52f178c8c09070cf
type Node_WadId struct{}

// 2c9fd77556119222
type Node_RawUniqueId struct{}

// 9bdca2fde055f622
type Node_RawStringHash struct{}

// 9df4beb458171ca0
type Node_GenericGuidPathReference struct{}

// 0849b07af0eb138e
type Node_GameObjectReference struct{}

// fd79eb0f8d223a5b
type Node_InstancedObjectReference struct{}

// 0d05dbd3dc6c4f76
type Node_InteractZoneReference struct{}

// c02f579701d823bc
type Node_ArrowId struct{}

// a6171fe17d54329e
type Node_ArrowEmitterId struct{}

// 40c6b7e9778f4ac5
type Node_BeamAttackId struct{}

// 95105785b3c257d9
type Node_CameraId struct{}

// cf228e78da60a4a5
type Node_CameraRecenterId struct{}

// e9de8c61895cdaf0
type Node_GatewayMarkerId struct{}

// 4c49a8c78c504caa
type Node_ConcussionId struct{}

// 29e7e9b37903ac59
type Node_DecisionId struct{}

// b49c7bea838afe04
type Node_EquipmentId struct{}

// 88f073d90a625908
type Node_IconId struct{}

// 79a0e934acbd6dfc
type Node_MeterId struct{}

// d62d4572caf7926d
type Node_PickupId struct{}

// d16793f5f58bb9b5
type Node_PickupSlotId struct{}

// fdaad40094d8a839
type Node_QuestId struct{}

// a478db9cd3a09ff6
type Node_RumbleId struct{}

// 2e76833856f86838
type Node_ScreenShakeId struct{}

// 622b77cd8bfdbfb4
type Node_FullScreenEffectId struct{}

// 5a56ed75136d76eb
type Node_BranchId struct{}

// d2b018e91c1a93fa
type Node_MapRealmId struct{}

// ae18acf3bfdcf152
type Node_MapRegionId struct{}

// 178dfc16926e0844
type Node_MapMarkerId struct{}

// d1470e0ce80bc0d8
type Node_MapAreaId struct{}

// cde34ffb67f87d97
type Node_PropId struct{}

// 048dca824cdd6c1d
type Node_TweakConstantId struct{}

// b50eceffb8abd291
type Node_PlayFXId struct{}

// 78852150a54e3f86
type Node_InteractZoneTemplateId struct{}

// efc78cf842cb91c1
type Node_AnimationId struct{}

// ba3cf0f381de832f
type Node_AnimJointReference struct{}

// c61a0e105014e038
type Node_TraversePathReference struct{}

// 492efafdec6d9a69
type Node_TraverseLinkReference struct{}

// 86644edfd60344c5
type Node_ContextActionReference struct{}

// ade3a3c7d5ae113c
type Node_NavCurveReference struct{}

// e8ddee845605f702
type Node_SoundEmitterReference struct{}

// 37615fcda854b0a4
type Node_SoundPortalReference struct{}

// 6b36b9a958046d4c
type Node_SoundProximityTriggerReference struct{}

// 36a9670bf45a9087
type Node_BehaviorTreeRootReference struct{}

// 4885673f08024735
type Node_BehaviorTreeSubtreeReference struct{}

// 23dcc4734469660b
type Node_WwiseEventId struct{}

// a72bb875f329502d
type Node_WwiseSwitchGroupId struct{}

// 78d91938944b6264
type Node_WwiseSwitchStateId struct{}

// 22321a16a46541cf
type Node_WwiseRTPCId struct{}

// d82f895868cbc7d9
type Node_LamsId struct{}

// 1f4bfa75e62f6476
type Node_BanterId struct{}

// 25805601e8a96131
type Node_RecipeId struct{}

// cdc7f7460c8cc34f
type Node_ResourceId struct{}

// a79f65af4efb7ab4
type Node_WalletId struct{}

// 7cf65792f2480290
type Node_WeaponId struct{}

// 47bdaa2fe5e117ba
type Node_WildlifeId struct{}

// 663726d070490221
type Node_HapticId struct{}

// 03920e3a93c6e400
type Node_LootConditionId struct{}

// 0ceeaba96ca34bc7
type Node_LootConditionSetId struct{}

// e41d8772c589f070
type Node_LootDistributorId struct{}

// 4a0ca99b7bf90834
type Node_RelativeReference struct{}

// 220e63265fd2c566
type Node_RelativeGameObjectReference struct{}

// 398229c1bf072da3
type Node_RelativeInteractZoneReference struct{}

// 42c59fcb07f5e71b
type Node_RelativeTraversePathReference struct{}

// c11435c61ac3673b
type Node_RelativeTraverseLinkReference struct{}

// 295d9e74f165fba5
type Node_RelativeContextActionReference struct{}

// 7d7f523a1e1c8c86
type Node_RelativeNavCurveReference struct{}

// 189cb06796275fbc
type Node_RelativeSoundEmitterReference struct{}

// cf1534664b91d850
type Node_RelativeSoundPortalReference struct{}

// 873bd08d804f158e
type Node_RelativeSoundProximityTriggerReference struct{}

// 657ccc60e8b3adb6
type Node_RelativeAnimJoint struct{}

// bc838245895da8de
type Node_AsWadId struct{}

// 1414bbfb8eafb551
type Node_TweaksIdAsStringHash struct{}

// 86de5d27641e2855
type Node_ConcatenateAndHashRawStrings struct{}

// 04ae331e307a3420
type Node_OnFastTravelWadsLoaded struct{}

// 926caeba606562e8
type Node_OnBeamCreated struct{}

// 514d9afff1681639
type Node_OnTimerComplete struct {
	node Node
}

// cb8641b853fbcd9c
type Node_OnFrameDelay struct{}

// 055359a614867cc6
type Node_OnSpawnWildlifeComplete struct{}

// a04a86f51c8e34c6
type Node_OnSpawnComplete struct{}

// 90217b02c282d226
type Node_OnSpawnedObjectFrameUpdate struct{}

// 7f2545f59c5ca5b8
type Node_OnSpawnedObjectFrameUpdateExpired struct{}

// a0261df565152550
type Node_OnSpawnedObjectDestroyed struct{}

// 407762df190d8e8a
type Node_OnSpawnedObjectRecycled struct{}

// dbc491e75c2b829c
type Node_OnSpawnFXComplete struct{}

// 5cd26a7454f3b9bb
type Node_OnMaterialAnimationDone struct{}

// 231fe60f7af11389
type Node_OnAnimationDone struct{}

// 8ddbc2744aa1a33f
type Node_RegisterCombatIndicator struct{}

// 174bf2b7c7b7bc6a
type Node_OnArrowEmitted struct{}

// f0a27848bbaeffc2
type Node_OnLoadCheckFinished struct{}

// af7b779367274516
type Node_OnBlackboardTimerExpired struct{}

// bf881bd3ad19d00e
type Node_OnWarpComplete struct{}

// 2c7a27caaeca647b
type Node_ModifyFocalZoneCameraAngleActivationTolerance struct{}

// 746ee266064e4228
type Node_OnFocalZonePreActivationEnter struct{}

// d49668ece1374231
type Node_OnFocalZonePreActivationExit struct{}

// 515c5a812d44cb8c
type Node_OnBanterStart struct{}

// 875031f40b6c907c
type Node_OnBanterDone struct{}

// 42160ff9e068b081
type Node_OnNextBanterDone struct{}

// 881ef391553740c9
type Node_OnBanterFailed struct{}

// c8da6c2869fe521f
type Node_RemoveEquipmentFromWallet struct{}

// e96efae266d22bb6
type Node_OnUIMessageClosed struct{}

// 374b9370e2741624
type Node_OnCallAndResponseTriggered struct{}

// 78028f7675b302d1
type Node_OnCallAndResponseFinished struct{}

// 19f88e7729f15fec
type Node_OnRegisteredFrameUpdate struct{}

// 8940960461f727b8
type Node_OnAnimFrame struct{}

// c52912a99e3029c2
type Node_OnMovePlay struct{}

// c94dc15319c6652a
type Node_OnHealthChange struct{}

// f3f35e11ad6680c9
type Node_OnDeath struct{}

// bdc033de21601dc5
type Node_OnLookAt struct{}

// 0ea3372bcf3edaf3
type Node_OnButtonPress struct{}

// 784b89e7c35a1b9d
type Node_OnUIButtonPress struct{}

// 87c93b285cf1afd5
type Node_OnControlPress struct{}

// 2c5d0590b77b7db1
type Node_OnControlMash struct{}

// b2efc2bf56af4552
type Node_OnDistanceCheck struct{}

// f0f5c75512f5557d
type Node_OnEventQueueProcessed struct{}

// 18136aec00f11253
type Node_OnRegisteredBreakableBroken struct{}

// 7b0c22f8691e872d
type Node_OnRegisteredMeterChanged struct{}

// 3311c868a305e711
type Node_OnControlSwipe struct{}

// 978d17b5d77e416d
type Node_OnReticleTargetInvalid struct{}

// 1d0b8860ff8e4f75
type Node_OnVariableChanged struct{}

// 7e41456b339534ce
type Node_LoadGate struct{}

// 9a5c017fe16eb2a7
type Node_OnStartGameFromThisWad struct{}

// a6d8ecc2eaaa4512
type Node_ClearAllWeather struct{}

// 078b1630af80bf08
type Node_CreatureSpawnOptionsStruct struct{}

// 3247023efaa55f82
type Node_OnEncounterStart struct{}

// 3e347f0a44f8fc8b
type Node_OnEncounterCreated struct{}

// a331adf5bd27ea6c
type Node_OnEncounterFinished struct{}

// e41883d29611e9ed
type Node_OnEncounterReset struct{}

// b7dcf355d3ae7665
type Node_OnEnemySpawn struct{}

// 6865f5f27fcd9a8b
type Node_OnDynamicCharacterSpawn struct{}

// e273b915d4249fe1
type Node_OnEnemyDeath struct{}

// db4e8cf9e7f209ab
type Node_Crank_Callback struct{}

// a9f5278bcde79690
type Node_Sluice_Callback struct{}

// 97da1f4f120b4345
type Node_Sluice_SoundCallback struct{}

// 49a41bd6dc3e2c6e
type Node_OnInteractStart struct{}

// 695dc97456f1b038
type Node_OnInteractFinished struct{}

// 2e8228b8c3c59c9c
type Node_OnInteractAborted struct{}

// c444784d3307e6c9
type Node_OnInteractDone struct{}

// 71744f76784cc25d
type Node_OnInteractAttemptedWhileOccupied struct{}

// e20e4150a4da67c3
type Node_OnTraverseGraphEntryEvent struct{}

// 14a3081dcbe7d1be
type Node_OnTraverseGraphEnteredEvent struct{}

// 1923576cf1912be1
type Node_OnCreatedVFSEntryChanged struct{}

// 7c088808310b8c8a
type Node_OnCreatedVFSEntryExecuted struct{}

// aefa146d7ad401fc
type Node_OnConcussionHitBase struct{}

// 6e82923ca6494322
type Node_OnArrowImpactBase struct{}

// 90daac765f2d2923
type Node_OnCombatCollisionBase struct{}

// 0a2458fed1a8eea5
type Node_OnWeaponThrown struct{}

// e5d1ab0456da1291
type Node_OnThrownWeaponBlockedBase struct{}

// daae217faa20e2d0
type Node_NamedEventBase struct{}

// 8ccb347d8f347e2a
type Node_OnImmediateEvent struct{}

// e33ecdcf4f6c0eee
type Node_SyncedSequenceActorInfo struct{}

// 0d04000a399016e0
type Node_SyncedSequenceSimpleObjectActorInfo struct{}

// 8be15850bea252ec
type Node_SyncedSequenceCreatureActorInfo struct{}

// d52f5120e4ca98ea
type Node_SyncedSequenceDynamicActorInfo struct{}

// 64378ab079270ae4
type Node_SyncedSequenceSimpleObjectActorInfoSet struct{}

// 3c27b77d71bb6a20
type Node_SyncedSequenceCreatureIdentifier struct{}

// 57ed7db35ce082ca
type Node_SyncedSequenceCreatureActorInfoSet struct{}

// 3582cf8c3c09106c
type Node_MotionWarpParameters struct{}

// b5cd5b784b03b76a
type Node_SyncedSequence struct{}

// 79d6757f054d54f7
type Node_RegisteredSyncedSequence struct{}

// f3ad9196ee051258
type Node_DynamicSyncedSequenceEventContainer struct{}

// 7bd6c2178771e9c4
type Node_SyncedSequenceEvent struct{}

// ff4aef85420d4a3f
type Node_SyncedSequenceOnSkip struct{}

// 4c59ee9fe3f68832
type Node_SyncedSequenceOnSkipExit struct{}

// fa0e90b2e5be7491
type Node_SyncedSequenceOnSequenceExit struct{}

// cc9b2cfc9968d3c9
type Node_SyncedSequenceOnTimeReached struct{}

// 20f7b326fadcb634
type Node_SyncedSequenceOnMoveTimeReached struct{}

// 5a48e5adb73eaf17
type Node_FocalZoneStrafe struct{}

// 8dd68602aba3bfff
type Node_SyncedSequenceTime struct{}

func (n *Node_Header) To_String() string {
	if n.node == nil {
		return "(null)"
	}
	return n.node.To_String()
}

func (n *Node_Dummy) To_String() string {
	return "(invalid)"
}

func (n *Node_Variable) To_String() string {
	return fmt.Sprintf("v_%s", n.str)
}

func (n *Node_SetEffectsSubtitle) To_String() string {
	return "SetEffectsSubtitle()"
}

func (n *Node_ShowUIItemCardMessage) To_String() string {
	return "ShowUIItemCardMessage()"
}

func (n *Node_ShowUIMessage) To_String() string {
	return "ShowUIMessage()"
}

func (n *Node_ShowUISidebarMessage) To_String() string {
	return "ShowUISidebarMessage()"
}

func (n *Node_ShowUISplashScreenMessage) To_String() string {
	return "ShowUISplashScreenMessage()"
}

func (n Node_ContextAction) To_String() string {
	return "ContextAction()"
}

func (n Node_Ceil) To_String() string {
	return "Ceil()"
}

func (n Node_HasInteractZoneTag) To_String() string {
	return "HasInteractZoneTag()"
}

func (n Node_AsLootDistributor) To_String() string {
	return "AsLootDistributor()"
}

func (n Node_IsPinType) To_String() string {
	return "IsPinType()"
}

func (n Node_GetCameraTargetPosition) To_String() string {
	return "GetCameraTargetPosition()"
}

func (n Node_GetActiveMovePercent) To_String() string {
	return "GetActiveMovePercent()"
}

func (n Node_AsBanter) To_String() string {
	return "AsBanter()"
}

func (n Node_InteractZoneTemplate) To_String() string {
	return "InteractZoneTemplate()"
}

func (n Node_GetScale) To_String() string {
	return "GetScale()"
}

func (n Node_GetBlackboardVariable) To_String() string {
	return "GetBlackboardVariable()"
}

func (n Node_Icon) To_String() string {
	return "Icon()"
}

func (n Node_GetFlawlessAresAvailable) To_String() string {
	return "GetFlawlessAresAvailable()"
}

func (n Node_IsInNewGamePlus) To_String() string {
	return "IsInNewGamePlus()"
}

func (n Node_GetSkillTreeTokensOfTypeOn) To_String() string {
	return "GetSkillTreeTokensOfTypeOn()"
}

func (n Node_CanCraftRecipe) To_String() string {
	return "CanCraftRecipe()"
}

func (n Node_IsVolumeEnabled) To_String() string {
	return "IsVolumeEnabled()"
}

func (n Node_WalletHasEquipment) To_String() string {
	return "WalletHasEquipment()"
}

func (n Node_IsSideQuestInAnotherRealm) To_String() string {
	return "IsSideQuestInAnotherRealm()"
}

func (n Node_CheckProgressionFact) To_String() string {
	return "CheckProgressionFact()"
}

func (n Node_WwiseRTPC) To_String() string {
	return "WwiseRTPC()"
}

func (n Node_VectorDot) To_String() string {
	return "VectorDot()"
}

func (n Node_GetVectorFromBlackboard) To_String() string {
	return "GetVectorFromBlackboard()"
}

func (n Node_InstancedTraversePath) To_String() string {
	return "InstancedTraversePath()"
}

func (n Node_GetLastAttacked) To_String() string {
	return "GetLastAttacked()"
}

func (n Node_InstancedContextAction) To_String() string {
	return "InstancedContextAction()"
}

func (n Node_GetClonedWeapons) To_String() string {
	return "GetClonedWeapons()"
}

func (n Node_GetContextActionApproachPoint) To_String() string {
	return "GetContextActionApproachPoint()"
}

func (n Node_IsAnySoundPlaying) To_String() string {
	return "IsAnySoundPlaying()"
}

func (n Node_IsFalling) To_String() string {
	return "IsFalling()"
}

func (n Node_GetPlayersCurrentRegion) To_String() string {
	return "GetPlayersCurrentRegion()"
}

func (n Node_IsControlDown) To_String() string {
	return "IsControlDown()"
}

func (n Node_GetSeekTargetRemainingCount) To_String() string {
	return "GetSeekTargetRemainingCount()"
}

func (n Node_IsInsideFocalZone) To_String() string {
	return "IsInsideFocalZone()"
}

func (n Node_QueryEquipmentTrait) To_String() string {
	return "QueryEquipmentTrait()"
}

func (n Node_GetAnimJoint) To_String() string {
	return "GetAnimJoint()"
}

func (n Node_IsHostile) To_String() string {
	return "IsHostile()"
}

func (n Node_VectorCross) To_String() string {
	return "VectorCross()"
}

func (n Node_GetMeterName) To_String() string {
	return "GetMeterName()"
}

func (n Node_FindRandomPositionOnNavMesh) To_String() string {
	return "FindRandomPositionOnNavMesh()"
}

func (n Node_AsMeter) To_String() string {
	return "AsMeter()"
}

func (n Node_VectorAdd) To_String() string {
	return "VectorAdd()"
}

func (n Node_QuestHasAnyFlags) To_String() string {
	return "QuestHasAnyFlags()"
}

func (n Node_CheckFeatureData) To_String() string {
	return "CheckFeatureData()"
}

func (n Node_GetStickInput) To_String() string {
	return "GetStickInput()"
}

func (n Node_IsInteractZonePromptEnabled) To_String() string {
	return "IsInteractZonePromptEnabled()"
}

func (n Node_CreatureBehaviorSetOptionData) To_String() string {
	return "CreatureBehaviorSetOptionData()"
}

func (n Node_Effect) To_String() string {
	return "Effect()"
}

func (n Node_Subtract) To_String() string {
	return fmt.Sprintf("(%s - %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_PlayerEitherHasOrCanCraftEquipmentRightNow) To_String() string {
	return "PlayerEitherHasOrCanCraftEquipmentRightNow()"
}

func (n Node_AsPlayFX) To_String() string {
	return "AsPlayFX()"
}

func (n Node_GetShieldValue) To_String() string {
	return "GetShieldValue()"
}

func (n Node_Equipment) To_String() string {
	return "Equipment()"
}

func (n Node_IsUnoccupied) To_String() string {
	return "IsUnoccupied()"
}

func (n Node_GetStringHashValueFromElement) To_String() string {
	return "GetStringHashValueFromElement()"
}

func (n Node_IsCollisionEnabled) To_String() string {
	return "IsCollisionEnabled()"
}

func (n Node_IsAttackableArrow) To_String() string {
	return "IsAttackableArrow()"
}

func (n Node_AsBranch) To_String() string {
	return "AsBranch()"
}

func (n Node_GetAnimationLength) To_String() string {
	return "GetAnimationLength()"
}

func (n Node_IsMovementLimitedByCurrentAOO) To_String() string {
	return "IsMovementLimitedByCurrentAOO()"
}

func (n Node_GetFloatValueFromElement) To_String() string {
	return "GetFloatValueFromElement()"
}

func (n Node_GetAvailabilityStateBySource) To_String() string {
	return "GetAvailabilityStateBySource()"
}

func (n Node_GetAllInteractsDisabled) To_String() string {
	return "GetAllInteractsDisabled()"
}

func (n Node_ACos) To_String() string {
	return "ACos()"
}

func (n Node_AsPickupSlot) To_String() string {
	return "AsPickupSlot()"
}

func (n Node_IsContainedInArray) To_String() string {
	return "IsContainedInArray()"
}

func (n Node_MakeVector) To_String() string {
	return "MakeVector()"
}

func (n Node_FullScreenEffect) To_String() string {
	return "FullScreenEffect()"
}

func (n Node_AsWwiseEvent) To_String() string {
	return fmt.Sprintf("AsWwiseEvent(%s)", n.param.To_String())
}

func (n Node_ASin) To_String() string {
	return "ASin()"
}

func (n Node_GetActiveWeapons) To_String() string {
	return "GetActiveWeapons()"
}

func (n Node_AsWallet) To_String() string {
	return "AsWallet()"
}

func (n Node_ArrowEmitter) To_String() string {
	return "ArrowEmitter()"
}

func (n Node_GetOptionControlUpForSetting) To_String() string {
	return "GetOptionControlUpForSetting()"
}

func (n Node_WasWeaponReflected) To_String() string {
	return "WasWeaponReflected()"
}

func (n Node_GetAnimDriverValue) To_String() string {
	return "GetAnimDriverValue()"
}

func (n Node_GetDockTargetType) To_String() string {
	return "GetDockTargetType()"
}

func (n Node_IsDynamicCharacterSpa) To_String() string {
	return "IsDynamicCharacterSpa()"
}

func (n Node_EncounterElementData) To_String() string {
	return "EncounterElementData()"
}

func (n Node_BeamAttack) To_String() string {
	return "BeamAttack()"
}

func (n Node_Rand) To_String() string {
	return "Random()"
}

func (n Node_IsMentalStateEqual) To_String() string {
	return "IsMentalStateEqual()"
}

func (n Node_IsSequenceSkipping) To_String() string {
	return "IsSequenceSkipping()"
}

func (n Node_GetEquipmentInWallet) To_String() string {
	return "GetEquipmentInWallet()"
}

func (n Node_GetRegionForMapMarker) To_String() string {
	return "GetRegionForMapMarker()"
}

func (n Node_GetTyr) To_String() string {
	return "GetTyr()"
}

func (n Node_HasResurrectionStoneBeenUsed) To_String() string {
	return "HasResurrectionStoneBeenUsed()"
}

func (n Node_MakeVector4) To_String() string {
	return "MakeVector4()"
}

func (n Node_GetCameraOrbitLeft) To_String() string {
	return "GetCameraOrbitLeft()"
}

func (n Node_GetVFSInt) To_String() string {
	return "GetVFSInt()"
}

func (n Node_IsValidEquipment) To_String() string {
	return "IsValidEquipment()"
}

func (n Node_IsMarkerClaimed) To_String() string {
	return "IsMarkerClaimed()"
}

func (n Node_GetTouchpadSwipeMapped) To_String() string {
	return "GetTouchpadSwipeMapped()"
}

func (n Node_RecipeHasFlag) To_String() string {
	return "RecipeHasFlag()"
}

func (n Node_GetBreakableStage) To_String() string {
	return "GetBreakableStage()"
}

func (n Node_IsTouchpadSwipeMapped) To_String() string {
	return "IsTouchpadSwipeMapped()"
}

func (n Node_TraverseLink) To_String() string {
	return "TraverseLink()"
}

func (n Node_HasHitFlag) To_String() string {
	return "HasHitFlag()"
}

func (n Node_CanEquipToEquipmentSlot) To_String() string {
	return "CanEquipToEquipmentSlot()"
}

func (n Node_ChooseByCondition) To_String() string {
	return "ChooseByCondition()"
}

func (n Node_Enum) To_String() string {
	return "Enum()"
}

func (n Node_IsSvrContextConnected) To_String() string {
	return "IsSvrContextConnected()"
}

func (n Node_IsTraverseLinkEnabled) To_String() string {
	return "IsTraverseLinkEnabled()"
}

func (n Node_GetEquipmentInCharacterSlot) To_String() string {
	return "GetEquipmentInCharacterSlot()"
}

func (n Node_GetListenerPositionInfo) To_String() string {
	return "GetListenerPositionInfo()"
}

func (n Node_LootCondition) To_String() string {
	return "LootCondition()"
}

func (n Node_MapMarkerHasAllFlags) To_String() string {
	return "MapMarkerHasAllFlags()"
}

func (n Node_IsWeatherCategoryActive) To_String() string {
	return "IsWeatherCategoryActive()"
}

func (n Node_Max) To_String() string {
	return "Max()"
}

func (n Node_GetContextActionFromBlackboard) To_String() string {
	return "GetContextActionFromBlackboard()"
}

func (n Node_GetWolfSledInertia) To_String() string {
	return "GetWolfSledInertia()"
}

func (n Node_AsConcussion) To_String() string {
	return "AsConcussion()"
}

func (n Node_IsPlatform) To_String() string {
	return "IsPlatform()"
}

func (n Node_AsNavCurve) To_String() string {
	return "AsNavCurve()"
}

func (n Node_IsRealmLocked) To_String() string {
	return "IsRealmLocked()"
}

func (n Node_GetCameraOrbit) To_String() string {
	return "GetCameraOrbit()"
}

func (n Node_AsWwiseSwitchGroup) To_String() string {
	return "AsWwiseSwitchGroup()"
}

func (n Node_AsSoundEmitter) To_String() string {
	return "AsSoundEmitter()"
}

func (n Node_GetPlayersCurrentArea) To_String() string {
	return "GetPlayersCurrentArea()"
}

func (n Node_GetScriptOwningWad) To_String() string {
	return "GetScriptOwningWad()"
}

func (n Node_Crank_GetCurrentC) To_String() string {
	return "Crank_GetCurrentC()"
}

func (n Node_GetAvailabilityState) To_String() string {
	return "GetAvailabilityState()"
}

func (n Node_IsRouteToOtherRealmsViaFastTravel) To_String() string {
	return "IsRouteToOtherRealmsViaFastTravel()"
}

func (n Node_SortIterator) To_String() string {
	return "SortIterator()"
}

func (n Node_AsInteractZoneTemplate) To_String() string {
	return "AsInteractZoneTemplate()"
}

func (n Node_GetTimerFromBlackboard) To_String() string {
	return "GetTimerFromBlackboard()"
}

func (n Node_GatewayMarker) To_String() string {
	return "GatewayMarker()"
}

func (n Node_HasPositionalFlag) To_String() string {
	return "HasPositionalFlag()"
}

func (n Node_VectorLength) To_String() string {
	return fmt.Sprintf("VectorLength(%s)", n.vector.To_String())
}

func (n Node_AreNavCurvesEnabledForObjects) To_String() string {
	return "AreNavCurvesEnabledForObjects()"
}

func (n Node_GetCompassTargetPositionLerpSpeed) To_String() string {
	return "GetCompassTargetPositionLerpSpeed()"
}

func (n Node_HitFlagsTo_String) To_String() string {
	return "HitFlagsTo_String()"
}

func (n Node_Add) To_String() string {
	return fmt.Sprintf("(%s + %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_MemoryAvailableToSpawnCreature) To_String() string {
	return "MemoryAvailableToSpawnCreature()"
}

func (n Node_HasAnimationComponent) To_String() string {
	return "HasAnimationComponent()"
}

func (n Node_GetVelocity) To_String() string {
	return "GetVelocity()"
}

func (n Node_Branch) To_String() string {
	return "Branch()"
}

func (n Node_IsPlayGoDownloaded) To_String() string {
	return "IsPlayGoDownloaded()"
}

func (n Node_ChooseByInt) To_String() string {
	return "ChooseByInt()"
}

func (n Node_GetKratos) To_String() string {
	return "GetKratos()"
}

func (n Node_IsPickupSlotOnCooldown) To_String() string {
	return "IsPickupSlotOnCooldown()"
}

func (n Node_GetCompanionReticleHitData) To_String() string {
	return "GetCompanionReticleHitData()"
}

func (n Node_GetCurrentOptionIndexForSetting) To_String() string {
	return "GetCurrentOptionIndexForSetting()"
}

func (n Node_AreBehaviorsEnabled) To_String() string {
	return "AreBehaviorsEnabled()"
}

func (n Node_GetEquipmentInWalletWithFlags) To_String() string {
	return "GetEquipmentInWalletWithFlags()"
}

func (n Node_Sin) To_String() string {
	return "Sin()"
}

func (n Node_GetBreakableHitPoints) To_String() string {
	return "GetBreakableHitPoints()"
}

func (n Node_GetGameObjectFromBlackboard) To_String() string {
	return "GetGameObjectFromBlackboard()"
}

func (n Node_GetGroundLevel) To_String() string {
	return "GetGroundLevel()"
}

func (n Node_GetVisibility) To_String() string {
	return "GetVisibility()"
}

func (n Node_RecipeHasFlags) To_String() string {
	return "RecipeHasFlags()"
}

func (n Node_GetPathDirection) To_String() string {
	return "GetPathDirection()"
}

func (n Node_GetBestNavmeshPosInAOO) To_String() string {
	return "GetBestNavmeshPosInAOO()"
}

func (n Node_GetFightPosition) To_String() string {
	return "GetFightPosition()"
}

func (n Node_GetTimeSinceBanterPlayed) To_String() string {
	return "GetTimeSinceBanterPlayed()"
}

func (n Node_ShuffleIterator) To_String() string {
	return "ShuffleIterator()"
}

func (n Node_IsButtonDown) To_String() string {
	return "IsButtonDown()"
}

func (n Node_AreContextActionsEnabled) To_String() string {
	return "AreContextActionsEnabled()"
}

func (n Node_CreatureHasAOO) To_String() string {
	return "CreatureHasAOO()"
}

func (n Node_AsCreature) To_String() string {
	return "AsCreature()"
}

func (n Node_GetStringLength) To_String() string {
	return "GetStringLength()"
}

func (n Node_LootConditionSet) To_String() string {
	return "LootConditionSet()"
}

func (n Node_SoundPortal) To_String() string {
	return "SoundPortal()"
}

func (n Node_GetClosestFastTravelMarkerPos) To_String() string {
	return "GetClosestFastTravelMarkerPos()"
}

func (n Node_GetQuestName) To_String() string {
	return "GetQuestName()"
}

func (n Node_Weapon) To_String() string {
	return "Weapon()"
}

func (n Node_AsEquipment) To_String() string {
	return "AsEquipment()"
}

func (n Node_GetWadName) To_String() string {
	return "GetWadName()"
}

func (n Node_GetSetting) To_String() string {
	return "GetSetting()"
}

func (n Node_GetSpawnFacing) To_String() string {
	return "GetSpawnFacing()"
}

func (n Node_Or) To_String() string {
	return fmt.Sprintf("(%s || %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_GetPickupStageCount) To_String() string {
	return "GetPickupStageCount()"
}

func (n Node_ScreenShake) To_String() string {
	return "ScreenShake()"
}

func (n Node_IsPickupSlotActive) To_String() string {
	return "IsPickupSlotActive()"
}

func (n Node_GetChildren) To_String() string {
	return "GetChildren()"
}

func (n Node_MapRealm) To_String() string {
	return "MapRealm()"
}

func (n Node_CompassPathDistance) To_String() string {
	return "CompassPathDistance()"
}

func (n Node_InteractZoneLocked) To_String() string {
	return "InteractZoneLocked()"
}

func (n Node_WwiseSwitchState) To_String() string {
	return "WwiseSwitchState()"
}

func (n Node_AsMapArea) To_String() string {
	return "AsMapArea()"
}

func (n Node_IsPositionOnScreen) To_String() string {
	return "IsPositionOnScreen()"
}

func (n Node_EquipmentHasFlags) To_String() string {
	return "EquipmentHasFlags()"
}

func (n Node_GetPlayerBoat) To_String() string {
	return "GetPlayerBoat()"
}

func (n Node_GetQuestProgressAndGoal) To_String() string {
	return "GetQuestProgressAndGoal()"
}

func (n Node_GetObjectsWithFlagsInRadius) To_String() string {
	return "GetObjectsWithFlagsInRadius()"
}

func (n Node_GetMotionDebuggerRecordLocation) To_String() string {
	return "GetMotionDebuggerRecordLocation()"
}

func (n Node_GetVariable) To_String() string {
	return fmt.Sprintf("GetVariable(\"%s\")", n.name)
}

func (n Node_AsAnimation) To_String() string {
	return "AsAnimation()"
}

func (n Node_GetVFSFloat) To_String() string {
	return "GetVFSFloat()"
}

func (n Node_GetCreatureGroun) To_String() string {
	return "GetCreatureGroun()"
}

func (n Node_Numeric) To_String() string {
	return fmt.Sprintf("%f", n.val)
}

func (n Node_IterateFixedSizeArray) To_String() string {
	return "IterateFixedSizeArray()"
}

func (n Node_GetCurrentOptionJoysticForSetting) To_String() string {
	return "GetCurrentOptionJoysticForSetting()"
}

func (n Node_GetExistingResourcesInWallet) To_String() string {
	return "GetExistingResourcesInWallet()"
}

func (n Node_GetSpawnPosition) To_String() string {
	return "GetSpawnPosition()"
}

func (n Node_GetBoolValueFromElement) To_String() string {
	return "GetBoolValueFromElement()"
}

func (n Node_GetTargetCreature) To_String() string {
	return "GetTargetCreature()"
}

func (n Node_AsQuest) To_String() string {
	return "AsQuest()"
}

func (n Node_GetAngleBetweenVectors) To_String() string {
	return "GetAngleBetweenVectors()"
}

func (n Node_IsNavAssistCompassMarker) To_String() string {
	return "IsNavAssistCompassMarker()"
}

func (n Node_IsPlayerDebugMovement) To_String() string {
	return "IsPlayerDebugMovement()"
}

func (n Node_DidCreatureRecentlyHaveFlag) To_String() string {
	return "DidCreatureRecentlyHaveFlag()"
}

func (n Node_GetPendulumAngle) To_String() string {
	return "GetPendulumAngle()"
}

func (n Node_Prop) To_String() string {
	return "Prop()"
}

func (n Node_AsSoundProximityTrigger) To_String() string {
	return "AsSoundProximityTrigger()"
}

func (n Node_GetMaxSpeedOverride) To_String() string {
	return "GetMaxSpeedOverride()"
}

func (n Node_IsInValhalla) To_String() string {
	return "IsInValhalla()"
}

func (n Node_GetWeaponGameObject) To_String() string {
	return "GetWeaponGameObject()"
}

func (n Node_GameObject) To_String() string {
	return "GameObject()"
}

func (n Node_GetCurrentSpline) To_String() string {
	return "GetCurrentSpline()"
}

func (n Node_GetAnimationTime) To_String() string {
	return "GetAnimationTime()"
}

func (n Node_AsMapRealm) To_String() string {
	return "AsMapRealm()"
}

func (n Node_IsSummaryCategoryInMapRegion) To_String() string {
	return "IsSummaryCategoryInMapRegion()"
}

func (n Node_GetMapMarkersInArea) To_String() string {
	return "GetMapMarkersInArea()"
}

func (n Node_GetCollisionRadius) To_String() string {
	return "GetCollisionRadius()"
}

func (n Node_GetMPIconObjectByName) To_String() string {
	return "GetMPIconObjectByName()"
}

func (n Node_GetName) To_String() string {
	return "GetName()"
}

func (n Node_InstancedNavCurve) To_String() string {
	return "InstancedNavCurve()"
}

func (n Node_AsCamera) To_String() string {
	return "AsCamera()"
}

func (n Node_GetWolfSledHarnessVisibility) To_String() string {
	return "GetWolfSledHarnessVisibility()"
}

func (n Node_HasBlackboardVariable) To_String() string {
	return fmt.Sprintf("HasBlackboardVariable(%s,%s,%d)", n.params[0].To_String(), n.params[1].To_String(), n.param)
}

func (n Node_IsPickupSlotFree) To_String() string {
	return "IsPickupSlotFree()"
}

func (n Node_GetAllSoundEmittersOnObject) To_String() string {
	return "GetAllSoundEmittersOnObject()"
}

func (n Node_AsTweakConstant) To_String() string {
	return "AsTweakConstant()"
}

func (n Node_InstancedInteractZone) To_String() string {
	return "InstancedInteractZone()"
}

func (n Node_IsInsideVolume) To_String() string {
	return "IsInsideVolume()"
}

func (n Node_GetCompanion) To_String() string {
	return "GetCompanion()"
}

func (n Node_GetCreatureUnderReticle) To_String() string {
	return "GetCreatureUnderReticle()"
}

func (n Node_GetQuestState) To_String() string {
	return "GetQuestState()"
}

func (n Node_VectorScale) To_String() string {
	return "VectorScale()"
}

func (n Node_VectorSubtract) To_String() string {
	return "VectorSubtract()"
}

func (n Node_ArePhysicsEnabled) To_String() string {
	return "ArePhysicsEnabled()"
}

func (n Node_GetArrowOwner) To_String() string {
	return "GetArrowOwner()"
}

func (n Node_GetCurrentOptionControlDownForSetting) To_String() string {
	return "GetCurrentOptionControlDownForSetting()"
}

func (n Node_ContextActionSequencerStage) To_String() string {
	return "ContextActionSequencerStage()"
}

func (n Node_GetPendulumSpeed) To_String() string {
	return "GetPendulumSpeed()"
}

func (n Node_IsValidReference) To_String() string {
	return "IsValidReference()"
}

func (n Node_GetBranchValueFromElement) To_String() string {
	return "GetBranchValueFromElement()"
}

func (n Node_IsTargetObstructed) To_String() string {
	return "IsTargetObstructed()"
}

func (n Node_AsRecipe) To_String() string {
	return "AsRecipe()"
}

func (n Node_IsPickupActive) To_String() string {
	return "IsPickupActive()"
}

func (n Node_GetPlayerDeathsOnCurrentCheckpoint) To_String() string {
	return "GetPlayerDeathsOnCurrentCheckpoint()"
}

func (n Node_CastToInt32) To_String() string {
	return fmt.Sprintf("CastToInt32(%s)", n.param.To_String())
}

func (n Node_GetWolfSledRopeVisibility) To_String() string {
	return "GetWolfSledRopeVisibility()"
}

func (n Node_CanAim) To_String() string {
	return "CanAim()"
}

func (n Node_Floor) To_String() string {
	return fmt.Sprintf("floor(%s)", n.param.To_String())
}

func (n Node_BitwiseAnd) To_String() string {
	return "BitwiseAnd()"
}

func (n Node_GetAllRecipesInWallet) To_String() string {
	return "GetAllRecipesInWallet()"
}

func (n Node_SoundEmitter) To_String() string {
	return "SoundEmitter()"
}

func (n Node_ContextActionSequencerStageData) To_String() string {
	return "ContextActionSequencerStageData()"
}

func (n Node_GetWeaponThrowStatus) To_String() string {
	return "GetWeaponThrowStatus()"
}

func (n Node_AsWwiseSwitchState) To_String() string {
	return "AsWwiseSwitchState()"
}

func (n Node_IsPlayingMove) To_String() string {
	return fmt.Sprintf("IsPlayingMove(%s,%s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_GetMappedButton) To_String() string {
	return "GetMappedButton()"
}

func (n Node_IsWeatherSystemActive) To_String() string {
	return "IsWeatherSystemActive()"
}

func (n Node_MapMarkerHasAnyFlags) To_String() string {
	return "MapMarkerHasAnyFlags()"
}

func (n Node_HasBeenUsed) To_String() string {
	return "HasBeenUsed()"
}

func (n Node_IsOnNavMesh) To_String() string {
	return "IsOnNavMesh()"
}

func (n Node_GetUniqueEquipmentByName) To_String() string {
	return "GetUniqueEquipmentByName()"
}

func (n Node_Sqrt) To_String() string {
	return "Sqrt()"
}

func (n Node_IsAnyBanterPlaying) To_String() string {
	return "IsAnyBanterPlaying()"
}

func (n Node_GetAnimationFrame) To_String() string {
	return "GetAnimationFrame()"
}

func (n Node_GetRandomPositionInVolume) To_String() string {
	return "GetRandomPositionInVolume()"
}

func (n Node_GetSplinePosition) To_String() string {
	return "GetSplinePosition()"
}

func (n Node_EncounterWaveData) To_String() string {
	return "EncounterWaveData()"
}

func (n Node_GetCameraOrbitForward) To_String() string {
	return "GetCameraOrbitForward()"
}

func (n Node_Bool) To_String() string {
	return fmt.Sprintf("%t", n.val)
}

func (n Node_Abs) To_String() string {
	return "Abs()"
}

func (n Node_GetSpawnedObjectVariable) To_String() string {
	return "GetSpawnedObjectVariable()"
}

func (n Node_GetTweakFloatConstant) To_String() string {
	return fmt.Sprintf("GetTweakFloatConstant(%s,%s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_GetRegionForWad) To_String() string {
	return "GetRegionForWad()"
}

func (n Node_AsWildlife) To_String() string {
	return "AsWildlife()"
}

func (n Node_AsPickup) To_String() string {
	return "AsPickup()"
}

func (n Node_AsAnimJoint) To_String() string {
	return "AsAnimJoint()"
}

func (n Node_IsPlayerOnWolfSled) To_String() string {
	return "IsPlayerOnWolfSled()"
}

func (n Node_Wallet) To_String() string {
	return "Wallet()"
}

func (n Node_QuestHasAllFlags) To_String() string {
	return "QuestHasAllFlags()"
}

func (n Node_GetDepthOfField) To_String() string {
	return "GetDepthOfField()"
}

func (n Node_GetFriendlyCreatures) To_String() string {
	return "GetFriendlyCreatures()"
}

func (n Node_AsColor) To_String() string {
	return "AsColor()"
}

func (n Node_IsMarkerLocked) To_String() string {
	return "IsMarkerLocked()"
}

func (n Node_Divide) To_String() string {
	return fmt.Sprintf("(%s / %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_PointOnPath) To_String() string {
	return "PointOnPath()"
}

func (n Node_GetBitFromVariable) To_String() string {
	return "GetBitFromVariable()"
}

func (n Node_IsHighlightDisableOverrideActive) To_String() string {
	return "IsHighlightDisableOverrideActive()"
}

func (n Node_AsProp) To_String() string {
	return "AsProp()"
}

func (n Node_Pickup) To_String() string {
	return fmt.Sprintf("Pickup(\"%s.%s\",%X,%X)", n.namespace, n.name, n.namespaceHash, n.nameHash)
}

func (n Node_Vector) To_String() string {
	return "Vector()"
}

func (n Node_AsRumble) To_String() string {
	return "AsRumble()"
}

func (n Node_GetIntFromBlackboard) To_String() string {
	return "GetIntFromBlackboard()"
}

func (n Node_IsAreaOfOperationEnabled) To_String() string {
	return "IsAreaOfOperationEnabled()"
}

func (n Node_IsPickupAvailable) To_String() string {
	return "IsPickupAvailable()"
}

func (n Node_AsEnum) To_String() string {
	return "AsEnum()"
}

func (n Node_GetLootObjectResult) To_String() string {
	return "GetLootObjectResult()"
}

func (n Node_IsWeaponEmbedded) To_String() string {
	return "IsWeaponEmbedded()"
}

func (n Node_GetLootConditionByName) To_String() string {
	return "GetLootConditionByName()"
}

func (n Node_GetAngrboda) To_String() string {
	return "GetAngrboda()"
}

func (n Node_InstancedSoundEmitter) To_String() string {
	return "InstancedSoundEmitter()"
}

func (n Node_AsGatewayMarker) To_String() string {
	return "AsGatewayMarker()"
}

func (n Node_IsAOOAssignmentTypeEqualTo) To_String() string {
	return "IsAOOAssignmentTypeEqualTo()"
}

func (n Node_GetCurrentContextAction) To_String() string {
	return "GetCurrentContextAction()"
}

func (n Node_IsSpawnEnemiesEnabled) To_String() string {
	return "IsSpawnEnemiesEnabled()"
}

func (n Node_ATan) To_String() string {
	return "ATan()"
}

func (n Node_CastToFloat) To_String() string {
	return fmt.Sprintf("CastToFloat(%s)", n.param.To_String())
}

func (n Node_GetClosestPositionInVolume) To_String() string {
	return "GetClosestPositionInVolume()"
}

func (n Node_IsDoingSyncMove) To_String() string {
	return "IsDoingSyncMove()"
}

func (n Node_ShouldPointToFastTravelMarker) To_String() string {
	return "ShouldPointToFastTravelMarker()"
}

func (n Node_AreContextActionsEnabledForObjects) To_String() string {
	return "AreContextActionsEnabledForObjects()"
}

func (n Node_FindPath) To_String() string {
	return "FindPath()"
}

func (n Node_GetAtreus) To_String() string {
	return "GetAtreus()"
}

func (n Node_CastToUInt64) To_String() string {
	return fmt.Sprintf("CastToUInt64(%s)", n.param.To_String())
}

func (n Node_GetPlayerLockOnTarget) To_String() string {
	return "GetPlayerLockOnTarget()"
}

func (n Node_IsPlayerInBoat) To_String() string {
	return "IsPlayerInBoat()"
}

func (n Node_Min) To_String() string {
	return "Min()"
}

func (n Node_WasEncounterRunning) To_String() string {
	return "WasEncounterRunning()"
}

func (n Node_CanEquipToCharacterSlot) To_String() string {
	return "CanEquipToCharacterSlot()"
}

func (n Node_GetCreaturesInRadius) To_String() string {
	return fmt.Sprintf("GetCreaturesInRadius(%s,%s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_GetAllSkillsInTree) To_String() string {
	return "GetAllSkillsInTree()"
}

func (n Node_IsWeaponActive) To_String() string {
	return fmt.Sprintf("IsWeaponActive(%s,%s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_GetCurrentVehicleCreature) To_String() string {
	return "GetCurrentVehicleCreature()"
}

func (n Node_AsFullScreenEffect) To_String() string {
	return "AsFullScreenEffect()"
}

func (n Node_GetMapMarkersInRegion) To_String() string {
	return "GetMapMarkersInRegion()"
}

func (n Node_GetClosestFastTravelMarker) To_String() string {
	return "GetClosestFastTravelMarker()"
}

func (n Node_GetGIBlend) To_String() string {
	return "GetGIBlend()"
}

func (n Node_GetSideQuestRealm) To_String() string {
	return "GetSideQuestRealm()"
}

func (n Node_IsJumping) To_String() string {
	return "IsJumping()"
}

func (n Node_IsWolfSledSpawned) To_String() string {
	return "IsWolfSledSpawned()"
}

func (n Node_AsVector) To_String() string {
	return "AsVector()"
}

func (n Node_GetFirstEquippableEquipmentSlot) To_String() string {
	return "GetFirstEquippableEquipmentSlot()"
}

func (n Node_GetEquipmentGeneratorFromHandle) To_String() string {
	return "GetEquipmentGeneratorFromHandle()"
}

func (n Node_GetLookAtTarget) To_String() string {
	return "GetLookAtTarget()"
}

func (n Node_IsPlayerInsideLeashZone) To_String() string {
	return "IsPlayerInsideLeashZone()"
}

func (n Node_GetNavCurveLength) To_String() string {
	return "GetNavCurveLength()"
}

func (n Node_AsGameObject) To_String() string {
	return "AsGameObject()"
}

func (n Node_Rumble) To_String() string {
	return "Rumble()"
}

func (n Node_IsAnyOtherSoundPlaying) To_String() string {
	return "IsAnyOtherSoundPlaying()"
}

func (n Node_CastToBool) To_String() string {
	return fmt.Sprintf("CastToBool(%s)", n.param.To_String())
}

func (n Node_IsAvailabilitySourceRequesting) To_String() string {
	return "IsAvailabilitySourceRequesting()"
}

func (n Node_WalletHasRecipe) To_String() string {
	return "WalletHasRecipe()"
}

func (n Node_BreakVector) To_String() string {
	return "BreakVector()"
}

func (n Node_GetMainQuestRealm) To_String() string {
	return "GetMainQuestRealm()"
}

func (n Node_GetVFSTokenFromPath) To_String() string {
	return "GetVFSTokenFromPath()"
}

func (n Node_GetButtonIndex) To_String() string {
	return "GetButtonIndex()"
}

func (n Node_IsModelEnabled) To_String() string {
	return "IsModelEnabled()"
}

func (n Node_GetOptionJoysticForSetting) To_String() string {
	return "GetOptionJoysticForSetting()"
}

func (n Node_GetPickupStage) To_String() string {
	return "GetPickupStage()"
}

func (n Node_GetWolfSledIsInDrift) To_String() string {
	return "GetWolfSledIsInDrift()"
}

func (n Node_GetOptionControlDownForSetting) To_String() string {
	return "GetOptionControlDownForSetting()"
}

func (n Node_GetBooleanFromBlackboard) To_String() string {
	return "GetBooleanFromBlackboard()"
}

func (n Node_AsSoundPortal) To_String() string {
	return "AsSoundPortal()"
}

func (n Node_CameraRecenter) To_String() string {
	return "CameraRecenter()"
}

func (n Node_And) To_String() string {
	return fmt.Sprintf("(%s && %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_CreatureOptionKeyValuePairData) To_String() string {
	return "CreatureOptionKeyValuePairData()"
}

func (n Node_AreLightsEnabled) To_String() string {
	return "AreLightsEnabled()"
}

func (n Node_StringHash) To_String() string {
	return fmt.Sprintf("StringHash(\"%s\", %X)", n.str, n.hash)
}

func (n Node_Round) To_String() string {
	return "Round()"
}

func (n Node_ResourceHasFlags) To_String() string {
	return "ResourceHasFlags()"
}

func (n Node_GetCameraRender) To_String() string {
	return "GetCameraRender()"
}

func (n Node_AsLootCondition) To_String() string {
	return "AsLootCondition()"
}

func (n Node_GetVFSBool) To_String() string {
	return "GetVFSBool()"
}

func (n Node_AsMapRegion) To_String() string {
	return "AsMapRegion()"
}

func (n Node_IsInteracting) To_String() string {
	return "IsInteracting()"
}

func (n Node_IsPositionInCreatureAOO) To_String() string {
	return "IsPositionInCreatureAOO()"
}

func (n Node_IsInContextAction) To_String() string {
	return "IsInContextAction()"
}

func (n Node_MapAreaHasAnyFlags) To_String() string {
	return "MapAreaHasAnyFlags()"
}

func (n Node_GetWolfSummoningPoints) To_String() string {
	return "GetWolfSummoningPoints()"
}

func (n Node_BitwiseOr) To_String() string {
	return "BitwiseOr()"
}

func (n Node_IsSideQuestPathfindingNonContiguous) To_String() string {
	return "IsSideQuestPathfindingNonContiguous()"
}

func (n Node_GetMapRegionsInRealm) To_String() string {
	return "GetMapRegionsInRealm()"
}

func (n Node_GetElement) To_String() string {
	return "GetElement()"
}

func (n Node_Clamp) To_String() string {
	return "Clamp()"
}

func (n Node_IsDying) To_String() string {
	return "IsDying()"
}

func (n Node_Float) To_String() string {
	return fmt.Sprintf("%f", n.val)
}

func (n Node_GetSkillTreeTokenIn) To_String() string {
	return "GetSkillTreeTokenIn()"
}

func (n Node_GetMusicIntensity) To_String() string {
	return "GetMusicIntensity()"
}

func (n Node_CanAcquireSkill) To_String() string {
	return "CanAcquireSkill()"
}

func (n Node_GetCreatureByName) To_String() string {
	return "GetCreatureByName()"
}

func (n Node_GetEnemiesInRadius) To_String() string {
	return "GetEnemiesInRadius()"
}

func (n Node_NavCurve) To_String() string {
	return "NavCurve()"
}

func (n Node_AreNavCurvesEnabled) To_String() string {
	return "AreNavCurvesEnabled()"
}

func (n Node_GetViewport) To_String() string {
	return "GetViewport()"
}

func (n Node_PartFlagsTo_String) To_String() string {
	return "PartFlagsTo_String()"
}

func (n Node_Resource) To_String() string {
	return "Resource()"
}

func (n Node_GetSplineProgression) To_String() string {
	return "GetSplineProgression()"
}

func (n Node_CreatureHasBehaviorTree) To_String() string {
	return "CreatureHasBehaviorTree()"
}

func (n Node_GetVariableInfo) To_String() string {
	return "GetVariableInfo()"
}

func (n Node_Wildlife) To_String() string {
	return "Wildlife()"
}

func (n Node_IsWadLoaded) To_String() string {
	return "IsWadLoaded()"
}

func (n Node_HashString) To_String() string {
	return "HashString()"
}

func (n Node_CreatureOptionData) To_String() string {
	return "CreatureOptionData()"
}

func (n Node_IsInsideLeashZone) To_String() string {
	return "IsInsideLeashZone()"
}

func (n Node_GameMap_IsItemStateV2) To_String() string {
	return "GameMap_IsItemStateV2()"
}

func (n Node_GetMaterialEntryValue) To_String() string {
	return "GetMaterialEntryValue()"
}

func (n Node_GetAllInteractZones) To_String() string {
	return "GetAllInteractZones()"
}

func (n Node_IsInPlaytest) To_String() string {
	return "IsInPlaytest()"
}

func (n Node_GetYakWaterHeight) To_String() string {
	return "GetYakWaterHeight()"
}

func (n Node_GetFirstChild) To_String() string {
	return "GetFirstChild()"
}

func (n Node_GetIntForContextActionEnum) To_String() string {
	return "GetIntForContextActionEnum()"
}

func (n Node_Haptic) To_String() string {
	return "Haptic()"
}

func (n Node_MapRegion) To_String() string {
	return "MapRegion()"
}

func (n Node_GetBanterDuration) To_String() string {
	return "GetBanterDuration()"
}

func (n Node_GetControlIndex) To_String() string {
	return "GetControlIndex()"
}

func (n Node_Recipe) To_String() string {
	return "Recipe()"
}

func (n Node_ATan2) To_String() string {
	return "ATan2()"
}

func (n Node_IsBreakableBroken) To_String() string {
	return "IsBreakableBroken()"
}

func (n Node_MapRegionHasAnyFlags) To_String() string {
	return "MapRegionHasAnyFlags()"
}

func (n Node_VectorDistance) To_String() string {
	return "VectorDistance()"
}

func (n Node_InstancedSoundProximityTrigger) To_String() string {
	return "InstancedSoundProximityTrigger()"
}

func (n Node_GetCurrentOptionControlUpForSetting) To_String() string {
	return "GetCurrentOptionControlUpForSetting()"
}

func (n Node_GetAssociatedGameObject) To_String() string {
	return "GetAssociatedGameObject()"
}

func (n Node_GreaterThan) To_String() string {
	return fmt.Sprintf("(%s > %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_GetAllResourcesInWallet) To_String() string {
	return "GetAllResourcesInWallet()"
}

func (n Node_GetAngleToTarget) To_String() string {
	return "GetAngleToTarget()"
}

func (n Node_Multiply) To_String() string {
	return fmt.Sprintf("(%s * %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_GetHitPoints) To_String() string {
	return "GetHitPoints()"
}

func (n Node_IsInteractCandidateSetActive) To_String() string {
	return "IsInteractCandidateSetActive()"
}

func (n Node_DebugLams) To_String() string {
	return "DebugLams()"
}

func (n Node_AsDecision) To_String() string {
	return "AsDecision()"
}

func (n Node_GetCurrentOptionEventForSetting) To_String() string {
	return "GetCurrentOptionEventForSetting()"
}

func (n Node_GetCurrentOptionEventModForSetting) To_String() string {
	return "GetCurrentOptionEventModForSetting()"
}

func (n Node_GetOptionEventForSetting) To_String() string {
	return "GetOptionEventForSetting()"
}

func (n Node_String) To_String() string {
	return fmt.Sprintf("\"%s\"", n.str)
}

func (n Node_InteractZone) To_String() string {
	return "InteractZone()"
}

func (n Node_HasCombatSetFlag) To_String() string {
	return "HasCombatSetFlag()"
}

func (n Node_HasMeter) To_String() string {
	return "HasMeter()"
}

func (n Node_AsLams) To_String() string {
	return "AsLams()"
}

func (n Node_GetObjectWithUniqueFlag) To_String() string {
	return "GetObjectWithUniqueFlag()"
}

func (n Node_BehaviorTreeSubtree) To_String() string {
	return "BehaviorTreeSubtree()"
}

func (n Node_Lams) To_String() string {
	return "Lams()"
}

func (n Node_GetProgressionFact) To_String() string {
	return "GetProgressionFact()"
}

func (n Node_GetPlayer) To_String() string {
	return "GetPlayer()"
}

func (n Node_GetStringHashFromBlackboard) To_String() string {
	return "GetStringHashFromBlackboard()"
}

func (n Node_AsContextAction) To_String() string {
	return "AsContextAction()"
}

func (n Node_IsMonitorLookingAt) To_String() string {
	return "IsMonitorLookingAt()"
}

func (n Node_BitwiseNot) To_String() string {
	return "BitwiseNot()"
}

func (n Node_IsTraversePathEnabled) To_String() string {
	return "IsTraversePathEnabled()"
}

func (n Node_HasAnySaveGames) To_String() string {
	return "HasAnySaveGames()"
}

func (n Node_IsJointVisible) To_String() string {
	return "IsJointVisible()"
}

func (n Node_AsFixedSizeArray) To_String() string {
	return "AsFixedSizeArray()"
}

func (n Node_GetSoundEmitterSplineTV) To_String() string {
	return "GetSoundEmitterSplineTV()"
}

func (n Node_Meter) To_String() string {
	return "Meter()"
}

func (n Node_HasHitOrPartFlags) To_String() string {
	return "HasHitOrPartFlags()"
}

func (n Node_Not) To_String() string {
	return fmt.Sprintf("!(%s)", n.param.To_String())
}

func (n Node_IsOnGround) To_String() string {
	return "IsOnGround()"
}

func (n Node_GetRecipesInWalletWithFlags) To_String() string {
	return "GetRecipesInWalletWithFlags()"
}

func (n Node_AsBehaviorTreeRoot) To_String() string {
	return "AsBehaviorTreeRoot()"
}

func (n Node_GetArrayLength) To_String() string {
	return "GetArrayLength()"
}

func (n Node_DoesCreatureHaveLookAtTarget) To_String() string {
	return "DoesCreatureHaveLookAtTarget()"
}

func (n Node_AsHaptic) To_String() string {
	return "AsHaptic()"
}

func (n Node_IsSoundPlaying) To_String() string {
	return "IsSoundPlaying()"
}

func (n Node_GetPosition) To_String() string {
	return fmt.Sprintf("GetPosition(%s)", n.param.To_String())
}

func (n Node_GetMeterMin) To_String() string {
	return "GetMeterMin()"
}

func (n Node_GetGlobalVariable) To_String() string {
	return "GetGlobalVariable()"
}

func (n Node_IsGoldBuild) To_String() string {
	return "IsGoldBuild()"
}

func (n Node_AsStringHash) To_String() string {
	return "AsStringHash()"
}

func (n Node_GetDistanceToTargetCapsule) To_String() string {
	return "GetDistanceToTargetCapsule()"
}

func (n Node_Animation) To_String() string {
	return "Animation()"
}

func (n Node_AsArrow) To_String() string {
	return "AsArrow()"
}

func (n Node_IsEncounterRunning) To_String() string {
	return "IsEncounterRunning()"
}

func (n Node_AsCameraRecenter) To_String() string {
	return "AsCameraRecenter()"
}

func (n Node_AreCreaturesOnSameTeam) To_String() string {
	return "AreCreaturesOnSameTeam()"
}

func (n Node_GetResourceAmountInWallet) To_String() string {
	return "GetResourceAmountInWallet()"
}

func (n Node_AreTraverseGraphsEnabledForObjects) To_String() string {
	return "AreTraverseGraphsEnabledForObjects()"
}

func (n Node_HasFlag) To_String() string {
	return fmt.Sprintf("HasFlag(%s,%s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_BreakVector4) To_String() string {
	return "BreakVector4()"
}

func (n Node_Vector4) To_String() string {
	return "Vector4()"
}

func (n Node_GetCallAndResponseObjectsInRadius) To_String() string {
	return "GetCallAndResponseObjectsInRadius()"
}

func (n Node_IsOnTraverseLink) To_String() string {
	return "IsOnTraverseLink()"
}

func (n Node_MapArea) To_String() string {
	return "MapArea()"
}

func (n Node_GetMapAreasInRegion) To_String() string {
	return "GetMapAreasInRegion()"
}

func (n Node_EquipmentGeneratorHasFlags) To_String() string {
	return "EquipmentGeneratorHasFlags()"
}

func (n Node_GetPickupInSlot) To_String() string {
	return "GetPickupInSlot()"
}

func (n Node_RotationFromEuler) To_String() string {
	return "RotationFromEuler()"
}

func (n Node_AsTraverseLink) To_String() string {
	return "AsTraverseLink()"
}

func (n Node_GetSelectedCreature) To_String() string {
	return "GetSelectedCreature()"
}

func (n Node_PlayFX) To_String() string {
	return "PlayFX()"
}

func (n Node_GetCreatureStatsValue) To_String() string {
	return "GetCreatureStatsValue()"
}

func (n Node_IsDead) To_String() string {
	return "IsDead()"
}

func (n Node_Modulo) To_String() string {
	return "Modulo()"
}

func (n Node_IsUnderSystemCon) To_String() string {
	return "IsUnderSystemCon()"
}

func (n Node_LootDistributor) To_String() string {
	return "LootDistributor()"
}

func (n Node_CreatureSpawnOptions) To_String() string {
	return "CreatureSpawnOptions()"
}

func (n Node_IsMusicPlayingAnything) To_String() string {
	return "IsMusicPlayingAnything()"
}

func (n Node_GetPlayersCurrentRealm) To_String() string {
	return "GetPlayersCurrentRealm()"
}

func (n Node_IsBanterPlayingOnCharacter) To_String() string {
	return "IsBanterPlayingOnCharacter()"
}

func (n Node_Array) To_String() string {
	out := ""
	for i, elem := range n.elems {
		if i > 0 {
			out += ", "
		}
		out += elem.To_String()
	}
	return fmt.Sprintf("[%s]", out)
}

func (n Node_GetMeterMax) To_String() string {
	return "GetMeterMax()"
}

func (n Node_AreTraverseLinksEnabledForObjects) To_String() string {
	return "AreTraverseLinksEnabledForObjects()"
}

func (n Node_CanEquipTokenOfTypeTo) To_String() string {
	return "CanEquipTokenOfTypeTo()"
}

func (n Node_InstancedSoundPortal) To_String() string {
	return "InstancedSoundPortal()"
}

func (n Node_GetFirstEquippableCharacterSlot) To_String() string {
	return "GetFirstEquippableCharacterSlot()"
}

func (n Node_GetContextActionTraversalData) To_String() string {
	return "GetContextActionTraversalData()"
}

func (n Node_Camera) To_String() string {
	return "Camera()"
}

func (n Node_IsAvailableForFollow) To_String() string {
	return "IsAvailableForFollow()"
}

func (n Node_GetLootConditionSetByName) To_String() string {
	return "GetLootConditionSetByName()"
}

func (n Node_IsCreature) To_String() string {
	return "IsCreature()"
}

func (n Node_GetCreaturesSpeedSettingFor) To_String() string {
	return "GetCreaturesSpeedSettingFor()"
}

func (n Node_MapMarker) To_String() string {
	return "MapMarker()"
}

func (n Node_GetNumAvailableGameObjects) To_String() string {
	return "GetNumAvailableGameObjects()"
}

func (n Node_IsAvailableForSplines) To_String() string {
	return "IsAvailableForSplines()"
}

func (n Node_AsIcon) To_String() string {
	return "AsIcon()"
}

func (n Node_IsMainQuestInAnotherRealm) To_String() string {
	return "IsMainQuestInAnotherRealm()"
}

func (n Node_GetClosestPositionOnNavMesh) To_String() string {
	return "GetClosestPositionOnNavMesh()"
}

func (n Node_GetPointOnFX) To_String() string {
	return "GetPointOnFX()"
}

func (n Node_Color) To_String() string {
	return "Color()"
}

func (n Node_AsArrowEmitter) To_String() string {
	return "AsArrowEmitter()"
}

func (n Node_GetEnemyFromEncounter) To_String() string {
	return "GetEnemyFromEncounter()"
}

func (n Node_GetTimeScale) To_String() string {
	return "GetTimeScale()"
}

func (n Node_AsBitfield) To_String() string {
	return "AsBitfield()"
}

func (n Node_IsNavObstacleEnabled) To_String() string {
	return "IsNavObstacleEnabled()"
}

func (n Node_IsEncounterCompleted) To_String() string {
	return "IsEncounterCompleted()"
}

func (n Node_GetCenterOfScreenEnemy) To_String() string {
	return "GetCenterOfScreenEnemy()"
}

func (n Node_InstancedTraverseLink) To_String() string {
	return "InstancedTraverseLink()"
}

func (n Node_IsAnyWeatherSystemActive) To_String() string {
	return "IsAnyWeatherSystemActive()"
}

func (n Node_GetParticleMonsterAliveTotemCount) To_String() string {
	return "GetParticleMonsterAliveTotemCount()"
}

func (n Node_LootResultHasEntry) To_String() string {
	return "LootResultHasEntry()"
}

func (n Node_TraversePath) To_String() string {
	return "TraversePath()"
}

func (n Node_AsScreenShake) To_String() string {
	return "AsScreenShake()"
}

func (n Node_IsFocalZoneLockInEnabled) To_String() string {
	return "IsFocalZoneLockInEnabled()"
}

func (n Node_SimpleStateMachine_GetState) To_String() string {
	return "SimpleStateMachine_GetState()"
}

func (n Node_LessThan) To_String() string {
	return fmt.Sprintf("(%s < %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_GetPlayingBanter) To_String() string {
	return "GetPlayingBanter()"
}

func (n Node_IsFightKnowledgeEnabled) To_String() string {
	return "IsFightKnowledgeEnabled()"
}

func (n Node_MapRegionHasAllFlags) To_String() string {
	return "MapRegionHasAllFlags()"
}

func (n Node_IsContextActionTraversalBehaviorInitialized) To_String() string {
	return "IsContextActionTraversalBehaviorInitialized()"
}

func (n Node_WwiseEvent) To_String() string {
	return fmt.Sprintf("WwiseEvent(%s)", n.name)
}

func (n Node_IsSettingEnabled) To_String() string {
	return "IsSettingEnabled()"
}

func (n Node_IsInNavigationMove) To_String() string {
	return "IsInNavigationMove()"
}

func (n Node_IsAvailableForCombat) To_String() string {
	return "IsAvailableForCombat()"
}

func (n Node_IsControlDisabled) To_String() string {
	return "IsControlDisabled()"
}

func (n Node_GetLastAttacker) To_String() string {
	return "GetLastAttacker()"
}

func (n Node_CanBanterPlay) To_String() string {
	return "CanBanterPlay()"
}

func (n Node_IsQuestInState) To_String() string {
	return "IsQuestInState()"
}

func (n Node_TweakConstant) To_String() string {
	return fmt.Sprintf("TweakConstant(%X,\"%s\")", n.val, n.name)
}

func (n Node_IsMainQuestPathfindingNonContiguous) To_String() string {
	return "IsMainQuestPathfindingNonContiguous()"
}

func (n Node_GetTweakIntegerConstant) To_String() string {
	return fmt.Sprintf("GetTweakIntegerConstant(%s,%s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_AsWeapon) To_String() string {
	return "AsWeapon()"
}

func (n Node_IsAvailableForBanter) To_String() string {
	return "IsAvailableForBanter()"
}

func (n Node_MapAreaHasAllFlags) To_String() string {
	return "MapAreaHasAllFlags()"
}

func (n Node_GetFixedSizeArrayCapacity) To_String() string {
	return "GetFixedSizeArrayCapacity()"
}

func (n Node_InteractZoneEnabled) To_String() string {
	return "InteractZoneEnabled()"
}

func (n Node_IsWithinAngle) To_String() string {
	return "IsWithinAngle()"
}

func (n Node_PickupSlot) To_String() string {
	return "PickupSlot()"
}

func (n Node_IsOnActiveTraversePath) To_String() string {
	return "IsOnActiveTraversePath()"
}

func (n Node_IsPaused) To_String() string {
	return "IsPaused()"
}

func (n Node_IsPickupAcquired) To_String() string {
	out := ""
	for i := 0; i < len(n.params2); i++ {
		if i > 0 {
			out += ","
		}
		out += n.params2[i].To_String()
	}
	return fmt.Sprintf("IsPickupAcquired(%s,%s,%s,[%s])", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), out)
}

func (n Node_AsWad) To_String() string {
	return "AsWad()"
}

func (n Node_AnimJoint) To_String() string {
	return "AnimJoint()"
}

func (n Node_GetActiveWeatherSystems) To_String() string {
	return "GetActiveWeatherSystems()"
}

func (n Node_GetQuest) To_String() string {
	return "GetQuest()"
}

func (n Node_GetSpawnedObjectOwningObject) To_String() string {
	return "GetSpawnedObjectOwningObject()"
}

func (n Node_GetCurrentAnimation) To_String() string {
	return "GetCurrentAnimation()"
}

func (n Node_GetCombatStatus) To_String() string {
	return "GetCombatStatus()"
}

func (n Node_Tan) To_String() string {
	return "Tan()"
}

func (n Node_GetWolfSledIsDriftAllowed) To_String() string {
	return "GetWolfSledIsDriftAllowed()"
}

func (n Node_GetFloatFromBlackboard) To_String() string {
	return "GetFloatFromBlackboard()"
}

func (n Node_GetResourcesInWalletWithFlags) To_String() string {
	return "GetResourcesInWalletWithFlags()"
}

func (n Node_GetCreatureAttributeValue) To_String() string {
	return "GetCreatureAttributeValue()"
}

func (n Node_Banter) To_String() string {
	return "Banter()"
}

func (n Node_GetNavCurveInformation) To_String() string {
	return "GetNavCurveInformation()"
}

func (n Node_FindRandomPositionOnNavMeshInRectangle) To_String() string {
	return "FindRandomPositionOnNavMeshInRectangle()"
}

func (n Node_GetEquipmentInEquipmentSlot) To_String() string {
	return "GetEquipmentInEquipmentSlot()"
}

func (n Node_GetAnimationPlayRate) To_String() string {
	return "GetAnimationPlayRate()"
}

func (n Node_GetValhallaComplete) To_String() string {
	return "GetValhallaComplete()"
}

func (n Node_GetMiniGameplaySkipped) To_String() string {
	return "GetMiniGameplaySkipped()"
}

func (n Node_IsWaveRunning) To_String() string {
	return "IsWaveRunning()"
}

func (n Node_AsBeamAttack) To_String() string {
	return "AsBeamAttack()"
}

func (n Node_IsAvailableForSync) To_String() string {
	return "IsAvailableForSync()"
}

func (n Node_GetNumPointsOnFX) To_String() string {
	return "GetNumPointsOnFX()"
}

func (n Node_GetPickupSlotName) To_String() string {
	return "GetPickupSlotName()"
}

func (n Node_IsParticleEmitterEnabled) To_String() string {
	return "IsParticleEmitterEnabled()"
}

func (n Node_GetValuesFromStageData) To_String() string {
	return "GetValuesFromStageData()"
}

func (n Node_Arrow) To_String() string {
	return "Arrow()"
}

func (n Node_AsVector4) To_String() string {
	return "AsVector4()"
}

func (n Node_GetFlawlessZeusAvailable) To_String() string {
	return "GetFlawlessZeusAvailable()"
}

func (n Node_GetNearestSoundPortal) To_String() string {
	return "GetNearestSoundPortal()"
}

func (n Node_IsMusicPlaying) To_String() string {
	return "IsMusicPlaying()"
}

func (n Node_BehaviorTreeRoot) To_String() string {
	return "BehaviorTreeRoot()"
}

func (n Node_GetIntValueFromElement) To_String() string {
	return "GetIntValueFromElement()"
}

func (n Node_AsWwiseRTPC) To_String() string {
	return "AsWwiseRTPC()"
}

func (n Node_IsVisualScriptEnabled) To_String() string {
	return "IsVisualScriptEnabled()"
}

func (n Node_GetGameObjectValueFromElement) To_String() string {
	return "GetGameObjectValueFromElement()"
}

func (n Node_GetGameTime) To_String() string {
	return "GetGameTime()"
}

func (n Node_AsBehaviorTreeSubtree) To_String() string {
	return "AsBehaviorTreeSubtree()"
}

func (n Node_GetParent) To_String() string {
	return "GetParent()"
}

func (n Node_GetCreatureContext) To_String() string {
	return "GetCreatureContext()"
}

func (n Node_Int32) To_String() string {
	return fmt.Sprintf("%d", n.val)
}

func (n Node_GetGameObjectWad) To_String() string {
	return "GetGameObjectWad()"
}

func (n Node_WwiseSwitchGroup) To_String() string {
	return "WwiseSwitchGroup()"
}

func (n Node_AsTraversePath) To_String() string {
	return "AsTraversePath()"
}

func (n Node_SoundProximityTrigger) To_String() string {
	return "SoundProximityTrigger()"
}

func (n Node_ChooseByValue) To_String() string {
	return "ChooseByValue()"
}

func (n Node_GetNumBreakableStages) To_String() string {
	return "GetNumBreakableStages()"
}

func (n Node_GetSpawnElementUID) To_String() string {
	return "GetSpawnElementUID()"
}

func (n Node_Wad) To_String() string {
	return "Wad()"
}

func (n Node_ConcatenateAndHashStrings) To_String() string {
	return "ConcatenateAndHashStrings()"
}

func (n Node_GetBitFromGlobalVariable) To_String() string {
	return "GetBitFromGlobalVariable()"
}

func (n Node_AsResource) To_String() string {
	return "AsResource()"
}

func (n Node_Creature) To_String() string {
	return "Creature()"
}

func (n Node_GetOptionEventModForSetting) To_String() string {
	return "GetOptionEventModForSetting()"
}

func (n Node_AreParticlesEnabled) To_String() string {
	return "AreParticlesEnabled()"
}

func (n Node_Decision) To_String() string {
	return "Decision()"
}

func (n Node_IsQuestTracked) To_String() string {
	return "IsQuestTracked()"
}

func (n Node_GetSpawnedObjectHierarchyRoot) To_String() string {
	return "GetSpawnedObjectHierarchyRoot()"
}

func (n Node_IsSummaryCategoryInMapArea) To_String() string {
	return "IsSummaryCategoryInMapArea()"
}

func (n Node_VectorNormalize) To_String() string {
	return "VectorNormalize()"
}

func (n Node_AreInteractZonesEnabledForObjects) To_String() string {
	return "AreInteractZonesEnabledForObjects()"
}

func (n Node_UInt64) To_String() string {
	return fmt.Sprintf("%d", n.val)
}

func (n Node_Cos) To_String() string {
	return "Cos()"
}

func (n Node_GetMeterValue) To_String() string {
	return "GetMeterValue()"
}

func (n Node_AsLootConditionSet) To_String() string {
	return "AsLootConditionSet()"
}

func (n Node_Concussion) To_String() string {
	return "Concussion()"
}

func (n Node_GetQuestPrimaryCompletionFact) To_String() string {
	return "GetQuestPrimaryCompletionFact()"
}

func (n Node_Equals) To_String() string {
	return fmt.Sprintf("(%s == %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_Quest) To_String() string {
	return "Quest()"
}

func (n Node_IsSkillAcquired) To_String() string {
	return "IsSkillAcquired()"
}

func (n Node_CheckProgressionFactGreaterEquals) To_String() string {
	return "CheckProgressionFactGreaterEquals()"
}

func (n Node_GetNameHash) To_String() string {
	return "GetNameHash()"
}

func (n Node_RotateVector) To_String() string {
	return "RotateVector()"
}

func (n Node_GetRotation) To_String() string {
	return "GetRotation()"
}

func (n Node_AsMapMarker) To_String() string {
	return "AsMapMarker()"
}

func (n Node_IsInAir) To_String() string {
	return "IsInAir()"
}

func (n Node_ResourceHasFlag) To_String() string {
	return "ResourceHasFlag()"
}

func (n Node_AsInteractZone) To_String() string {
	return "AsInteractZone()"
}

func (n Node_ControlClothSim) To_String() string {
	return "ControlClothSim()"
}

func (n Node_DrawPoint) To_String() string {
	return "DrawPoint()"
}

func (n Node_MarkResurrectionStoneUsed) To_String() string {
	return "MarkResurrectionStoneUsed()"
}

func (n Node_ClearWeatherCategory) To_String() string {
	return "ClearWeatherCategory()"
}

func (n Node_SetAnimationPlayRate) To_String() string {
	return "SetAnimationPlayRate()"
}

func (n Node_SetDisableAllInteracts) To_String() string {
	return "SetDisableAllInteracts()"
}

func (n Node_AddExcludeTraverseLinkFilterToCreature) To_String() string {
	return "AddExcludeTraverseLinkFilterToCreature()"
}

func (n Node_PlayCombatConcussion) To_String() string {
	return "PlayCombatConcussion()"
}

func (n Node_SetHapticInstanceParameter) To_String() string {
	return "SetHapticInstanceParameter()"
}

func (n Node_RequestCinematicModeEnabled) To_String() string {
	return "RequestCinematicModeEnabled()"
}

func (n Node_Sluice_StartWater) To_String() string {
	return "Sluice_StartWater()"
}

func (n Node_StopRotatingObjectTowardsTarget) To_String() string {
	return "StopRotatingObjectTowardsTarget()"
}

func (n Node_RegisterOnEncounterFinished) To_String() string {
	return "RegisterOnEncounterFinished()"
}

func (n Node_SetTraversePathFlagEnabled) To_String() string {
	return "SetTraversePathFlagEnabled()"
}

func (n Node_SetSplineLeashing) To_String() string {
	return "SetSplineLeashing()"
}

func (n Node_DrawLine2D) To_String() string {
	return "DrawLine2D()"
}

func (n Node_SetWolfSledAutoRunSplineAngle) To_String() string {
	return "SetWolfSledAutoRunSplineAngle()"
}

func (n Node_Crank_Enable) To_String() string {
	return "Crank_Enable()"
}

func (n Node_CacheValuesOnStack) To_String() string {
	return "CacheValuesOnStack()"
}

func (n Node_SetAggressivePriorityOverride) To_String() string {
	return "SetAggressivePriorityOverride()"
}

func (n Node_Crank_UpdateDrivenObjectAnim) To_String() string {
	return "Crank_UpdateDrivenObjectAnim()"
}

func (n Node_LootPot_Reroll) To_String() string {
	return "LootPot_Reroll()"
}

func (n Node_DisableBoatForceTurnAroundControlMode) To_String() string {
	return "DisableBoatForceTurnAroundControlMode()"
}

func (n Node_AddRecipeToWallet) To_String() string {
	return "AddRecipeToWallet()"
}

func (n Node_RemoveTraverseLinkFilter) To_String() string {
	return "RemoveTraverseLinkFilter()"
}

func (n Node_StopAllSounds) To_String() string {
	return "StopAllSounds()"
}

func (n Node_SimpleStateMachine_ClearStateMac) To_String() string {
	return "SimpleStateMachine_ClearStateMac()"
}

func (n Node_SendTelemetryEvent) To_String() string {
	return "SendTelemetryEvent()"
}

func (n Node_SetSplineFollowDistances) To_String() string {
	return "SetSplineFollowDistances()"
}

func (n Node_ClearFinishedEncounterData) To_String() string {
	return "ClearFinishedEncounterData()"
}

func (n Node_ClearPlayerNavCurve) To_String() string {
	return "ClearPlayerNavCurve()"
}

func (n Node_SetNavObstacleEnabled) To_String() string {
	return "SetNavObstacleEnabled()"
}

func (n Node_SetPreventSoftSave) To_String() string {
	return "SetPreventSoftSave()"
}

func (n Node_FindInteractZoneWithName) To_String() string {
	return "FindInteractZoneWithName()"
}

func (n Node_SubtractFromGlobalVariable) To_String() string {
	return "SubtractFromGlobalVariable()"
}

func (n Node_CheckCineModeIntegrity) To_String() string {
	return "CheckCineModeIntegrity()"
}

func (n Node_InterruptBanterOnCharacter) To_String() string {
	return "InterruptBanterOnCharacter()"
}

func (n Node_SetContextActionApproachData) To_String() string {
	return "SetContextActionApproachData()"
}

func (n Node_ForceCompanionTraverseLink) To_String() string {
	return "ForceCompanionTraverseLink()"
}

func (n Node_SendImmediateEvent) To_String() string {
	return "SendImmediateEvent()"
}

func (n Node_SpawnNPC) To_String() string {
	return "SpawnNPC()"
}

func (n Node_SetBehaviorTreeUpdatePriority) To_String() string {
	return "SetBehaviorTreeUpdatePriority()"
}

func (n Node_CreateVFSExec) To_String() string {
	return "CreateVFSExec()"
}

func (n Node_RemoveRecipeFromWallet) To_String() string {
	return "RemoveRecipeFromWallet()"
}

func (n Node_EndCreatureInteract) To_String() string {
	return "EndCreatureInteract()"
}

func (n Node_SendTelemetryEventDLC) To_String() string {
	return "SendTelemetryEventDLC()"
}

func (n Node_SetNavCurveReversal) To_String() string {
	return "SetNavCurveReversal()"
}

func (n Node_SetJointHighlightCategory) To_String() string {
	return "SetJointHighlightCategory()"
}

func (n Node_AlertWave) To_String() string {
	return "AlertWave()"
}

func (n Node_DrawArc) To_String() string {
	return "DrawArc()"
}

func (n Node_DismissEffectsSubtitle) To_String() string {
	return "DismissEffectsSubtitle()"
}

func (n Node_ModifyFocalZoneCameraAngleDeactivationThreshold) To_String() string {
	return "ModifyFocalZoneCameraAngleDeactivationThreshold()"
}

func (n Node_SetChainVisible) To_String() string {
	return "SetChainVisible()"
}

func (n Node_AddFlag) To_String() string {
	return "AddFlag()"
}

func (n Node_DistributeLootResult) To_String() string {
	return "DistributeLootResult()"
}

func (n Node_ClearFocalZones) To_String() string {
	return "ClearFocalZones()"
}

func (n Node_SetPhysicsEnabled) To_String() string {
	return "SetPhysicsEnabled()"
}

func (n Node_UnmapTouchpadSwipe) To_String() string {
	return "UnmapTouchpadSwipe()"
}

func (n Node_StartYakEnter) To_String() string {
	return "StartYakEnter()"
}

func (n Node_GetClosestElementToPlayer) To_String() string {
	return "GetClosestElementToPlayer()"
}

func (n Node_TransferWalletContents) To_String() string {
	return "TransferWalletContents()"
}

func (n Node_SetVFSBool) To_String() string {
	return "SetVFSBool()"
}

func (n Node_Switch) To_String() string {
	return "Switch()"
}

func (n Node_DebugHighlightObject) To_String() string {
	return "DebugHighlightObject()"
}

func (n Node_SendBehaviorTreeEvent) To_String() string {
	return "SendBehaviorTreeEvent()"
}

func (n Node_RemoveByValue) To_String() string {
	return "RemoveByValue()"
}

func (n Node_SetEnableCounterWind) To_String() string {
	return "SetEnableCounterWind()"
}

func (n Node_GrantLootCondition) To_String() string {
	return "GrantLootCondition()"
}

func (n Node_SetCombatStatus) To_String() string {
	return "SetCombatStatus()"
}

func (n Node_ClearAllAnimationCallbacks) To_String() string {
	return "ClearAllAnimationCallbacks()"
}

func (n Node_MarkCurrentWeaponModeForPlayer) To_String() string {
	return "MarkCurrentWeaponModeForPlayer()"
}

func (n Node_SetBehaviorsEnabled) To_String() string {
	return "SetBehaviorsEnabled()"
}

func (n Node_ResetListenerPosition) To_String() string {
	return "ResetListenerPosition()"
}

func (n Node_SetAOOAssignment) To_String() string {
	return "SetAOOAssignment()"
}

func (n Node_CreateEquipment) To_String() string {
	return "CreateEquipment()"
}

func (n Node_MapTouchpadSwipe) To_String() string {
	return "MapTouchpadSwipe()"
}

func (n Node_ActivateWeatherCategory) To_String() string {
	return "ActivateWeatherCategory()"
}

func (n Node_TriggerBlenderEffect) To_String() string {
	return "TriggerBlenderEffect()"
}

func (n Node_SetCompassTargetPositionLerpSpeed) To_String() string {
	return "SetCompassTargetPositionLerpSpeed()"
}

func (n Node_SetUDSActivityAvailable) To_String() string {
	return "SetUDSActivityAvailable()"
}

func (n Node_SetJointVisible) To_String() string {
	return "SetJointVisible()"
}

func (n Node_Crank_Unlock) To_String() string {
	return "Crank_Unlock()"
}

func (n Node_Sluice_WaterPourStopped) To_String() string {
	return "Sluice_WaterPourStopped()"
}

func (n Node_SetContextActionMentalStateTag) To_String() string {
	return "SetContextActionMentalStateTag()"
}

func (n Node_SetCombatIndicator) To_String() string {
	return "SetCombatIndicator()"
}

func (n Node_ForceCompanionTraversePath) To_String() string {
	return "ForceCompanionTraversePath()"
}

func (n Node_SetOverrideSpeedControlData) To_String() string {
	return "SetOverrideSpeedControlData()"
}

func (n Node_ActivateBoatNavCurveFunctionalityFunneling) To_String() string {
	return "ActivateBoatNavCurveFunctionalityFunneling()"
}

func (n Node_ClearVariable) To_String() string {
	if n.next != nil {
		return fmt.Sprintf("ClearVariable(%s);%s", n.name, n.next.To_String())
	}
	return fmt.Sprintf("ClearVariable(%s)", n.name)
}

func (n Node_ToggleBitOnVariable) To_String() string {
	return "ToggleBitOnVariable()"
}

func (n Node_StartCheckpointedMusic) To_String() string {
	return "StartCheckpointedMusic()"
}

func (n Node_ClearPreventSoftSave) To_String() string {
	return "ClearPreventSoftSave()"
}

func (n Node_SetNavAssistCompassMarker) To_String() string {
	return "SetNavAssistCompassMarker()"
}

func (n Node_ClearAllDisabledControls) To_String() string {
	return "ClearAllDisabledControls()"
}

func (n Node_SetMusicCaption) To_String() string {
	return "SetMusicCaption()"
}

func (n Node_StartTimer) To_String() string {
	if n.next != nil {
		return fmt.Sprintf("StartTimer(%s, %s);%s", n.params[0].To_String(), n.params[1].To_String(), n.next.To_String())
	}
	return fmt.Sprintf("StartTimer(%s, %s)", n.params[0].To_String(), n.params[1].To_String())
}

func (n Node_SetSplineAvoidDistances) To_String() string {
	return "SetSplineAvoidDistances()"
}

func (n Node_IncrementMeterValue) To_String() string {
	return "IncrementMeterValue()"
}

func (n Node_OverridePowerLevelForEncounterElement) To_String() string {
	return "OverridePowerLevelForEncounterElement()"
}

func (n Node_ClearWeather) To_String() string {
	return "ClearWeather()"
}

func (n Node_StartMusicSting) To_String() string {
	return "StartMusicSting()"
}

func (n Node_GrantLootConditionSet) To_String() string {
	return "GrantLootConditionSet()"
}

func (n Node_SetInteractZoneLocked) To_String() string {
	return "SetInteractZoneLocked()"
}

func (n Node_CreatePendulum) To_String() string {
	return "CreatePendulum()"
}

func (n Node_SetBoatIntoTriggeredLaunchState) To_String() string {
	return "SetBoatIntoTriggeredLaunchState()"
}

func (n Node_LootPot_SetLootType) To_String() string {
	return "LootPot_SetLootType()"
}

func (n Node_StopRestartableSoundLoop) To_String() string {
	return "StopRestartableSoundLoop()"
}

func (n Node_SetOverridePositionOnDock) To_String() string {
	return "SetOverridePositionOnDock()"
}

func (n Node_SetAnimDriverValue) To_String() string {
	return "SetAnimDriverValue()"
}

func (n Node_ClearForceLookAt) To_String() string {
	return "ClearForceLookAt()"
}

func (n Node_SetDynamicNavmeshEnabled) To_String() string {
	return "SetDynamicNavmeshEnabled()"
}

func (n Node_SetWolfSledAutoRunAcceleration) To_String() string {
	return "SetWolfSledAutoRunAcceleration()"
}

func (n Node_SetAIUseTraverseLinksWithoutPath) To_String() string {
	return "SetAIUseTraverseLinksWithoutPath()"
}

func (n Node_StartCreatureInteract) To_String() string {
	return "StartCreatureInteract()"
}

func (n Node_SetContextActionFactionTag) To_String() string {
	return "SetContextActionFactionTag()"
}

func (n Node_AcquireSkill) To_String() string {
	return "AcquireSkill()"
}

func (n Node_RemoveCreatureAvailability) To_String() string {
	return "RemoveCreatureAvailability()"
}

func (n Node_SetTraversePathsEnabled) To_String() string {
	return "SetTraversePathsEnabled()"
}

func (n Node_RecenterToFacing) To_String() string {
	return "RecenterToFacing()"
}

func (n Node_StopTimer) To_String() string {
	return fmt.Sprintf("StopTimer(%s)", n.param.To_String())
}

func (n Node_ResetCompassPathfinding) To_String() string {
	return "ResetCompassPathfinding()"
}

func (n Node_ActivatePickup) To_String() string {
	return "ActivatePickup()"
}

func (n Node_KillPlayer) To_String() string {
	return "KillPlayer()"
}

func (n Node_EndPaperboatSync) To_String() string {
	return "EndPaperboatSync()"
}

func (n Node_WolfSledSplineSectionEnter) To_String() string {
	return "WolfSledSplineSectionEnter()"
}

func (n Node_SetTraverseLinkOverridePrompt) To_String() string {
	return "SetTraverseLinkOverridePrompt()"
}

func (n Node_RevertAggressivePriorityOverride) To_String() string {
	return "RevertAggressivePriorityOverride()"
}

func (n Node_ClearVisualVariantOnDock) To_String() string {
	return "ClearVisualVariantOnDock()"
}

func (n Node_PrepareDynamicSyncedSequence) To_String() string {
	return "PrepareDynamicSyncedSequence()"
}

func (n Node_OverrideCreatureLootDistributor) To_String() string {
	return "OverrideCreatureLootDistributor()"
}

func (n Node_Sequence) To_String() string {
	out := ""
	for i, node := range n.nodes {
		if i > 0 {
			out += ";"
		}
		if node != nil {
			out += node.To_String()
		} else {
			out += "(invalid)"
		}
	}
	return fmt.Sprintf("%s", out)
}

func (n Node_SetBreakableEnabled) To_String() string {
	return "SetBreakableEnabled()"
}

func (n Node_DrawCircle) To_String() string {
	return "DrawCircle()"
}

func (n Node_RequestDespawnDynamicCharacter) To_String() string {
	return "RequestDespawnDynamicCharacter()"
}

func (n Node_AddCharacterFX) To_String() string {
	return "AddCharacterFX()"
}

func (n Node_ClearOverrideSpeedControlData) To_String() string {
	return "ClearOverrideSpeedControlData()"
}

func (n Node_SetWaitGateEnabledState) To_String() string {
	return "SetWaitGateEnabledState()"
}

func (n Node_SimpleStateMachine_SwitchState) To_String() string {
	return "SimpleStateMachine_SwitchState()"
}

func (n Node_SetContextActionTraversalCompanionLead) To_String() string {
	return "SetContextActionTraversalCompanionLead()"
}

func (n Node_SetClonePositionRotation) To_String() string {
	return "SetClonePositionRotation()"
}

func (n Node_SetPaperBoatOffsetSpringValues) To_String() string {
	return "SetPaperBoatOffsetSpringValues()"
}

func (n Node_EndSyncedSequenceByActor) To_String() string {
	return "EndSyncedSequenceByActor()"
}

func (n Node_ProfileBegin) To_String() string {
	return "ProfileBegin()"
}

func (n Node_SwitchOnPinType) To_String() string {
	return "SwitchOnPinType()"
}

func (n Node_StopAndSettlePendulum) To_String() string {
	return "StopAndSettlePendulum()"
}

func (n Node_SetWolfSledAutoRunStrafeMultiplier) To_String() string {
	return "SetWolfSledAutoRunStrafeMultiplier()"
}

func (n Node_DrawZone) To_String() string {
	return "DrawZone()"
}

func (n Node_DeactivateBoatNavCurveFunctionality) To_String() string {
	return "DeactivateBoatNavCurveFunctionality()"
}

func (n Node_SetSoundProximityTriggerEnabled) To_String() string {
	return "SetSoundProximityTriggerEnabled()"
}

func (n Node_DebugSpawner) To_String() string {
	return "DebugSpawner()"
}

func (n Node_SetContextActionTraversalSegmentDiffThreshold) To_String() string {
	return "SetContextActionTraversalSegmentDiffThreshold()"
}

func (n Node_NotifySluiceIceChanged) To_String() string {
	return "NotifySluiceIceChanged()"
}

func (n Node_HideUIMessage) To_String() string {
	return "HideUIMessage()"
}

func (n Node_AddAreaOfOperation) To_String() string {
	return "AddAreaOfOperation()"
}

func (n Node_SendContainerRewardAcquiredTelemetry) To_String() string {
	return "SendContainerRewardAcquiredTelemetry()"
}

func (n Node_ShowUIMessageWithTokensBase) To_String() string {
	return "ShowUIMessageWithTokensBase()"
}

func (n Node_SetInfiniteSpawning) To_String() string {
	return "SetInfiniteSpawning()"
}

func (n Node_FixBreakable) To_String() string {
	return "FixBreakable()"
}

func (n Node_SetMusicWwiseSwitch) To_String() string {
	return "SetMusicWwiseSwitch()"
}

func (n Node_OverridePowerLevelForEncounter) To_String() string {
	return "OverridePowerLevelForEncounter()"
}

func (n Node_CreateVFSInt) To_String() string {
	return "CreateVFSInt()"
}

func (n Node_WildlifeDestroy) To_String() string {
	return "WildlifeDestroy()"
}

func (n Node_StartCamera) To_String() string {
	return "StartCamera()"
}

func (n Node_SetPendulumPosition) To_String() string {
	return "SetPendulumPosition()"
}

func (n Node_ClearParticleMonsterCheckpointing) To_String() string {
	return "ClearParticleMonsterCheckpointing()"
}

func (n Node_SetEnableAreaOfOperation) To_String() string {
	return "SetEnableAreaOfOperation()"
}

func (n Node_Sluice_RemoveManualBlocker) To_String() string {
	return "Sluice_RemoveManualBlocker()"
}

func (n Node_RegisterSpawnedObjectForFrameUpdate) To_String() string {
	return "RegisterSpawnedObjectForFrameUpdate()"
}

func (n Node_RemovePropToSyncWithCAOwner) To_String() string {
	return "RemovePropToSyncWithCAOwner()"
}

func (n Node_RegisterForDistanceEvent) To_String() string {
	return "RegisterForDistanceEvent()"
}

func (n Node_AddInteractZoneTag) To_String() string {
	return "AddInteractZoneTag()"
}

func (n Node_RemoveIncludeTraverseLinkFilterFromCreature) To_String() string {
	return "RemoveIncludeTraverseLinkFilterFromCreature()"
}

func (n Node_SetContextActionInitialized) To_String() string {
	return "SetContextActionInitialized()"
}

func (n Node_TeleportWolfSled) To_String() string {
	return "TeleportWolfSled()"
}

func (n Node_ModifyNavCurveProgressionLimit) To_String() string {
	return "ModifyNavCurveProgressionLimit()"
}

func (n Node_SetWolfSledAutoRunSpeedModifier) To_String() string {
	return "SetWolfSledAutoRunSpeedModifier()"
}

func (n Node_GrantCreatureLoot) To_String() string {
	return "GrantCreatureLoot()"
}

func (n Node_QuitAndJumpToWad) To_String() string {
	return "QuitAndJumpToWad()"
}

func (n Node_HideAllUIMessages) To_String() string {
	return "HideAllUIMessages()"
}

func (n Node_AddCreatureToCustomSplineGroup) To_String() string {
	return "AddCreatureToCustomSplineGroup()"
}

func (n Node_SetContextActionTraversalDistanceThreshold) To_String() string {
	return "SetContextActionTraversalDistanceThreshold()"
}

func (n Node_SetPendulumRotation) To_String() string {
	return "SetPendulumRotation()"
}

func (n Node_ClearOverridePositionOnDock) To_String() string {
	return "ClearOverridePositionOnDock()"
}

func (n Node_RemoveContextActionCreatureTag) To_String() string {
	return "RemoveContextActionCreatureTag()"
}

func (n Node_RegisterBoatDock) To_String() string {
	return "RegisterBoatDock()"
}

func (n Node_SetWolfSledHarnessVisibility) To_String() string {
	return "SetWolfSledHarnessVisibility()"
}

func (n Node_CreateVFSBool) To_String() string {
	return "CreateVFSBool()"
}

func (n Node_DEBUG_ONLY_SetQuestStateAndBackpropagate) To_String() string {
	return "DEBUG_ONLY_SetQuestStateAndBackpropagate()"
}

func (n Node_For) To_String() string {
	return "For()"
}

func (n Node_ModifyFocalZoneLockInEnabled) To_String() string {
	return "ModifyFocalZoneLockInEnabled()"
}

func (n Node_SendEnvironment) To_String() string {
	return "SendEnvironment()"
}

func (n Node_ResetEncounter) To_String() string {
	return "ResetEncounter()"
}

func (n Node_CrankSetup) To_String() string {
	return "CrankSetup()"
}

func (n Node_SetPosition) To_String() string {
	return "SetPosition()"
}

func (n Node_Crank_ForcePlayerDetach) To_String() string {
	return "Crank_ForcePlayerDetach()"
}

func (n Node_SetCheckpointedMusic) To_String() string {
	return "SetCheckpointedMusic()"
}

func (n Node_AttachZiplineToSpear) To_String() string {
	return "AttachZiplineToSpear()"
}

func (n Node_SetSplineAttributes) To_String() string {
	return "SetSplineAttributes()"
}

func (n Node_PauseEncounter) To_String() string {
	return "PauseEncounter()"
}

func (n Node_SetEncounterLODRange) To_String() string {
	return "SetEncounterLODRange()"
}

func (n Node_RegisterWaitGateSetup) To_String() string {
	return "RegisterWaitGateSetup()"
}

func (n Node_SummonWolfSled) To_String() string {
	return "SummonWolfSled()"
}

func (n Node_SetPickupStage) To_String() string {
	if n.next != nil {
		return fmt.Sprintf("SetPickupStage(%s,%s,%s,%s);%s", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), n.params[3].To_String(), n.next.To_String())
	}
	return fmt.Sprintf("SetPickupStage(%s,%s,%s,%s)", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), n.params[3].To_String())
}

func (n Node_RestoreMarkedWeaponModeForPlayer) To_String() string {
	return "RestoreMarkedWeaponModeForPlayer()"
}

func (n Node_SpawnVehicle) To_String() string {
	return "SpawnVehicle()"
}

func (n Node_DestroyBeam) To_String() string {
	return "DestroyBeam()"
}

func (n Node_SetMaterialEntryValue) To_String() string {
	return "SetMaterialEntryValue()"
}

func (n Node_AnimationSync) To_String() string {
	return "AnimationSync()"
}

func (n Node_RequestSpawnDynamicCharacter) To_String() string {
	return "RequestSpawnDynamicCharacter()"
}

func (n Node_SetControlEnabledWithType) To_String() string {
	return "SetControlEnabledWithType()"
}

func (n Node_SetControlEnabled) To_String() string {
	return "SetControlEnabled()"
}

func (n Node_PauseCurrentAnimation) To_String() string {
	return "PauseCurrentAnimation()"
}

func (n Node_GetTriggerIndex) To_String() string {
	return "GetTriggerIndex()"
}

func (n Node_ClearForcedSpline) To_String() string {
	return "ClearForcedSpline()"
}

func (n Node_UpdateGameObjectSpline) To_String() string {
	return "UpdateGameObjectSpline()"
}

func (n Node_SetDynamicCharacterSlot) To_String() string {
	return "SetDynamicCharacterSlot()"
}

func (n Node_ModifyFocalZonePriority) To_String() string {
	return "ModifyFocalZonePriority()"
}

func (n Node_SetInteractZoneHintSubPrompt) To_String() string {
	return "SetInteractZoneHintSubPrompt()"
}

func (n Node_StopSound) To_String() string {
	return fmt.Sprintf("StopSound(%s,%s,%s,%s,%s,%s)", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), n.params[3].To_String(), n.params[4].To_String(), n.params[5].To_String())
}

func (n Node_RaycastCollision) To_String() string {
	return "RaycastCollision()"
}

func (n Node_ClearOverrideEnterBranchesOnDock) To_String() string {
	return "ClearOverrideEnterBranchesOnDock()"
}

func (n Node_UnpauseEncounter) To_String() string {
	return "UnpauseEncounter()"
}

func (n Node_Sluice_StopWater) To_String() string {
	return "Sluice_StopWater()"
}

func (n Node_Crank_Lock) To_String() string {
	return "Crank_Lock()"
}

func (n Node_EndSyncedSequenceByName) To_String() string {
	return "EndSyncedSequenceByName()"
}

func (n Node_EndCinematicSkip) To_String() string {
	return "EndCinematicSkip()"
}

func (n Node_StartHapticVibration) To_String() string {
	return "StartHapticVibration()"
}

func (n Node_EnableSetting) To_String() string {
	return "EnableSetting()"
}

func (n Node_DiscoverSummaryCategoryInAllRegions) To_String() string {
	return "DiscoverSummaryCategoryInAllRegions()"
}

func (n Node_AddEmbedWeaponToGameObject) To_String() string {
	return "AddEmbedWeaponToGameObject()"
}

func (n Node_DisableBoatForceDirectionControlMode) To_String() string {
	return "DisableBoatForceDirectionControlMode()"
}

func (n Node_ShowUIMessageBase) To_String() string {
	return "ShowUIMessageBase()"
}

func (n Node_ModifyResourceInWallet) To_String() string {
	return "ModifyResourceInWallet()"
}

func (n Node_Crank_API) To_String() string {
	return "Crank_API()"
}

func (n Node_DrawSphere) To_String() string {
	return "DrawSphere()"
}

func (n Node_TransitionBetweenFocalZones) To_String() string {
	return "TransitionBetweenFocalZones()"
}

func (n Node_MarkUsed) To_String() string {
	return "MarkUsed()"
}

func (n Node_RemoveFocalZone) To_String() string {
	return "RemoveFocalZone()"
}

func (n Node_ModifyFocalZoneMovementPriorityOverrideWhenLockedIn) To_String() string {
	return "ModifyFocalZoneMovementPriorityOverrideWhenLockedIn()"
}

func (n Node_SetOverrideVisualVariantOnDock) To_String() string {
	return "SetOverrideVisualVariantOnDock()"
}

func (n Node_EnableLoadGate) To_String() string {
	return "EnableLoadGate()"
}

func (n Node_SetCompassGatewayMarkerIsOpen) To_String() string {
	return "SetCompassGatewayMarkerIsOpen()"
}

func (n Node_OverrideCreatureLootConditionSet) To_String() string {
	return "OverrideCreatureLootConditionSet()"
}

func (n Node_ActivateBoatNavCurveFunctionality) To_String() string {
	return "ActivateBoatNavCurveFunctionality()"
}

func (n Node_ResetHeightField) To_String() string {
	return "ResetHeightField()"
}

func (n Node_GetInteractLerpObject) To_String() string {
	return "GetInteractLerpObject()"
}

func (n Node_PushPendulum) To_String() string {
	return "PushPendulum()"
}

func (n Node_UnregisterTraverseGraphEntryEvent) To_String() string {
	return "UnregisterTraverseGraphEntryEvent()"
}

func (n Node_SetWolfSledIsDriftAllowed) To_String() string {
	return "SetWolfSledIsDriftAllowed()"
}

func (n Node_RemoveThrowableTarget) To_String() string {
	return "RemoveThrowableTarget()"
}

func (n Node_ApplyStatusMeterDamage) To_String() string {
	return "ApplyStatusMeterDamage()"
}

func (n Node_FreezePendulum) To_String() string {
	return "FreezePendulum()"
}

func (n Node_DiscoverSingleRegionSummaryCategory) To_String() string {
	return "DiscoverSingleRegionSummaryCategory()"
}

func (n Node_SetRouteToOtherRealmsViaFastTravel) To_String() string {
	return "SetRouteToOtherRealmsViaFastTravel()"
}

func (n Node_CreateBeam) To_String() string {
	return "CreateBeam()"
}

func (n Node_RemoveAreaOfOperation) To_String() string {
	return "RemoveAreaOfOperation()"
}

func (n Node_BreakBreakable) To_String() string {
	return "BreakBreakable()"
}

func (n Node_PlayBanter) To_String() string {
	return "PlayBanter()"
}

func (n Node_ModifyFocalZoneEnablePitchOnHeadTracking) To_String() string {
	return "ModifyFocalZoneEnablePitchOnHeadTracking()"
}

func (n Node_UnregisterMinibossCheckpoint) To_String() string {
	return "UnregisterMinibossCheckpoint()"
}

func (n Node_TriggerPendulumSlowdown) To_String() string {
	return "TriggerPendulumSlowdown()"
}

func (n Node_StartNewDLCTelemetryRun) To_String() string {
	return "StartNewDLCTelemetryRun()"
}

func (n Node_SetMeterValue) To_String() string {
	return "SetMeterValue()"
}

func (n Node_SetManualSaveEnabled) To_String() string {
	return "SetManualSaveEnabled()"
}

func (n Node_MarkSkipAIUpdate) To_String() string {
	return "MarkSkipAIUpdate()"
}

func (n Node_SetTraverseLinkPromptOffset) To_String() string {
	return "SetTraverseLinkPromptOffset()"
}

func (n Node_AddCombatCollision) To_String() string {
	return "AddCombatCollision()"
}

func (n Node_EquipToEquipmentSlot) To_String() string {
	return "EquipToEquipmentSlot()"
}

func (n Node_SetupPaperboatSync) To_String() string {
	return "SetupPaperboatSync()"
}

func (n Node_UsePlayerHeightForFloating) To_String() string {
	return "UsePlayerHeightForFloating()"
}

func (n Node_SendNamedEventBase) To_String() string {
	return "SendNamedEventBase()"
}

func (n Node_CreateEmptyLootResult) To_String() string {
	return "CreateEmptyLootResult()"
}

func (n Node_RegisterForBreakableBrokenEvent) To_String() string {
	return "RegisterForBreakableBrokenEvent()"
}

func (n Node_ActivateContextAction) To_String() string {
	return "ActivateContextAction()"
}

func (n Node_SetScale) To_String() string {
	return "SetScale()"
}

func (n Node_Crank_JamCrank) To_String() string {
	return "Crank_JamCrank()"
}

func (n Node_SetPreventCinematicSkip) To_String() string {
	return "SetPreventCinematicSkip()"
}

func (n Node_SetContextActionMultiCreatureData) To_String() string {
	return "SetContextActionMultiCreatureData()"
}

func (n Node_SetGlobalVariable) To_String() string {
	return "SetGlobalVariable()"
}

func (n Node_DestroyVFSEntry) To_String() string {
	return "DestroyVFSEntry()"
}

func (n Node_SetTraversePathPromptEnabled) To_String() string {
	return "SetTraversePathPromptEnabled()"
}

func (n Node_OverrideCreatureCullingFadeDistance) To_String() string {
	return "OverrideCreatureCullingFadeDistance()"
}

func (n Node_SetFightKnowledgeEnabled) To_String() string {
	return "SetFightKnowledgeEnabled()"
}

func (n Node_WolfSledForceEnter) To_String() string {
	return "WolfSledForceEnter()"
}

func (n Node_ClearWolfSledSpeedControlOverride) To_String() string {
	return "ClearWolfSledSpeedControlOverride()"
}

func (n Node_SetWolfSledRopeVisibility) To_String() string {
	return "SetWolfSledRopeVisibility()"
}

func (n Node_SendNamedEvent) To_String() string {
	return "SendNamedEvent()"
}

func (n Node_PositionObjectsForSyncedSequence) To_String() string {
	return "PositionObjectsForSyncedSequence()"
}

func (n Node_UnequipFromEquipmentSlot) To_String() string {
	return "UnequipFromEquipmentSlot()"
}

func (n Node_DespawnVehicle) To_String() string {
	return "DespawnVehicle()"
}

func (n Node_SimpleStateMachine_Clea) To_String() string {
	return "SimpleStateMachine_Clea()"
}

func (n Node_RegisterTraverseGraphEntryEvent) To_String() string {
	return "RegisterTraverseGraphEntryEvent()"
}

func (n Node_RemoveChild) To_String() string {
	return "RemoveChild()"
}

func (n Node_DrawDebugTable) To_String() string {
	return "DrawDebugTable()"
}

func (n Node_Crank_Disable) To_String() string {
	return "Crank_Disable()"
}

func (n Node_ActivateWeather) To_String() string {
	return "ActivateWeather()"
}

func (n Node_AddGameObjectToProfileGroup) To_String() string {
	return "AddGameObjectToProfileGroup()"
}

func (n Node_SetZiplineSlack) To_String() string {
	return "SetZiplineSlack()"
}

func (n Node_TriggerQuestManualActiveBehaviors) To_String() string {
	return "TriggerQuestManualActiveBehaviors()"
}

func (n Node_SetPaperBoatFadeIn) To_String() string {
	return "SetPaperBoatFadeIn()"
}

func (n Node_Print) To_String() string {
	if n.next != nil {
		return fmt.Sprintf("Print(%s);%s", n.param.To_String(), n.next.To_String())
	}
	return fmt.Sprintf("Print(%s)", n.param.To_String())
}

func (n Node_SetWolfSledAutoRun) To_String() string {
	return "SetWolfSledAutoRun()"
}

func (n Node_ResetAccesibilityToggleState) To_String() string {
	return "ResetAccesibilityToggleState()"
}

func (n Node_SpawnVehicleAtPosition) To_String() string {
	return "SpawnVehicleAtPosition()"
}

func (n Node_OverrideCreatureMotionParams) To_String() string {
	return "OverrideCreatureMotionParams()"
}

func (n Node_SetMarkerLocked) To_String() string {
	return "SetMarkerLocked()"
}

func (n Node_PrestreamRequest) To_String() string {
	return "PrestreamRequest()"
}

func (n Node_TriggerEncounterEnd) To_String() string {
	return "TriggerEncounterEnd()"
}

func (n Node_BreakLoop) To_String() string {
	return "BreakLoop()"
}

func (n Node_AddSpawnToEncounter) To_String() string {
	return "AddSpawnToEncounter()"
}

func (n Node_SetSplineSpeed) To_String() string {
	return "SetSplineSpeed()"
}

func (n Node_AddPointToFX) To_String() string {
	return "AddPointToFX()"
}

func (n Node_ClearCallAndResponseObject) To_String() string {
	return "ClearCallAndResponseObject()"
}

func (n Node_OpenSoundPortal) To_String() string {
	return "OpenSoundPortal()"
}

func (n Node_SetPlayerControlledCreature) To_String() string {
	return "SetPlayerControlledCreature()"
}

func (n Node_SetWwiseRTPCValue) To_String() string {
	return "SetWwiseRTPCValue()"
}

func (n Node_GetBranchStartPositionAndOrientation) To_String() string {
	return "GetBranchStartPositionAndOrientation()"
}

func (n Node_SetWolfSledAutoRunSlopeBoostMultiplier) To_String() string {
	return "SetWolfSledAutoRunSlopeBoostMultiplier()"
}

func (n Node_SetSoundPortalOpen) To_String() string {
	return "SetSoundPortalOpen()"
}

func (n Node_StartPaperBoatBounceFadeOut) To_String() string {
	return "StartPaperBoatBounceFadeOut()"
}

func (n Node_SetContextActionEncounterMarkerIDTag) To_String() string {
	return "SetContextActionEncounterMarkerIDTag()"
}

func (n Node_DrawRect2D) To_String() string {
	return "DrawRect2D()"
}

func (n Node_SetInfluenceConeAttributes) To_String() string {
	return "SetInfluenceConeAttributes()"
}

func (n Node_SetPlayerNavCurve) To_String() string {
	return "SetPlayerNavCurve()"
}

func (n Node_DestroyEquipment) To_String() string {
	return "DestroyEquipment()"
}

func (n Node_SetFlawlessZeusAvailable) To_String() string {
	return "SetFlawlessZeusAvailable()"
}

func (n Node_OverrideLookAtConversationSettings) To_String() string {
	return "OverrideLookAtConversationSettings()"
}

func (n Node_SetTextOnTextObject) To_String() string {
	return "SetTextOnTextObject()"
}

func (n Node_SetYakInteractOptions) To_String() string {
	return "SetYakInteractOptions()"
}

func (n Node_GameMap_SetItemStateV2) To_String() string {
	return "GameMap_SetItemStateV2()"
}

func (n Node_TriggerCallAndResponse) To_String() string {
	return "TriggerCallAndResponse()"
}

func (n Node_CheckFeature) To_String() string {
	return "CheckFeature()"
}

func (n Node_CoilChainedObject) To_String() string {
	return "CoilChainedObject()"
}

func (n Node_ModifyFocalZoneFacingAngleDeactivationThreshold) To_String() string {
	return "ModifyFocalZoneFacingAngleDeactivationThreshold()"
}

func (n Node_StartWolfSledExit) To_String() string {
	return "StartWolfSledExit()"
}

func (n Node_SetSlowdown) To_String() string {
	return "SetSlowdown()"
}

func (n Node_InterruptAllBanter) To_String() string {
	return "InterruptAllBanter()"
}

func (n Node_Sluice_AddManualBlocker) To_String() string {
	return "Sluice_AddManualBlocker()"
}

func (n Node_StopCamera) To_String() string {
	return "StopCamera()"
}

func (n Node_SetWolfSledAutoRunRequiresInput) To_String() string {
	return "SetWolfSledAutoRunRequiresInput()"
}

func (n Node_RegisterForControlMashEvent) To_String() string {
	return "RegisterForControlMashEvent()"
}

func (n Node_UnpauseCurrentAnimation) To_String() string {
	return "UnpauseCurrentAnimation()"
}

func (n Node_SetCreatureAvailabilityRequest) To_String() string {
	return "SetCreatureAvailabilityRequest()"
}

func (n Node_SetCreatureFocus) To_String() string {
	return "SetCreatureFocus()"
}

func (n Node_ShowHiddenWads) To_String() string {
	return "ShowHiddenWads()"
}

func (n Node_OverrideCreatureLootCategory) To_String() string {
	return "OverrideCreatureLootCategory()"
}

func (n Node_ClearAllWeaponsCinematicMode) To_String() string {
	return "ClearAllWeaponsCinematicMode()"
}

func (n Node_ModifyFocalZoneAttachObject) To_String() string {
	return "ModifyFocalZoneAttachObject()"
}

func (n Node_SetBitOnGlobalVariable) To_String() string {
	return "SetBitOnGlobalVariable()"
}

func (n Node_ClearAndHideAllNavAssistCompassMarkers) To_String() string {
	return "ClearAndHideAllNavAssistCompassMarkers()"
}

func (n Node_ClearCreatureFocus) To_String() string {
	return "ClearCreatureFocus()"
}

func (n Node_RemoveSecondaryActorFromSequence) To_String() string {
	return "RemoveSecondaryActorFromSequence()"
}

func (n Node_RegisterSpawnedObjectForVariables) To_String() string {
	return "RegisterSpawnedObjectForVariables()"
}

func (n Node_PrepareSyncedSequence) To_String() string {
	return "PrepareSyncedSequence()"
}

func (n Node_SetBoatDockInteractRadius) To_String() string {
	return "SetBoatDockInteractRadius()"
}

func (n Node_RegisterAsChainedObject) To_String() string {
	return "RegisterAsChainedObject()"
}

func (n Node_EnableContextAction) To_String() string {
	return "EnableContextAction()"
}

func (n Node_SimpleStateMachine_SwitchStateUsingEnumPin) To_String() string {
	return "SimpleStateMachine_SwitchStateUsingEnumPin()"
}

func (n Node_ForEach) To_String() string {
	if n.next != nil {
		return fmt.Sprintf("ForEach(%s,function(v_itr){%s});%s", n.param.To_String(), n.operation.To_String(), n.next.To_String())
	}
	return fmt.Sprintf("ForEach(%s,function(v_itr){%s})", n.param.To_String(), n.operation.To_String())
}

func (n Node_SetBanterFact) To_String() string {
	return "SetBanterFact()"
}

func (n Node_SetWeaponMode) To_String() string {
	return "SetWeaponMode()"
}

func (n Node_SetWolfSledAutoRunControlRange) To_String() string {
	return "SetWolfSledAutoRunControlRange()"
}

func (n Node_TransferLootResultEntry) To_String() string {
	return "TransferLootResultEntry()"
}

func (n Node_SwitchIntRange) To_String() string {
	return "SwitchIntRange()"
}

func (n Node_DropThrowable) To_String() string {
	return "DropThrowable()"
}

func (n Node_RegisterForMeterEvent) To_String() string {
	return "RegisterForMeterEvent()"
}

func (n Node_SoundEmitterEnabled) To_String() string {
	return "SoundEmitterEnabled()"
}

func (n Node_SetFlawlessAresAvailable) To_String() string {
	return "SetFlawlessAresAvailable()"
}

func (n Node_SetContextActionBasicData) To_String() string {
	return "SetContextActionBasicData()"
}

func (n Node_FakeInteraction) To_String() string {
	return "FakeInteraction()"
}

func (n Node_SetSpawnedObjectVariable) To_String() string {
	return "SetSpawnedObjectVariable()"
}

func (n Node_StartBoatSyncMove) To_String() string {
	return "StartBoatSyncMove()"
}

func (n Node_SpawnWildlife) To_String() string {
	return "SpawnWildlife()"
}

func (n Node_SetPointOnFX) To_String() string {
	return "SetPointOnFX()"
}

func (n Node_GetNextArrayValue) To_String() string {
	return "GetNextArrayValue()"
}

func (n Node_ProcessInteractByCandidateSet) To_String() string {
	return "ProcessInteractByCandidateSet()"
}

func (n Node_RegisterForBlackboardTimerExpired) To_String() string {
	return "RegisterForBlackboardTimerExpired()"
}

func (n Node_ClearHideHealthBar) To_String() string {
	return "ClearHideHealthBar()"
}

func (n Node_SetTraversePathPromptLOSEnabled) To_String() string {
	return "SetTraversePathPromptLOSEnabled()"
}

func (n Node_TriggerCamera) To_String() string {
	return "TriggerCamera()"
}

func (n Node_CacheCallAndResponseObject) To_String() string {
	return "CacheCallAndResponseObject()"
}

func (n Node_SetPropEnabled) To_String() string {
	return "SetPropEnabled()"
}

func (n Node_ResetStickDamping) To_String() string {
	return "ResetStickDamping()"
}

func (n Node_SetResourceDebugAddMode) To_String() string {
	return "SetResourceDebugAddMode()"
}

func (n Node_SetUDSActivityStart) To_String() string {
	return "SetUDSActivityStart()"
}

func (n Node_AddNewLookAtTarget) To_String() string {
	return "AddNewLookAtTarget()"
}

func (n Node_RegisterAsZipline) To_String() string {
	return "RegisterAsZipline()"
}

func (n Node_SetAccelerationOverride) To_String() string {
	return "SetAccelerationOverride()"
}

func (n Node_CustomSortIterator) To_String() string {
	return "CustomSortIterator()"
}

func (n Node_SetSplineAvoidAvoidancePreferences) To_String() string {
	return "SetSplineAvoidAvoidancePreferences()"
}

func (n Node_UnregisterForMonitorEvent) To_String() string {
	return "UnregisterForMonitorEvent()"
}

func (n Node_RegisterForMovePlayEvent) To_String() string {
	return "RegisterForMovePlayEvent()"
}

func (n Node_AddContextActionCreatureTag) To_String() string {
	return "AddContextActionCreatureTag()"
}

func (n Node_ClearPreventCheckpoint) To_String() string {
	return "ClearPreventCheckpoint()"
}

func (n Node_SetValhallaComplete) To_String() string {
	return "SetValhallaComplete()"
}

func (n Node_SimpleCombatSync) To_String() string {
	return "SimpleCombatSync()"
}

func (n Node_SetBlackboardTimer) To_String() string {
	if n.next != nil {
		return fmt.Sprintf("SetBlackboardTimer(%s,%s,%s,%s,%d);%s", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), n.params[3].To_String(), n.param, n.next.To_String())
	}
	return fmt.Sprintf("SetBlackboardTimer(%s,%s,%s,%s,%d)", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), n.params[3].To_String(), n.param)
}

func (n Node_ParticleMonsterCheckpointTotemState) To_String() string {
	return "ParticleMonsterCheckpointTotemState()"
}

func (n Node_WildlifeFlee) To_String() string {
	return "WildlifeFlee()"
}

func (n Node_OverrideListenerPosition) To_String() string {
	return "OverrideListenerPosition()"
}

func (n Node_SetRaceBackHomeEnableSledControl) To_String() string {
	return "SetRaceBackHomeEnableSledControl()"
}

func (n Node_SetForceOddNumTraverseLinkConnections) To_String() string {
	return "SetForceOddNumTraverseLinkConnections()"
}

func (n Node_StartYakExit) To_String() string {
	return "StartYakExit()"
}

func (n Node_ForceEngageContextAction) To_String() string {
	return "ForceEngageContextAction()"
}

func (n Node_SetCombatTarget) To_String() string {
	return "SetCombatTarget()"
}

func (n Node_SetPaperBoatSpeed) To_String() string {
	return "SetPaperBoatSpeed()"
}

func (n Node_RegisterTraverseGraphEnteredEvent) To_String() string {
	return "RegisterTraverseGraphEnteredEvent()"
}

func (n Node_ClearBoatDockInteractRadius) To_String() string {
	return "ClearBoatDockInteractRadius()"
}

func (n Node_SetVisibility) To_String() string {
	return "SetVisibility()"
}

func (n Node_GetAdditionalGravityOnPendulum) To_String() string {
	return "GetAdditionalGravityOnPendulum()"
}

func (n Node_SetObjectHighlightCategory) To_String() string {
	return "SetObjectHighlightCategory()"
}

func (n Node_OverrideCreatureCullingRadius) To_String() string {
	return "OverrideCreatureCullingRadius()"
}

func (n Node_SendUIEvent) To_String() string {
	return "SendUIEvent()"
}

func (n Node_ToggleBitOnGlobalVariable) To_String() string {
	return "ToggleBitOnGlobalVariable()"
}

func (n Node_RemovePositionalFlag) To_String() string {
	return "RemovePositionalFlag()"
}

func (n Node_UntrackAllSideQuests) To_String() string {
	return "UntrackAllSideQuests()"
}

func (n Node_ClearAllDisabledControlsWithType) To_String() string {
	return "ClearAllDisabledControlsWithType()"
}

func (n Node_AddThrowableTarget) To_String() string {
	return "AddThrowableTarget()"
}

func (n Node_SetInteractZoneEnabled) To_String() string {
	return "SetInteractZoneEnabled()"
}

func (n Node_PrepareRegisteredSyncedSequence) To_String() string {
	return "PrepareRegisteredSyncedSequence()"
}

func (n Node_SetBitOnVariable) To_String() string {
	return "SetBitOnVariable()"
}

func (n Node_OverridePowerLevelForEncounterWave) To_String() string {
	return "OverridePowerLevelForEncounterWave()"
}

func (n Node_SetTraversePathDirectionEnabled) To_String() string {
	return "SetTraversePathDirectionEnabled()"
}

func (n Node_SetDockEnabled) To_String() string {
	return "SetDockEnabled()"
}

func (n Node_SetCharacterFX) To_String() string {
	return "SetCharacterFX()"
}

func (n Node_PlayMaterialAnim) To_String() string {
	return "PlayMaterialAnim()"
}

func (n Node_SetSubtitleDistanceThresholds) To_String() string {
	return "SetSubtitleDistanceThresholds()"
}

func (n Node_WolfSledSplineSectionExit) To_String() string {
	return "WolfSledSplineSectionExit()"
}

func (n Node_SetFlightVolumeEnabledByType) To_String() string {
	return "SetFlightVolumeEnabledByType()"
}

func (n Node_CreateEncounterData) To_String() string {
	return "CreateEncounterData()"
}

func (n Node_SetSubtitleDistanceThresholdsAlternate) To_String() string {
	return "SetSubtitleDistanceThresholdsAlternate()"
}

func (n Node_ClearEquipmentPreview) To_String() string {
	return "ClearEquipmentPreview()"
}

func (n Node_ClearMaxSpeedOverride) To_String() string {
	return "ClearMaxSpeedOverride()"
}

func (n Node_SetCreatureHighlightCategoryOverride) To_String() string {
	return "SetCreatureHighlightCategoryOverride()"
}

func (n Node_ForceZeroTime) To_String() string {
	return "ForceZeroTime()"
}

func (n Node_TriggerQuestManualCompleteBehaviors) To_String() string {
	return "TriggerQuestManualCompleteBehaviors()"
}

func (n Node_DrawCone) To_String() string {
	return "DrawCone()"
}

func (n Node_Warp) To_String() string {
	return "Warp()"
}

func (n Node_ForceLookAt) To_String() string {
	return "ForceLookAt()"
}

func (n Node_SetElement) To_String() string {
	return "SetElement()"
}

func (n Node_EnableStopToAvoid) To_String() string {
	return "EnableStopToAvoid()"
}

func (n Node_EmitArrowWithEmitter) To_String() string {
	return "EmitArrowWithEmitter()"
}

func (n Node_IncrementBanterFact) To_String() string {
	return "IncrementBanterFact()"
}

func (n Node_ExecuteVFS) To_String() string {
	return "ExecuteVFS()"
}

func (n Node_DespawnDynamicCharacter) To_String() string {
	return "DespawnDynamicCharacter()"
}

func (n Node_EnableStoryTime) To_String() string {
	return "EnableStoryTime()"
}

func (n Node_Crank_SetComplete) To_String() string {
	return "Crank_SetComplete()"
}

func (n Node_EvaluateLoadZones) To_String() string {
	return "EvaluateLoadZones()"
}

func (n Node_TriggerMoveEvent) To_String() string {
	return "TriggerMoveEvent()"
}

func (n Node_DrawTextInWorld) To_String() string {
	return "DrawTextInWorld()"
}

func (n Node_ModifyFocalZoneEnabled) To_String() string {
	return "ModifyFocalZoneEnabled()"
}

func (n Node_RemoveExcludeTraverseLinkFilterToCreature) To_String() string {
	return "RemoveExcludeTraverseLinkFilterToCreature()"
}

func (n Node_Pop) To_String() string {
	return "Pop()"
}

func (n Node_LockPlayerOnWolfsled) To_String() string {
	return "LockPlayerOnWolfsled()"
}

func (n Node_DrawText2D) To_String() string {
	return "DrawText2D()"
}

func (n Node_SetPreventCheckpoint) To_String() string {
	return "SetPreventCheckpoint()"
}

func (n Node_Crank_FreeCrank) To_String() string {
	return "Crank_FreeCrank()"
}

func (n Node_SetInteractZonesEnabled) To_String() string {
	return "SetInteractZonesEnabled()"
}

func (n Node_RemoveCombatCollision) To_String() string {
	return "RemoveCombatCollision()"
}

func (n Node_SetContextActionRangeData) To_String() string {
	return "SetContextActionRangeData()"
}

func (n Node_RegisterForLookAtEvent) To_String() string {
	return "RegisterForLookAtEvent()"
}

func (n Node_SetWolfSledInteractOptions) To_String() string {
	return "SetWolfSledInteractOptions()"
}

func (n Node_SetInteractZoneTemplate) To_String() string {
	return "SetInteractZoneTemplate()"
}

func (n Node_PlaySound) To_String() string {
	return fmt.Sprintf("PlaySound(%s,%s,%s,%s,%s,%s,%s,%s,%s)", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), n.params[3].To_String(), n.params[4].To_String(), n.params[5].To_String(), n.params[6].To_String(), n.params[7].To_String(), n.params[8].To_String())
}

func (n Node_ThawPendulum) To_String() string {
	return "ThawPendulum()"
}

func (n Node_DestroyAllCombatConcussions) To_String() string {
	return "DestroyAllCombatConcussions()"
}

func (n Node_SetSoundEmitterSplineCap) To_String() string {
	return "SetSoundEmitterSplineCap()"
}

func (n Node_ShowCompassMarker) To_String() string {
	return "ShowCompassMarker()"
}

func (n Node_SetPendulumRestrictionPlane) To_String() string {
	return "SetPendulumRestrictionPlane()"
}

func (n Node_CreateVFSFloat) To_String() string {
	return "CreateVFSFloat()"
}

func (n Node_ModifyFocalZoneObstacle) To_String() string {
	return "ModifyFocalZoneObstacle()"
}

func (n Node_SetAdditionalGravityOnPendulum) To_String() string {
	return "SetAdditionalGravityOnPendulum()"
}

func (n Node_PlayUISound) To_String() string {
	return "PlayUISound()"
}

func (n Node_SetGIBlend) To_String() string {
	return "SetGIBlend()"
}

func (n Node_SetSplineEnabled) To_String() string {
	return "SetSplineEnabled()"
}

func (n Node_ClearWeaponCinematicMode) To_String() string {
	return "ClearWeaponCinematicMode()"
}

func (n Node_SluiceSetup) To_String() string {
	return "SluiceSetup()"
}

func (n Node_SetSingleContextActionEnabled) To_String() string {
	return "SetSingleContextActionEnabled()"
}

func (n Node_SetPlayerNavCurveAutorun) To_String() string {
	return "SetPlayerNavCurveAutorun()"
}

func (n Node_SetRotation) To_String() string {
	return "SetRotation()"
}

func (n Node_RegisterOnEncounterStart) To_String() string {
	return "RegisterOnEncounterStart()"
}

func (n Node_CreatePendulumWithAnchorJoint) To_String() string {
	return "CreatePendulumWithAnchorJoint()"
}

func (n Node_SetWolfSledSpeedControlOverride) To_String() string {
	return "SetWolfSledSpeedControlOverride()"
}

func (n Node_AddIncludeTraverseLinkFilterToCreature) To_String() string {
	return "AddIncludeTraverseLinkFilterToCreature()"
}

func (n Node_AddPropToSyncWithCAOwner) To_String() string {
	return "AddPropToSyncWithCAOwner()"
}

func (n Node_RegisterForControlPressEvent) To_String() string {
	return "RegisterForControlPressEvent()"
}

func (n Node_SetSplineSpeedType) To_String() string {
	return "SetSplineSpeedType()"
}

func (n Node_PlayAnimation) To_String() string {
	return "PlayAnimation()"
}

func (n Node_SetCreatureExitedWaitGate) To_String() string {
	return "SetCreatureExitedWaitGate()"
}

func (n Node_RemoveEmbedWeaponFromGameObject) To_String() string {
	return "RemoveEmbedWeaponFromGameObject()"
}

func (n Node_ChangeBehaviorTree) To_String() string {
	return "ChangeBehaviorTree()"
}

func (n Node_StopMusic) To_String() string {
	return "StopMusic()"
}

func (n Node_SetWwiseState) To_String() string {
	return "SetWwiseState()"
}

func (n Node_ActivateBoatNavCurveFunctionalityAutoMovement) To_String() string {
	return "ActivateBoatNavCurveFunctionalityAutoMovement()"
}

func (n Node_SendSystemicBanterResponse) To_String() string {
	return "SendSystemicBanterResponse()"
}

func (n Node_PlayRestartableSoundLoop) To_String() string {
	return "PlayRestartableSoundLoop()"
}

func (n Node_SetTraversePathCustomEntry) To_String() string {
	return "SetTraversePathCustomEntry()"
}

func (n Node_RecenterToFramePosition) To_String() string {
	return "RecenterToFramePosition()"
}

func (n Node_SetZiplineSpeedOverride) To_String() string {
	return "SetZiplineSpeedOverride()"
}

func (n Node_DetermineSplineSpeedAndInputLocks) To_String() string {
	return "DetermineSplineSpeedAndInputLocks()"
}

func (n Node_SetVFSInt) To_String() string {
	return "SetVFSInt()"
}

func (n Node_CopyUpgradeEquipState) To_String() string {
	return "CopyUpgradeEquipState()"
}

func (n Node_SetMaxSpeedOverride) To_String() string {
	return "SetMaxSpeedOverride()"
}

func (n Node_AddFocalZone) To_String() string {
	return "AddFocalZone()"
}

func (n Node_SetRealmLocked) To_String() string {
	return "SetRealmLocked()"
}

func (n Node_AddPositionalFlag) To_String() string {
	return "AddPositionalFlag()"
}

func (n Node_DisableContextAction) To_String() string {
	return "DisableContextAction()"
}

func (n Node_ProfileEnd) To_String() string {
	return "ProfileEnd()"
}

func (n Node_AddToGlobalVariable) To_String() string {
	return "AddToGlobalVariable()"
}

func (n Node_SetCameraTargetEnabled) To_String() string {
	return "SetCameraTargetEnabled()"
}

func (n Node_SendLuaHook) To_String() string {
	return "SendLuaHook()"
}

func (n Node_GetArrayElement) To_String() string {
	return "GetArrayElement()"
}

func (n Node_RemoveCreatureFromCustomSplineGroup) To_String() string {
	return "RemoveCreatureFromCustomSplineGroup()"
}

func (n Node_TransferFullLootResult) To_String() string {
	return "TransferFullLootResult()"
}

func (n Node_OverrideStickDamping) To_String() string {
	return "OverrideStickDamping()"
}

func (n Node_SetHighlightEnabledOverride) To_String() string {
	return "SetHighlightEnabledOverride()"
}

func (n Node_SetMaxSpeedOverrideWithType) To_String() string {
	return "SetMaxSpeedOverrideWithType()"
}

func (n Node_RegisterOnEncounterReset) To_String() string {
	return "RegisterOnEncounterReset()"
}

func (n Node_RegisterForDeathEvent) To_String() string {
	return "RegisterForDeathEvent()"
}

func (n Node_RestoreDefaultConversationSettings) To_String() string {
	return "RestoreDefaultConversationSettings()"
}

func (n Node_DespawnEncounter) To_String() string {
	return "DespawnEncounter()"
}

func (n Node_SetIsDeathAllowed) To_String() string {
	return "SetIsDeathAllowed()"
}

func (n Node_ClearMaxSpeedOverrideWithType) To_String() string {
	return "ClearMaxSpeedOverrideWithType()"
}

func (n Node_SkipWave) To_String() string {
	return "SkipWave()"
}

func (n Node_ClearAllMarkUsed) To_String() string {
	return "ClearAllMarkUsed()"
}

func (n Node_EnableBoatForceDirectionControlMode) To_String() string {
	return "EnableBoatForceDirectionControlMode()"
}

func (n Node_EnableAllInteractTags) To_String() string {
	return "EnableAllInteractTags()"
}

func (n Node_RegisterForHealthChangeEvent) To_String() string {
	return "RegisterForHealthChangeEvent()"
}

func (n Node_AcquirePickup) To_String() string {
	out := ""
	for i := 0; i < len(n.params2); i++ {
		if i > 0 {
			out += ","
		}
		out += n.params2[i].To_String()
	}
	return fmt.Sprintf("AcquirePickup(%s,%s,%s,%s,[%s])", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), n.params[3].To_String(), out)
}

func (n Node_StoreCheckpointWithPositionOverride) To_String() string {
	return "StoreCheckpointWithPositionOverride()"
}

func (n Node_HideCompassMarker) To_String() string {
	return "HideCompassMarker()"
}

func (n Node_CheckTriggerWithI) To_String() string {
	return "CheckTriggerWithI()"
}

func (n Node_BucketEquipmentByTrait) To_String() string {
	return "BucketEquipmentByTrait()"
}

func (n Node_CloseSoundPortal) To_String() string {
	return "CloseSoundPortal()"
}

func (n Node_HideWads) To_String() string {
	return "HideWads()"
}

func (n Node_ApplyEquipmentJointsToObject) To_String() string {
	return "ApplyEquipmentJointsToObject()"
}

func (n Node_HideHealthBar) To_String() string {
	return "HideHealthBar()"
}

func (n Node_Remove) To_String() string {
	return "Remove()"
}

func (n Node_SetTraverseLinkEnabled) To_String() string {
	return "SetTraverseLinkEnabled()"
}

func (n Node_RemoveFlag) To_String() string {
	return "RemoveFlag()"
}

func (n Node_ModifyFocalZone) To_String() string {
	return "ModifyFocalZone()"
}

func (n Node_FunctionCall) To_String() string {
	return "FunctionCall()"
}

func (n Node_SuspendCreatureCulling) To_String() string {
	return "SuspendCreatureCulling()"
}

func (n Node_SetContextActionCustomData) To_String() string {
	return "SetContextActionCustomData()"
}

func (n Node_Spawn) To_String() string {
	return "Spawn()"
}

func (n Node_SetInteractZoneMainPrompts) To_String() string {
	return "SetInteractZoneMainPrompts()"
}

func (n Node_FilterIterator) To_String() string {
	return "FilterIterator()"
}

func (n Node_AddLookAtPriorityOverride) To_String() string {
	return "AddLookAtPriorityOverride()"
}

func (n Node_SetInteractZoneSubPrompt) To_String() string {
	return "SetInteractZoneSubPrompt()"
}

func (n Node_DeactivatePickup) To_String() string {
	return "DeactivatePickup()"
}

func (n Node_TriggerPlayFX) To_String() string {
	return "TriggerPlayFX()"
}

func (n Node_SetTeam) To_String() string {
	return "SetTeam()"
}

func (n Node_ClonePose) To_String() string {
	return "ClonePose()"
}

func (n Node_ClearSoundEmitterSplineCap) To_String() string {
	return "ClearSoundEmitterSplineCap()"
}

func (n Node_RemoveCombatIndicator) To_String() string {
	return "RemoveCombatIndicator()"
}

func (n Node_ConfigureLogicLoadGroup) To_String() string {
	return "ConfigureLogicLoadGroup()"
}

func (n Node_RegisterForEnemySpawn) To_String() string {
	return "RegisterForEnemySpawn()"
}

func (n Node_AssignCreatureToEncounter) To_String() string {
	return "AssignCreatureToEncounter()"
}

func (n Node_SetSkillTreeToken) To_String() string {
	return "SetSkillTreeToken()"
}

func (n Node_FinishHapticVibration) To_String() string {
	return "FinishHapticVibration()"
}

func (n Node_SetParticlesEnabled) To_String() string {
	return "SetParticlesEnabled()"
}

func (n Node_RegisterForAnimFrameEvent) To_String() string {
	return "RegisterForAnimFrameEvent()"
}

func (n Node_SetThrowable) To_String() string {
	return "SetThrowable()"
}

func (n Node_SetBlackboardVariable) To_String() string {
	if n.next != nil {
		return fmt.Sprintf("SetBlackboardVariable(%s,%s,%s,%s,%d);%s", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), n.params[3].To_String(), n.param, n.next.To_String())
	}
	return fmt.Sprintf("SetBlackboardVariable(%s,%s,%s,%s,%d)", n.params[0].To_String(), n.params[1].To_String(), n.params[2].To_String(), n.params[3].To_String(), n.param)
}

func (n Node_If) To_String() string {
	if n.False == nil {
		return fmt.Sprintf("if (%s) { %s }", n.Condition.To_String(), n.True.To_String())
	} else if n.True == nil {
		return fmt.Sprintf("if (%s) { } else { %s }", n.Condition.To_String(), n.False.To_String())
	}
	return fmt.Sprintf("if (%s) { %s } else { %s }", n.Condition.To_String(), n.True.To_String(), n.False.To_String())
}

func (n Node_ResetPadLightColor) To_String() string {
	return "ResetPadLightColor()"
}

func (n Node_SetWindMotorsEnabled) To_String() string {
	return "SetWindMotorsEnabled()"
}

func (n Node_MeterAdjustPauseDelay) To_String() string {
	return "MeterAdjustPauseDelay()"
}

func (n Node_ZoomSnap) To_String() string {
	return "ZoomSnap()"
}

func (n Node_RelinquishPickup) To_String() string {
	return "RelinquishPickup()"
}

func (n Node_SetWolfSledAutoRunSplineSpeed) To_String() string {
	return "SetWolfSledAutoRunSplineSpeed()"
}

func (n Node_SetPaperBoatDistanceFromSplineClamp) To_String() string {
	return "SetPaperBoatDistanceFromSplineClamp()"
}

func (n Node_RemoveLookAtTarget) To_String() string {
	return "RemoveLookAtTarget()"
}

func (n Node_DrawCircle2D) To_String() string {
	return "DrawCircle2D()"
}

func (n Node_SetPendulumRestrictionAngle) To_String() string {
	return "SetPendulumRestrictionAngle()"
}

func (n Node_RegisterForFrameUpdateEvent) To_String() string {
	return "RegisterForFrameUpdateEvent()"
}

func (n Node_ClearGlobalVariable) To_String() string {
	return "ClearGlobalVariable()"
}

func (n Node_RemoveLookAtPriorityOverride) To_String() string {
	return "RemoveLookAtPriorityOverride()"
}

func (n Node_SetZiplineStartAndEnd) To_String() string {
	return "SetZiplineStartAndEnd()"
}

func (n Node_MusicFadeOutAndLogVolume) To_String() string {
	return "MusicFadeOutAndLogVolume()"
}

func (n Node_StartMainMusicTrack) To_String() string {
	return "StartMainMusicTrack()"
}

func (n Node_SetPadLightColor) To_String() string {
	return "SetPadLightColor()"
}

func (n Node_EmitArrow) To_String() string {
	return "EmitArrow()"
}

func (n Node_RegisterForEventQueueProcessedEvent) To_String() string {
	return "RegisterForEventQueueProcessedEvent()"
}

func (n Node_SetProgressionFact) To_String() string {
	return "SetProgressionFact()"
}

func (n Node_ResetTraverseLinkFilterForCreature) To_String() string {
	return "ResetTraverseLinkFilterForCreature()"
}

func (n Node_CheckPersistentCreaturesEnabled) To_String() string {
	return "CheckPersistentCreaturesEnabled()"
}

func (n Node_DisableLoadGate) To_String() string {
	return "DisableLoadGate()"
}

func (n Node_RemoveFX) To_String() string {
	return "RemoveFX()"
}

func (n Node_RemoveArrows) To_String() string {
	return "RemoveArrows()"
}

func (n Node_ResetTraverseLinkFilter) To_String() string {
	return "ResetTraverseLinkFilter()"
}

func (n Node_FreeRequestedPrestreams) To_String() string {
	return "FreeRequestedPrestreams()"
}

func (n Node_RecenterPendulum) To_String() string {
	return "RecenterPendulum()"
}

func (n Node_SetWolfSledAutoRunEnableSteering) To_String() string {
	return "SetWolfSledAutoRunEnableSteering()"
}

func (n Node_SetSingleNavCurveEnabled) To_String() string {
	return "SetSingleNavCurveEnabled()"
}

func (n Node_SetWeaponCinematicMode) To_String() string {
	return "SetWeaponCinematicMode()"
}

func (n Node_SendContextActionStim) To_String() string {
	return "SendContextActionStim()"
}

func (n Node_AddEquipmentPreview) To_String() string {
	return "AddEquipmentPreview()"
}

func (n Node_EquipToCharacterSlot) To_String() string {
	return "EquipToCharacterSlot()"
}

func (n Node_StoreSoftSave) To_String() string {
	return "StoreSoftSave()"
}

func (n Node_ForceInterruptContextAction) To_String() string {
	return "ForceInterruptContextAction()"
}

func (n Node_TriggerEncounter) To_String() string {
	return "TriggerEncounter()"
}

func (n Node_SetCreaturePersistent) To_String() string {
	return "SetCreaturePersistent()"
}

func (n Node_SetSoundGeometryEnabled) To_String() string {
	return "SetSoundGeometryEnabled()"
}

func (n Node_RegisterForButtonPressEvent) To_String() string {
	return "RegisterForButtonPressEvent()"
}

func (n Node_SetHapticGlobalParameter) To_String() string {
	return "SetHapticGlobalParameter()"
}

func (n Node_ForceCompanionStriderPath) To_String() string {
	return "ForceCompanionStriderPath()"
}

func (n Node_Push) To_String() string {
	return "Push()"
}

func (n Node_Sluice_WaterPourHit) To_String() string {
	return "Sluice_WaterPourHit()"
}

func (n Node_GetParticleMonsterTotemCheckpointState) To_String() string {
	return "GetParticleMonsterTotemCheckpointState()"
}

func (n Node_AddTraverseLinkFilter) To_String() string {
	return "AddTraverseLinkFilter()"
}

func (n Node_SaveYakWaterSpeedType) To_String() string {
	return "SaveYakWaterSpeedType()"
}

func (n Node_UpdateGameObjectSplineCheap) To_String() string {
	return "UpdateGameObjectSplineCheap()"
}

func (n Node_SetInfluenceCircleAttributes) To_String() string {
	return "SetInfluenceCircleAttributes()"
}

func (n Node_WildlifeEvent) To_String() string {
	return "WildlifeEvent()"
}

func (n Node_RegisterOnEncounterCreated) To_String() string {
	return "RegisterOnEncounterCreated()"
}

func (n Node_SpawnDynamicCharacter) To_String() string {
	return "SpawnDynamicCharacter()"
}

func (n Node_SendAwarenessStimWithPa) To_String() string {
	return "SendAwarenessStimWithPa()"
}

func (n Node_SetVariable) To_String() string {
	if n.next != nil {
		return fmt.Sprintf("SetVariable(%s,%s);%s", n.name, n.val.To_String(), n.next.To_String())
	}
	return fmt.Sprintf("SetVariable(%s,%s)", n.name, n.val.To_String())
}

func (n Node_SetTraversePathSegmentEnabled) To_String() string {
	return "SetTraversePathSegmentEnabled()"
}

func (n Node_SetCreatureEnteredWaitGate) To_String() string {
	return "SetCreatureEnteredWaitGate()"
}

func (n Node_ClearCreatureMotionParamsOverride) To_String() string {
	return "ClearCreatureMotionParamsOverride()"
}

func (n Node_RegisterForReticleTargetInvalid) To_String() string {
	return "RegisterForReticleTargetInvalid()"
}

func (n Node_SetCompanionFollowStriderPath) To_String() string {
	return "SetCompanionFollowStriderPath()"
}

func (n Node_SetCompassEnabled) To_String() string {
	return "SetCompassEnabled()"
}

func (n Node_AddChild) To_String() string {
	return "AddChild()"
}

func (n Node_ForceUseSpline) To_String() string {
	return "ForceUseSpline()"
}

func (n Node_SetModelEnabled) To_String() string {
	return "SetModelEnabled()"
}

func (n Node_SendAwareness) To_String() string {
	return "SendAwareness()"
}

func (n Node_SetVisualScriptEnabled) To_String() string {
	return "SetVisualScriptEnabled()"
}

func (n Node_UnregisterTraverseGraphEnteredEvent) To_String() string {
	return "UnregisterTraverseGraphEnteredEvent()"
}

func (n Node_CompleteWave) To_String() string {
	return "CompleteWave()"
}

func (n Node_SetWolfSledAutoRunResetSettings) To_String() string {
	return "SetWolfSledAutoRunResetSettings()"
}

func (n Node_SetVFSFloat) To_String() string {
	return "SetVFSFloat()"
}

func (n Node_SetAnimFloatChannelDriverValue) To_String() string {
	return "SetAnimFloatChannelDriverValue()"
}

func (n Node_UpdateStaticWindClientPositions) To_String() string {
	return "UpdateStaticWindClientPositions()"
}

func (n Node_RotateObjectTowardsTarget) To_String() string {
	return "RotateObjectTowardsTarget()"
}

func (n Node_SetPaperBoatSpringValues) To_String() string {
	return "SetPaperBoatSpringValues()"
}

func (n Node_CancelLockOn) To_String() string {
	return "CancelLockOn()"
}

func (n Node_StartMusic) To_String() string {
	return "StartMusic()"
}

func (n Node_RegisterForUIButtonPressEvent) To_String() string {
	return "RegisterForUIButtonPressEvent()"
}

func (n Node_StopBlenderEffect) To_String() string {
	return "StopBlenderEffect()"
}

func (n Node_SetForceCompanionTraverseLink) To_String() string {
	return "SetForceCompanionTraverseLink()"
}

func (n Node_SetDecalsEnabled) To_String() string {
	return "SetDecalsEnabled()"
}

func (n Node_ClearRage) To_String() string {
	return "ClearRage()"
}

func (n Node_GetInteractObject) To_String() string {
	return "GetInteractObject()"
}

func (n Node_SetSavedNavCurveAttributes) To_String() string {
	return "SetSavedNavCurveAttributes()"
}

func (n Node_SetSplineLeadDistances) To_String() string {
	return "SetSplineLeadDistances()"
}

func (n Node_ModifyFocalZoneObstacleEnabled) To_String() string {
	return "ModifyFocalZoneObstacleEnabled()"
}

func (n Node_SpawnFX) To_String() string {
	return "SpawnFX()"
}

func (n Node_SpherecastCollision) To_String() string {
	return "SpherecastCollision()"
}

func (n Node_SimulateSpearZipline) To_String() string {
	return "SimulateSpearZipline()"
}

func (n Node_SetAllWeaponsCinematicMode) To_String() string {
	return "SetAllWeaponsCinematicMode()"
}

func (n Node_IncrementProgressionFact) To_String() string {
	return "IncrementProgressionFact()"
}

func (n Node_RegisterOnEnemyDeath) To_String() string {
	return "RegisterOnEnemyDeath()"
}

func (n Node_UnlockCombatStatus) To_String() string {
	return "UnlockCombatStatus()"
}

func (n Node_CompleteWaveElement) To_String() string {
	return "CompleteWaveElement()"
}

func (n Node_SetTraverseLinksEnabled) To_String() string {
	return "SetTraverseLinksEnabled()"
}

func (n Node_SetCharacterConfig) To_String() string {
	return "SetCharacterConfig()"
}

func (n Node_EnableBoatForceTurnAroundControlMode) To_String() string {
	return "EnableBoatForceTurnAroundControlMode()"
}

func (n Node_FrameDelay) To_String() string {
	return "FrameDelay()"
}

func (n Node_SetInteractTagEnabled) To_String() string {
	return "SetInteractTagEnabled()"
}

func (n Node_DestroyPendulum) To_String() string {
	return "DestroyPendulum()"
}

func (n Node_SetZiplineEndDisconnected) To_String() string {
	return "SetZiplineEndDisconnected()"
}

func (n Node_DestroyGameObject) To_String() string {
	return "DestroyGameObject()"
}

func (n Node_SetVolumeEnabled) To_String() string {
	return "SetVolumeEnabled()"
}

func (n Node_ClearAllInteractZoneTags) To_String() string {
	return "ClearAllInteractZoneTags()"
}

func (n Node_AddSimpleObjectActorToSequence) To_String() string {
	return "AddSimpleObjectActorToSequence()"
}

func (n Node_PushWeaponStateForSequence) To_String() string {
	return "PushWeaponStateForSequence()"
}

func (n Node_ClearCustomSplineGroup) To_String() string {
	return "ClearCustomSplineGroup()"
}

func (n Node_AddEquipmentToWallet) To_String() string {
	return "AddEquipmentToWallet()"
}

func (n Node_CheckDecision) To_String() string {
	return "CheckDecision()"
}

func (n Node_SetContextActionLoopActionData) To_String() string {
	return "SetContextActionLoopActionData()"
}

func (n Node_SetOverrideEnterBranchesOnDock) To_String() string {
	return "SetOverrideEnterBranchesOnDock()"
}

func (n Node_SetOverrideStatusEffectIconParent) To_String() string {
	return "SetOverrideStatusEffectIconParent()"
}

func (n Node_SetWwiseSwitch) To_String() string {
	return "SetWwiseSwitch()"
}

func (n Node_Error) To_String() string {
	return "Error()"
}

func (n Node_RequestMaterialSwap) To_String() string {
	return "RequestMaterialSwap()"
}

func (n Node_SetContextActionTraversalBehavior) To_String() string {
	return "SetContextActionTraversalBehavior()"
}

func (n Node_BatchQueryEquipmentTraits) To_String() string {
	return "BatchQueryEquipmentTraits()"
}

func (n Node_ApplyContextIdle) To_String() string {
	return "ApplyContextIdle()"
}

func (n Node_SetContextActionBasicBtreeData) To_String() string {
	return "SetContextActionBasicBtreeData()"
}

func (n Node_SubtractFromVariable) To_String() string {
	return "SubtractFromVariable()"
}

func (n Node_SetCompanion) To_String() string {
	return "SetCompanion()"
}

func (n Node_MusicFadeIn) To_String() string {
	return "MusicFadeIn()"
}

func (n Node_SpawnGameObject) To_String() string {
	return "SpawnGameObject()"
}

func (n Node_SendPreStartNamedEvent) To_String() string {
	return "SendPreStartNamedEvent()"
}

func (n Node_FindChildGameObjectWithName) To_String() string {
	return "FindChildGameObjectWithName()"
}

func (n Node_UnequipFromCharacterSlot) To_String() string {
	return "UnequipFromCharacterSlot()"
}

func (n Node_LoadFastTravelWads) To_String() string {
	return "LoadFastTravelWads()"
}

func (n Node_SetComponentsEnabled) To_String() string {
	return "SetComponentsEnabled()"
}

func (n Node_AddLootResultToWallet) To_String() string {
	return "AddLootResultToWallet()"
}

func (n Node_SetAttachmentConfig) To_String() string {
	return "SetAttachmentConfig()"
}

func (n Node_SetPlayGoRequired) To_String() string {
	return "SetPlayGoRequired()"
}

func (n Node_SetMarkerClaimed) To_String() string {
	return "SetMarkerClaimed()"
}

func (n Node_StoreCheckpoint) To_String() string {
	return "StoreCheckpoint()"
}

func (n Node_SetPaperBoatPitchRollValues) To_String() string {
	return "SetPaperBoatPitchRollValues()"
}

func (n Node_RelinquishAllPickups) To_String() string {
	return "RelinquishAllPickups()"
}

func (n Node_ModifyFocalZoneFacingAngleActivationTolerance) To_String() string {
	return "ModifyFocalZoneFacingAngleActivationTolerance()"
}

func (n Node_Insert) To_String() string {
	return "Insert()"
}

func (n Node_SetCollisionEnabled) To_String() string {
	return "SetCollisionEnabled()"
}

func (n Node_IncrementPickupStage) To_String() string {
	return "IncrementPickupStage()"
}

func (n Node_RemoveInteractZoneTag) To_String() string {
	return "RemoveInteractZoneTag()"
}

func (n Node_SetUDSActivityEnd) To_String() string {
	return "SetUDSActivityEnd()"
}

func (n Node_SetLightsEnabled) To_String() string {
	return "SetLightsEnabled()"
}

func (n Node_SetTraversePathEnabled) To_String() string {
	return "SetTraversePathEnabled()"
}

func (n Node_RegisterForControlSwipeEvent) To_String() string {
	return "RegisterForControlSwipeEvent()"
}

func (n Node_AbortCreatureInteract) To_String() string {
	return "AbortCreatureInteract()"
}

func (n Node_ClearBlackboardVariable) To_String() string {
	if n.next != nil {
		return fmt.Sprintf("ClearBlackboardVariable(%s,%s,%d);%s", n.params[0].To_String(), n.params[1].To_String(), n.param, n.next.To_String())
	}
	return fmt.Sprintf("ClearBlackboardVariable(%s,%s,%d)", n.params[0].To_String(), n.params[1].To_String(), n.param)
}

func (n Node_SetParticleEmitterEnabled) To_String() string {
	return "SetParticleEmitterEnabled()"
}

func (n Node_AddToVariable) To_String() string {
	return "AddToVariable()"
}

func (n Node_RecenterToAngle) To_String() string {
	return "RecenterToAngle()"
}

func (n Node_RelinquishPickupBySlot) To_String() string {
	return "RelinquishPickupBySlot()"
}

func (n Node_AddCreatureActorToSequence) To_String() string {
	return "AddCreatureActorToSequence()"
}

func (n Node_TriggerWave) To_String() string {
	return "TriggerWave()"
}

func (n Node_SetUDSActivityUnavailable) To_String() string {
	return "SetUDSActivityUnavailable()"
}

func (n Node_RegisterForWeaponThrown) To_String() string {
	return "RegisterForWeaponThrown()"
}

func (n Node_DestroyArrows) To_String() string {
	return "DestroyArrows()"
}

func (n Node_DecrementPickupStage) To_String() string {
	return "DecrementPickupStage()"
}

func (n Node_SetContextActionsEnabled) To_String() string {
	return "SetContextActionsEnabled()"
}

func (n Node_ScentTracker_SetUpdateFrequency) To_String() string {
	return "ScentTracker_SetUpdateFrequency()"
}

func (n Node_SetContextActionTraversalExitSegment) To_String() string {
	return "SetContextActionTraversalExitSegment()"
}

func (n Node_DropChainedObject) To_String() string {
	return "DropChainedObject()"
}

func (n Node_DrawNavCurve) To_String() string {
	return "DrawNavCurve()"
}

func (n Node_SetTimeScale) To_String() string {
	return "SetTimeScale()"
}

func (n Node_SetContextActionBranchData) To_String() string {
	return "SetContextActionBranchData()"
}

func (n Node_DrawLine) To_String() string {
	return "DrawLine()"
}

func (n Node_SetInteractZonePromptEnabled) To_String() string {
	return "SetInteractZonePromptEnabled()"
}

func (n Node_RegisterMinibossCheckpoint) To_String() string {
	return "RegisterMinibossCheckpoint()"
}

func (n Node_RegisterObjectAsProp) To_String() string {
	return "RegisterObjectAsProp()"
}

func (n Node_AlertWaveElement) To_String() string {
	return "AlertWaveElement()"
}

func (n Node_GetCreature) To_String() string {
	return fmt.Sprintf("GetCreature(%s)", n.param.To_String())
}

func (n Node_LockOn) To_String() string {
	return "LockOn()"
}

func (n Node_SetShowPlaceholderBoatOnDockDisable) To_String() string {
	return "SetShowPlaceholderBoatOnDockDisable()"
}

func (n Node_SetCompanionWaitOnTraversePath) To_String() string {
	return "SetCompanionWaitOnTraversePath()"
}

func (n Node_CineOnly_SwapCreatureToExistingObject) To_String() string {
	return "CineOnly_SwapCreatureToExistingObject()"
}

func (n Node_LoadCheck) To_String() string {
	return "LoadCheck()"
}

func (n Node_ResetCompass) To_String() string {
	return "ResetCompass()"
}

func (n Node_Event) To_String() string {
	return "Event()"
}

func (n Node_RegisteredEvent) To_String() string {
	return "RegisteredEvent()"
}

func (n Node_Data) To_String() string {
	return "Data()"
}

func (n Node_Triggerable) To_String() string {
	return "Triggerable()"
}

func (n Node_ConstantDisplay) To_String() string {
	return "ConstantDisplay()"
}

func (n Node_Constant) To_String() string {
	return "Constant()"
}

func (n Node_EditorAssetReference) To_String() string {
	return "EditorAssetReference()"
}

func (n Node_Struct) To_String() string {
	return "Struct()"
}

func (n Node_EnumReference) To_String() string {
	return "EnumReference()"
}

func (n Node_WadId) To_String() string {
	return "WadId()"
}

func (n Node_RawUniqueId) To_String() string {
	return "RawUniqueId()"
}

func (n Node_RawStringHash) To_String() string {
	return "RawStringHash()"
}

func (n Node_GenericGuidPathReference) To_String() string {
	return "GenericGuidPathReference()"
}

func (n Node_GameObjectReference) To_String() string {
	return "GameObjectReference()"
}

func (n Node_InstancedObjectReference) To_String() string {
	return "InstancedObjectReference()"
}

func (n Node_InteractZoneReference) To_String() string {
	return "InteractZoneReference()"
}

func (n Node_ArrowId) To_String() string {
	return "ArrowId()"
}

func (n Node_ArrowEmitterId) To_String() string {
	return "ArrowEmitterId()"
}

func (n Node_BeamAttackId) To_String() string {
	return "BeamAttackId()"
}

func (n Node_CameraId) To_String() string {
	return "CameraId()"
}

func (n Node_CameraRecenterId) To_String() string {
	return "CameraRecenterId()"
}

func (n Node_GatewayMarkerId) To_String() string {
	return "GatewayMarkerId()"
}

func (n Node_ConcussionId) To_String() string {
	return "ConcussionId()"
}

func (n Node_DecisionId) To_String() string {
	return "DecisionId()"
}

func (n Node_EquipmentId) To_String() string {
	return "EquipmentId()"
}

func (n Node_IconId) To_String() string {
	return "IconId()"
}

func (n Node_MeterId) To_String() string {
	return "MeterId()"
}

func (n Node_PickupId) To_String() string {
	return "PickupId()"
}

func (n Node_PickupSlotId) To_String() string {
	return "PickupSlotId()"
}

func (n Node_QuestId) To_String() string {
	return "QuestId()"
}

func (n Node_RumbleId) To_String() string {
	return "RumbleId()"
}

func (n Node_ScreenShakeId) To_String() string {
	return "ScreenShakeId()"
}

func (n Node_FullScreenEffectId) To_String() string {
	return "FullScreenEffectId()"
}

func (n Node_BranchId) To_String() string {
	return "BranchId()"
}

func (n Node_MapRealmId) To_String() string {
	return "MapRealmId()"
}

func (n Node_MapRegionId) To_String() string {
	return "MapRegionId()"
}

func (n Node_MapMarkerId) To_String() string {
	return "MapMarkerId()"
}

func (n Node_MapAreaId) To_String() string {
	return "MapAreaId()"
}

func (n Node_PropId) To_String() string {
	return "PropId()"
}

func (n Node_TweakConstantId) To_String() string {
	return "TweakConstantId()"
}

func (n Node_PlayFXId) To_String() string {
	return "PlayFXId()"
}

func (n Node_InteractZoneTemplateId) To_String() string {
	return "InteractZoneTemplateId()"
}

func (n Node_AnimationId) To_String() string {
	return "AnimationId()"
}

func (n Node_AnimJointReference) To_String() string {
	return "AnimJointReference()"
}

func (n Node_TraversePathReference) To_String() string {
	return "TraversePathReference()"
}

func (n Node_TraverseLinkReference) To_String() string {
	return "TraverseLinkReference()"
}

func (n Node_ContextActionReference) To_String() string {
	return "ContextActionReference()"
}

func (n Node_NavCurveReference) To_String() string {
	return "NavCurveReference()"
}

func (n Node_SoundEmitterReference) To_String() string {
	return "SoundEmitterReference()"
}

func (n Node_SoundPortalReference) To_String() string {
	return "SoundPortalReference()"
}

func (n Node_SoundProximityTriggerReference) To_String() string {
	return "SoundProximityTriggerReference()"
}

func (n Node_BehaviorTreeRootReference) To_String() string {
	return "BehaviorTreeRootReference()"
}

func (n Node_BehaviorTreeSubtreeReference) To_String() string {
	return "BehaviorTreeSubtreeReference()"
}

func (n Node_WwiseEventId) To_String() string {
	return "WwiseEventId()"
}

func (n Node_WwiseSwitchGroupId) To_String() string {
	return "WwiseSwitchGroupId()"
}

func (n Node_WwiseSwitchStateId) To_String() string {
	return "WwiseSwitchStateId()"
}

func (n Node_WwiseRTPCId) To_String() string {
	return "WwiseRTPCId()"
}

func (n Node_LamsId) To_String() string {
	return "LamsId()"
}

func (n Node_BanterId) To_String() string {
	return "BanterId()"
}

func (n Node_RecipeId) To_String() string {
	return "RecipeId()"
}

func (n Node_ResourceId) To_String() string {
	return "ResourceId()"
}

func (n Node_WalletId) To_String() string {
	return "WalletId()"
}

func (n Node_WeaponId) To_String() string {
	return "WeaponId()"
}

func (n Node_WildlifeId) To_String() string {
	return "WildlifeId()"
}

func (n Node_HapticId) To_String() string {
	return "HapticId()"
}

func (n Node_LootConditionId) To_String() string {
	return "LootConditionId()"
}

func (n Node_LootConditionSetId) To_String() string {
	return "LootConditionSetId()"
}

func (n Node_LootDistributorId) To_String() string {
	return "LootDistributorId()"
}

func (n Node_RelativeReference) To_String() string {
	return "RelativeReference()"
}

func (n Node_RelativeGameObjectReference) To_String() string {
	return "RelativeGameObjectReference()"
}

func (n Node_RelativeInteractZoneReference) To_String() string {
	return "RelativeInteractZoneReference()"
}

func (n Node_RelativeTraversePathReference) To_String() string {
	return "RelativeTraversePathReference()"
}

func (n Node_RelativeTraverseLinkReference) To_String() string {
	return "RelativeTraverseLinkReference()"
}

func (n Node_RelativeContextActionReference) To_String() string {
	return "RelativeContextActionReference()"
}

func (n Node_RelativeNavCurveReference) To_String() string {
	return "RelativeNavCurveReference()"
}

func (n Node_RelativeSoundEmitterReference) To_String() string {
	return "RelativeSoundEmitterReference()"
}

func (n Node_RelativeSoundPortalReference) To_String() string {
	return "RelativeSoundPortalReference()"
}

func (n Node_RelativeSoundProximityTriggerReference) To_String() string {
	return "RelativeSoundProximityTriggerReference()"
}

func (n Node_RelativeAnimJoint) To_String() string {
	return "RelativeAnimJoint()"
}

func (n Node_AsWadId) To_String() string {
	return "AsWadId()"
}

func (n Node_TweaksIdAsStringHash) To_String() string {
	return "TweaksIdAsStringHash()"
}

func (n Node_ConcatenateAndHashRawStrings) To_String() string {
	return "ConcatenateAndHashRawStrings()"
}

func (n Node_OnFastTravelWadsLoaded) To_String() string {
	return "OnFastTravelWadsLoaded()"
}

func (n Node_OnBeamCreated) To_String() string {
	return "OnBeamCreated()"
}

func (n Node_OnTimerComplete) To_String() string {
	return fmt.Sprintf("function(){%s}", n.node.To_String())
}

func (n Node_OnFrameDelay) To_String() string {
	return "OnFrameDelay()"
}

func (n Node_OnSpawnWildlifeComplete) To_String() string {
	return "OnSpawnWildlifeComplete()"
}

func (n Node_OnSpawnComplete) To_String() string {
	return "OnSpawnComplete()"
}

func (n Node_OnSpawnedObjectFrameUpdate) To_String() string {
	return "OnSpawnedObjectFrameUpdate()"
}

func (n Node_OnSpawnedObjectFrameUpdateExpired) To_String() string {
	return "OnSpawnedObjectFrameUpdateExpired()"
}

func (n Node_OnSpawnedObjectDestroyed) To_String() string {
	return "OnSpawnedObjectDestroyed()"
}

func (n Node_OnSpawnedObjectRecycled) To_String() string {
	return "OnSpawnedObjectRecycled()"
}

func (n Node_OnSpawnFXComplete) To_String() string {
	return "OnSpawnFXComplete()"
}

func (n Node_OnMaterialAnimationDone) To_String() string {
	return "OnMaterialAnimationDone()"
}

func (n Node_OnAnimationDone) To_String() string {
	return "OnAnimationDone()"
}

func (n Node_RegisterCombatIndicator) To_String() string {
	return "RegisterCombatIndicator()"
}

func (n Node_OnArrowEmitted) To_String() string {
	return "OnArrowEmitted()"
}

func (n Node_OnLoadCheckFinished) To_String() string {
	return "OnLoadCheckFinished()"
}

func (n Node_OnBlackboardTimerExpired) To_String() string {
	return "OnBlackboardTimerExpired()"
}

func (n Node_OnWarpComplete) To_String() string {
	return "OnWarpComplete()"
}

func (n Node_ModifyFocalZoneCameraAngleActivationTolerance) To_String() string {
	return "ModifyFocalZoneCameraAngleActivationTolerance()"
}

func (n Node_OnFocalZonePreActivationEnter) To_String() string {
	return "OnFocalZonePreActivationEnter()"
}

func (n Node_OnFocalZonePreActivationExit) To_String() string {
	return "OnFocalZonePreActivationExit()"
}

func (n Node_OnBanterStart) To_String() string {
	return "OnBanterStart()"
}

func (n Node_OnBanterDone) To_String() string {
	return "OnBanterDone()"
}

func (n Node_OnNextBanterDone) To_String() string {
	return "OnNextBanterDone()"
}

func (n Node_OnBanterFailed) To_String() string {
	return "OnBanterFailed()"
}

func (n Node_RemoveEquipmentFromWallet) To_String() string {
	return "RemoveEquipmentFromWallet()"
}

func (n Node_OnUIMessageClosed) To_String() string {
	return "OnUIMessageClosed()"
}

func (n Node_OnCallAndResponseTriggered) To_String() string {
	return "OnCallAndResponseTriggered()"
}

func (n Node_OnCallAndResponseFinished) To_String() string {
	return "OnCallAndResponseFinished()"
}

func (n Node_OnRegisteredFrameUpdate) To_String() string {
	return "OnRegisteredFrameUpdate()"
}

func (n Node_OnAnimFrame) To_String() string {
	return "OnAnimFrame()"
}

func (n Node_OnMovePlay) To_String() string {
	return "OnMovePlay()"
}

func (n Node_OnHealthChange) To_String() string {
	return "OnHealthChange()"
}

func (n Node_OnDeath) To_String() string {
	return "OnDeath()"
}

func (n Node_OnLookAt) To_String() string {
	return "OnLookAt()"
}

func (n Node_OnButtonPress) To_String() string {
	return "OnButtonPress()"
}

func (n Node_OnUIButtonPress) To_String() string {
	return "OnUIButtonPress()"
}

func (n Node_OnControlPress) To_String() string {
	return "OnControlPress()"
}

func (n Node_OnControlMash) To_String() string {
	return "OnControlMash()"
}

func (n Node_OnDistanceCheck) To_String() string {
	return "OnDistanceCheck()"
}

func (n Node_OnEventQueueProcessed) To_String() string {
	return "OnEventQueueProcessed()"
}

func (n Node_OnRegisteredBreakableBroken) To_String() string {
	return "OnRegisteredBreakableBroken()"
}

func (n Node_OnRegisteredMeterChanged) To_String() string {
	return "OnRegisteredMeterChanged()"
}

func (n Node_OnControlSwipe) To_String() string {
	return "OnControlSwipe()"
}

func (n Node_OnReticleTargetInvalid) To_String() string {
	return "OnReticleTargetInvalid()"
}

func (n Node_OnVariableChanged) To_String() string {
	return "OnVariableChanged()"
}

func (n Node_LoadGate) To_String() string {
	return "LoadGate()"
}

func (n Node_OnStartGameFromThisWad) To_String() string {
	return "OnStartGameFromThisWad()"
}

func (n Node_ClearAllWeather) To_String() string {
	return "ClearAllWeather()"
}

func (n Node_CreatureSpawnOptionsStruct) To_String() string {
	return "CreatureSpawnOptionsStruct()"
}

func (n Node_OnEncounterStart) To_String() string {
	return "OnEncounterStart()"
}

func (n Node_OnEncounterCreated) To_String() string {
	return "OnEncounterCreated()"
}

func (n Node_OnEncounterFinished) To_String() string {
	return "OnEncounterFinished()"
}

func (n Node_OnEncounterReset) To_String() string {
	return "OnEncounterReset()"
}

func (n Node_OnEnemySpawn) To_String() string {
	return "OnEnemySpawn()"
}

func (n Node_OnDynamicCharacterSpawn) To_String() string {
	return "OnDynamicCharacterSpawn()"
}

func (n Node_OnEnemyDeath) To_String() string {
	return "OnEnemyDeath()"
}

func (n Node_Crank_Callback) To_String() string {
	return "Crank_Callback()"
}

func (n Node_Sluice_Callback) To_String() string {
	return "Sluice_Callback()"
}

func (n Node_Sluice_SoundCallback) To_String() string {
	return "Sluice_SoundCallback()"
}

func (n Node_OnInteractStart) To_String() string {
	return "OnInteractStart()"
}

func (n Node_OnInteractFinished) To_String() string {
	return "OnInteractFinished()"
}

func (n Node_OnInteractAborted) To_String() string {
	return "OnInteractAborted()"
}

func (n Node_OnInteractDone) To_String() string {
	return "OnInteractDone()"
}

func (n Node_OnInteractAttemptedWhileOccupied) To_String() string {
	return "OnInteractAttemptedWhileOccupied()"
}

func (n Node_OnTraverseGraphEntryEvent) To_String() string {
	return "OnTraverseGraphEntryEvent()"
}

func (n Node_OnTraverseGraphEnteredEvent) To_String() string {
	return "OnTraverseGraphEnteredEvent()"
}

func (n Node_OnCreatedVFSEntryChanged) To_String() string {
	return "OnCreatedVFSEntryChanged()"
}

func (n Node_OnCreatedVFSEntryExecuted) To_String() string {
	return "OnCreatedVFSEntryExecuted()"
}

func (n Node_OnConcussionHitBase) To_String() string {
	return "OnConcussionHitBase()"
}

func (n Node_OnArrowImpactBase) To_String() string {
	return "OnArrowImpactBase()"
}

func (n Node_OnCombatCollisionBase) To_String() string {
	return "OnCombatCollisionBase()"
}

func (n Node_OnWeaponThrown) To_String() string {
	return "OnWeaponThrown()"
}

func (n Node_OnThrownWeaponBlockedBase) To_String() string {
	return "OnThrownWeaponBlockedBase()"
}

func (n Node_NamedEventBase) To_String() string {
	return "NamedEventBase()"
}

func (n Node_OnImmediateEvent) To_String() string {
	return "OnImmediateEvent()"
}

func (n Node_SyncedSequenceActorInfo) To_String() string {
	return "SyncedSequenceActorInfo()"
}

func (n Node_SyncedSequenceSimpleObjectActorInfo) To_String() string {
	return "SyncedSequenceSimpleObjectActorInfo()"
}

func (n Node_SyncedSequenceCreatureActorInfo) To_String() string {
	return "SyncedSequenceCreatureActorInfo()"
}

func (n Node_SyncedSequenceDynamicActorInfo) To_String() string {
	return "SyncedSequenceDynamicActorInfo()"
}

func (n Node_SyncedSequenceSimpleObjectActorInfoSet) To_String() string {
	return "SyncedSequenceSimpleObjectActorInfoSet()"
}

func (n Node_SyncedSequenceCreatureIdentifier) To_String() string {
	return "SyncedSequenceCreatureIdentifier()"
}

func (n Node_SyncedSequenceCreatureActorInfoSet) To_String() string {
	return "SyncedSequenceCreatureActorInfoSet()"
}

func (n Node_MotionWarpParameters) To_String() string {
	return "MotionWarpParameters()"
}

func (n Node_SyncedSequence) To_String() string {
	return "SyncedSequence()"
}

func (n Node_RegisteredSyncedSequence) To_String() string {
	return "RegisteredSyncedSequence()"
}

func (n Node_DynamicSyncedSequenceEventContainer) To_String() string {
	return "DynamicSyncedSequenceEventContainer()"
}

func (n Node_SyncedSequenceEvent) To_String() string {
	return "SyncedSequenceEvent()"
}

func (n Node_SyncedSequenceOnSkip) To_String() string {
	return "SyncedSequenceOnSkip()"
}

func (n Node_SyncedSequenceOnSkipExit) To_String() string {
	return "SyncedSequenceOnSkipExit()"
}

func (n Node_SyncedSequenceOnSequenceExit) To_String() string {
	return "SyncedSequenceOnSequenceExit()"
}

func (n Node_SyncedSequenceOnTimeReached) To_String() string {
	return "SyncedSequenceOnTimeReached()"
}

func (n Node_SyncedSequenceOnMoveTimeReached) To_String() string {
	return "SyncedSequenceOnMoveTimeReached()"
}

func (n Node_FocalZoneStrafe) To_String() string {
	return "FocalZoneStrafe()"
}

func (n Node_SyncedSequenceTime) To_String() string {
	return "SyncedSequenceTime()"
}

func ReadNodeHead(file *os.File, offset uint32) (Node, error) {
	_, err := file.Seek(int64(offset), 0)
	if err != nil {
		return nil, err
	}
	var nodeId uint64
	err = binary.Read(file, binary.BigEndian, &nodeId)
	if err != nil {
		return nil, err
	}
	switch nodeId {
	case 0x50088220865F7D78:
		fallthrough
	case 0x28467E577E4BF816:
		fallthrough
	case 0x2D7869A0C6650E33:
		fallthrough
	case 0x85B43DEC04D96B7B:
		fallthrough
	case 0x703FD9308848087C:
		fallthrough
	case 0x30298D107985F295:
		fallthrough
	case 0x397D1DE8A6DE37B0:
		fallthrough
	case 0xAD686B21CD01B5B3:
		fallthrough
	case 0xC69765B16ED92CE5:
		fallthrough
	case 0x6FB55EE2FA4C6FEC:
		fallthrough
	case 0xCE0218F8A6C2B9F9:
		fallthrough
	case 0x8EB080F430E6BC3B:
		fallthrough
	case 0x8AAE018DF0E9C5B0:
		fallthrough
	case 0x02448C70A1903303:
		n := &Node_Header{}
		n.nodeMap = make(map[uint16]uint32, 0)
		n.node, err = ReadNode(file, []uint16{}, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x04ae331e307a3420:
		return &Node_OnFastTravelWadsLoaded{}, nil
	case 0x926caeba606562e8:
		return &Node_OnBeamCreated{}, nil
	case 0x514d9afff1681639:
		return &Node_OnTimerComplete{}, nil
	case 0xcb8641b853fbcd9c:
		return &Node_OnFrameDelay{}, nil
	case 0x055359a614867cc6:
		return &Node_OnSpawnWildlifeComplete{}, nil
	case 0xa04a86f51c8e34c6:
		return &Node_OnSpawnComplete{}, nil
	case 0x90217b02c282d226:
		return &Node_OnSpawnedObjectFrameUpdate{}, nil
	case 0x7f2545f59c5ca5b8:
		return &Node_OnSpawnedObjectFrameUpdateExpired{}, nil
	case 0xa0261df565152550:
		return &Node_OnSpawnedObjectDestroyed{}, nil
	case 0x407762df190d8e8a:
		return &Node_OnSpawnedObjectRecycled{}, nil
	case 0xdbc491e75c2b829c:
		return &Node_OnSpawnFXComplete{}, nil
	case 0x5cd26a7454f3b9bb:
		return &Node_OnMaterialAnimationDone{}, nil
	case 0x231fe60f7af11389:
		return &Node_OnAnimationDone{}, nil
	case 0x174bf2b7c7b7bc6a:
		return &Node_OnArrowEmitted{}, nil
	case 0xf0a27848bbaeffc2:
		return &Node_OnLoadCheckFinished{}, nil
	case 0xaf7b779367274516:
		return &Node_OnBlackboardTimerExpired{}, nil
	case 0xbf881bd3ad19d00e:
		return &Node_OnWarpComplete{}, nil
	case 0x746ee266064e4228:
		return &Node_OnFocalZonePreActivationEnter{}, nil
	case 0xd49668ece1374231:
		return &Node_OnFocalZonePreActivationExit{}, nil
	case 0x515c5a812d44cb8c:
		return &Node_OnBanterStart{}, nil
	case 0x875031f40b6c907c:
		return &Node_OnBanterDone{}, nil
	case 0x42160ff9e068b081:
		return &Node_OnNextBanterDone{}, nil
	case 0x881ef391553740c9:
		return &Node_OnBanterFailed{}, nil
	case 0xc8da6c2869fe521f:
		return &Node_RemoveEquipmentFromWallet{}, nil
	case 0xe96efae266d22bb6:
		return &Node_OnUIMessageClosed{}, nil
	case 0x374b9370e2741624:
		return &Node_OnCallAndResponseTriggered{}, nil
	case 0x78028f7675b302d1:
		return &Node_OnCallAndResponseFinished{}, nil
	case 0x19f88e7729f15fec:
		return &Node_OnRegisteredFrameUpdate{}, nil
	case 0x8940960461f727b8:
		return &Node_OnAnimFrame{}, nil
	case 0xc52912a99e3029c2:
		return &Node_OnMovePlay{}, nil
	case 0xc94dc15319c6652a:
		return &Node_OnHealthChange{}, nil
	case 0xf3f35e11ad6680c9:
		return &Node_OnDeath{}, nil
	case 0xbdc033de21601dc5:
		return &Node_OnLookAt{}, nil
	case 0x0ea3372bcf3edaf3:
		return &Node_OnButtonPress{}, nil
	case 0x784b89e7c35a1b9d:
		return &Node_OnUIButtonPress{}, nil
	case 0x87c93b285cf1afd5:
		return &Node_OnControlPress{}, nil
	case 0x2c5d0590b77b7db1:
		return &Node_OnControlMash{}, nil
	case 0xb2efc2bf56af4552:
		return &Node_OnDistanceCheck{}, nil
	case 0xf0f5c75512f5557d:
		return &Node_OnEventQueueProcessed{}, nil
	case 0x18136aec00f11253:
		return &Node_OnRegisteredBreakableBroken{}, nil
	case 0x7b0c22f8691e872d:
		return &Node_OnRegisteredMeterChanged{}, nil
	case 0x3311c868a305e711:
		return &Node_OnControlSwipe{}, nil
	case 0x978d17b5d77e416d:
		return &Node_OnReticleTargetInvalid{}, nil
	case 0x1d0b8860ff8e4f75:
		return &Node_OnVariableChanged{}, nil
	case 0x9a5c017fe16eb2a7:
		return &Node_OnStartGameFromThisWad{}, nil
	case 0x3247023efaa55f82:
		return &Node_OnEncounterStart{}, nil
	case 0x3e347f0a44f8fc8b:
		return &Node_OnEncounterCreated{}, nil
	case 0xa331adf5bd27ea6c:
		return &Node_OnEncounterFinished{}, nil
	case 0xe41883d29611e9ed:
		return &Node_OnEncounterReset{}, nil
	case 0xb7dcf355d3ae7665:
		return &Node_OnEnemySpawn{}, nil
	case 0x6865f5f27fcd9a8b:
		return &Node_OnDynamicCharacterSpawn{}, nil
	case 0xe273b915d4249fe1:
		return &Node_OnEnemyDeath{}, nil
	case 0x49a41bd6dc3e2c6e:
		return &Node_OnInteractStart{}, nil
	case 0x695dc97456f1b038:
		return &Node_OnInteractFinished{}, nil
	case 0x2e8228b8c3c59c9c:
		return &Node_OnInteractAborted{}, nil
	case 0xc444784d3307e6c9:
		return &Node_OnInteractDone{}, nil
	case 0x71744f76784cc25d:
		return &Node_OnInteractAttemptedWhileOccupied{}, nil
	case 0xe20e4150a4da67c3:
		return &Node_OnTraverseGraphEntryEvent{}, nil
	case 0x14a3081dcbe7d1be:
		return &Node_OnTraverseGraphEnteredEvent{}, nil
	case 0x1923576cf1912be1:
		return &Node_OnCreatedVFSEntryChanged{}, nil
	case 0x7c088808310b8c8a:
		return &Node_OnCreatedVFSEntryExecuted{}, nil
	case 0xaefa146d7ad401fc:
		return &Node_OnConcussionHitBase{}, nil
	case 0x6e82923ca6494322:
		return &Node_OnArrowImpactBase{}, nil
	case 0x90daac765f2d2923:
		return &Node_OnCombatCollisionBase{}, nil
	case 0x0a2458fed1a8eea5:
		return &Node_OnWeaponThrown{}, nil
	case 0xe5d1ab0456da1291:
		return &Node_OnThrownWeaponBlockedBase{}, nil
	case 0x8ccb347d8f347e2a:
		return &Node_OnImmediateEvent{}, nil
	}
	return nil, fmt.Errorf("Unknown node type: %X", nodeId)
}
func ReadNode(file *os.File, stack []uint16, offset uint32) (Node, error) {
	_, err := file.Seek(int64(offset), 0)
	if err != nil {
		return nil, err
	}
	nodeIndex := uint16(0)
	err = binary.Read(file, binary.LittleEndian, &nodeIndex)
	if err != nil {
		return nil, err
	}
	if nodeIndex == 0xFFFF {
		return nil, nil
	}
	seen := false
	for _, v := range stack {
		if v == nodeIndex {
			seen = true
			break
		}
	}
	stack = append(stack, nodeIndex)
	// if _, ok := header.nodeMap[nodeIndex]; !ok {
	// 	header.nodeMap[nodeIndex] = uint32(len(header.nodeMap))
	// } else {
	// 	seen = true
	// }
	offset = nodeOffsets[nodeIndex]
	_, err = file.Seek(int64(offset), 0)
	if err != nil {
		return nil, err
	}
	var nodeId uint64
	err = binary.Read(file, binary.BigEndian, &nodeId)
	if err != nil {
		return nil, err
	}
	switch nodeId {
	case 0x50088220865F7D78:
		fallthrough
	case 0x28467E577E4BF816:
		fallthrough
	case 0x2D7869A0C6650E33:
		fallthrough
	case 0x85B43DEC04D96B7B:
		fallthrough
	case 0x703FD9308848087C:
		fallthrough
	case 0x30298D107985F295:
		fallthrough
	case 0x397D1DE8A6DE37B0:
		fallthrough
	case 0xAD686B21CD01B5B3:
		fallthrough
	case 0xC69765B16ED92CE5:
		fallthrough
	case 0x6FB55EE2FA4C6FEC:
		fallthrough
	case 0xCE0218F8A6C2B9F9:
		fallthrough
	case 0x8EB080F430E6BC3B:
		fallthrough
	case 0x8AAE018DF0E9C5B0:
		fallthrough
	case 0x02448C70A1903303:
		//fmt.Println("Warning: Nested Header nodes")
		return &Node_Dummy{}, nil
	case 0x339749F41CDA6CF3:
		return &Node_SetEffectsSubtitle{}, nil
	case 0x93B1E21ECE378ACB:
		return &Node_ShowUIItemCardMessage{}, nil
	case 0xDAE225FB7D050062:
		return &Node_ShowUIMessage{}, nil
	case 0xC041E37D26A4C555:
		return &Node_ShowUISidebarMessage{}, nil
	case 0xEC90133CE7D29A2E:
		return &Node_ShowUISplashScreenMessage{}, nil
	case 0x52C0B11E31001F7E:
		return &Node_ContextAction{}, nil
	case 0xA815ADC12F455A3F:
		return &Node_Ceil{}, nil
	case 0x6415879C6D60BE1F:
		return &Node_HasInteractZoneTag{}, nil
	case 0x1FCDE7BD7CAB4A12:
		return &Node_AsLootDistributor{}, nil
	case 0x416FCD56CBF9CB08:
		return &Node_IsPinType{}, nil
	case 0x73D5138B1D10A806:
		return &Node_GetCameraTargetPosition{}, nil
	case 0xF99BA6CE08A2E503:
		return &Node_GetActiveMovePercent{}, nil
	case 0xCBE7681D2D081500:
		return &Node_AsBanter{}, nil
	case 0x2A6DB82686190E01:
		return &Node_InteractZoneTemplate{}, nil
	case 0xBB587847BD5C2401:
		return &Node_GetScale{}, nil
	case 0x20083C263AAB8E02:
		return &Node_GetBlackboardVariable{}, nil
	case 0x178838A79AFC1504:
		return &Node_Icon{}, nil
	case 0x751884C159741E04:
		return &Node_GetFlawlessAresAvailable{}, nil
	case 0x979C59704EF99404:
		return &Node_IsInNewGamePlus{}, nil
	case 0xEDC78D0A7699A205:
		return &Node_GetSkillTreeTokensOfTypeOn{}, nil
	case 0xE87A372A7C499707:
		return &Node_CanCraftRecipe{}, nil
	case 0xFBD721EE0B2BDA06:
		return &Node_IsVolumeEnabled{}, nil
	case 0xBB362CCC34741507:
		return &Node_WalletHasEquipment{}, nil
	case 0x7F4BDEBE43771907:
		return &Node_IsSideQuestInAnotherRealm{}, nil
	case 0x4C25C63B98998607:
		return &Node_CheckProgressionFact{}, nil
	case 0x2E1B42EF540D9807:
		return &Node_WwiseRTPC{}, nil
	case 0x328187E7E557E707:
		return &Node_VectorDot{}, nil
	case 0x172E5CE9A06E6008:
		return &Node_GetVectorFromBlackboard{}, nil
	case 0x0266DCBCB18CCB08:
		return &Node_InstancedTraversePath{}, nil
	case 0x55507B2CFDF28A0D:
		return &Node_GetLastAttacked{}, nil
	case 0x1401E22515CB8109:
		return &Node_InstancedContextAction{}, nil
	case 0xF60CC1933C77D408:
		return &Node_GetClonedWeapons{}, nil
	case 0xB4401A111A7EE308:
		return &Node_GetContextActionApproachPoint{}, nil
	case 0x8A517611D6CB4009:
		return &Node_IsAnySoundPlaying{}, nil
	case 0x599B0A794FE05309:
		return &Node_IsFalling{}, nil
	case 0x9CFAF828B7AB9A0A:
		return &Node_GetPlayersCurrentRegion{}, nil
	case 0xDA7EB07E8079DA0A:
		return &Node_IsControlDown{}, nil
	case 0x945A99A6DFFC5C0B:
		return &Node_GetSeekTargetRemainingCount{}, nil
	case 0x579FEEFB38B0D60B:
		return &Node_IsInsideFocalZone{}, nil
	case 0xB106A6654350800F:
		return &Node_QueryEquipmentTrait{}, nil
	case 0xFB33A2532545D70D:
		return &Node_GetAnimJoint{}, nil
	case 0x7E6943A6AFEDC40E:
		return &Node_IsHostile{}, nil
	case 0x7013A1BADB00C80E:
		return &Node_VectorCross{}, nil
	case 0x3424C43128A3EF0E:
		return &Node_GetMeterName{}, nil
	case 0x3CA708DA86781810:
		return &Node_FindRandomPositionOnNavMesh{}, nil
	case 0x6521D1F12CF09B10:
		return &Node_AsMeter{}, nil
	case 0xC3C6BA077B56C710:
		return &Node_VectorAdd{}, nil
	case 0x3A50F9AE60CF1C11:
		return &Node_QuestHasAnyFlags{}, nil
	case 0x21BBBA62C3816017:
		return &Node_CheckFeatureData{}, nil
	case 0x368A6F5DB229CE14:
		return &Node_GetStickInput{}, nil
	case 0x04E1E41094515D13:
		return &Node_IsInteractZonePromptEnabled{}, nil
	case 0x1885624E8D82B412:
		return &Node_CreatureBehaviorSetOptionData{}, nil
	case 0x0EFE86250CACCF12:
		return &Node_Effect{}, nil
	case 0x6D4996EE826AEF12:
		n := &Node_Subtract{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xA16AD5E6F6402613:
		return &Node_PlayerEitherHasOrCanCraftEquipmentRightNow{}, nil
	case 0xA0A140DEDA0B6313:
		return &Node_AsPlayFX{}, nil
	case 0xC8A240163665F613:
		return &Node_GetShieldValue{}, nil
	case 0x965120C0AFB71814:
		return &Node_Equipment{}, nil
	case 0xA4D12B9B6FD28714:
		return &Node_IsUnoccupied{}, nil
	case 0xA12BF9D5B5BA6C15:
		return &Node_GetStringHashValueFromElement{}, nil
	case 0x854A33EA70E6EA14:
		return &Node_IsCollisionEnabled{}, nil
	case 0xBD53422167DE0B15:
		return &Node_IsAttackableArrow{}, nil
	case 0x836EA18B55D51715:
		return &Node_AsBranch{}, nil
	case 0xFE4A56C4F82F5415:
		return &Node_GetAnimationLength{}, nil
	case 0xBACD3F39DCD98E15:
		return &Node_IsMovementLimitedByCurrentAOO{}, nil
	case 0x5F812DF89D2B2E16:
		return &Node_GetFloatValueFromElement{}, nil
	case 0x4610D0BAC8516A16:
		return &Node_GetAvailabilityStateBySource{}, nil
	case 0xCE06F1CE3BC99416:
		return &Node_GetAllInteractsDisabled{}, nil
	case 0xC141BC765096711D:
		return &Node_ACos{}, nil
	case 0x42D173CFEBD58B1A:
		return &Node_AsPickupSlot{}, nil
	case 0x99766FC467773A19:
		return &Node_IsContainedInArray{}, nil
	case 0xA93F370F32B35519:
		return &Node_MakeVector{}, nil
	case 0x8D6670BB9249C119:
		return &Node_FullScreenEffect{}, nil
	case 0x6F14F08AB5B9EC19:
		n := &Node_AsWwiseEvent{}
		n.param, err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xFE2E85C681F4721B:
		return &Node_ASin{}, nil
	case 0xC659BC3FC5998A1B:
		return &Node_GetActiveWeapons{}, nil
	case 0xDE24398CE73F781C:
		return &Node_AsWallet{}, nil
	case 0xA62DE0C1AEAF3E1D:
		return &Node_ArrowEmitter{}, nil
	case 0x54D0297AABD2E71E:
		return &Node_GetOptionControlUpForSetting{}, nil
	case 0xC9B10F8369EA261E:
		return &Node_WasWeaponReflected{}, nil
	case 0x8DD1EC5A6C75541E:
		return &Node_GetAnimDriverValue{}, nil
	case 0x036E13BEA402A71E:
		return &Node_GetDockTargetType{}, nil
	case 0x74C487E17719C31E:
		return &Node_IsDynamicCharacterSpa{}, nil
	case 0x62FA573E079D0E1F:
		return &Node_EncounterElementData{}, nil
	case 0x768BDC18FCF91D1F:
		return &Node_BeamAttack{}, nil
	case 0xB08AB75C16F69E1F:
		return &Node_Rand{}, nil
	case 0x4C9EAB4ACBE1CE2F:
		return &Node_IsMentalStateEqual{}, nil
	case 0xE1438E33E5989F28:
		return &Node_IsSequenceSkipping{}, nil
	case 0x50C0D28308EF2124:
		return &Node_GetEquipmentInWallet{}, nil
	case 0xB304B794456B0322:
		return &Node_GetRegionForMapMarker{}, nil
	case 0x27EEAAE00585A120:
		return &Node_GetTyr{}, nil
	case 0x3E12DE27B934C920:
		return &Node_HasResurrectionStoneBeenUsed{}, nil
	case 0x2858D742F7EC5E21:
		return &Node_MakeVector4{}, nil
	case 0x2508E084C4C1A621:
		return &Node_GetCameraOrbitLeft{}, nil
	case 0xFBA91B0E60677C22:
		return &Node_GetVFSInt{}, nil
	case 0x53EF795BE8954523:
		return &Node_IsValidEquipment{}, nil
	case 0x92D8229CA2E56823:
		return &Node_IsMarkerClaimed{}, nil
	case 0xAF7B822BADB36D23:
		return &Node_GetTouchpadSwipeMapped{}, nil
	case 0xF546D53574C7D824:
		return &Node_RecipeHasFlag{}, nil
	case 0x2747BC2296105E24:
		return &Node_GetBreakableStage{}, nil
	case 0x4F01A73671379024:
		return &Node_IsTouchpadSwipeMapped{}, nil
	case 0x33E132C52E35C624:
		return &Node_TraverseLink{}, nil
	case 0xD7D95157B103D824:
		return &Node_HasHitFlag{}, nil
	case 0x67CF4AEA90A29126:
		return &Node_CanEquipToEquipmentSlot{}, nil
	case 0xEA9343A22E819927:
		return &Node_ChooseByCondition{}, nil
	case 0x3733EAFEDB939828:
		return &Node_Enum{}, nil
	case 0xB3ECE0E6AA649E28:
		return &Node_IsSvrContextConnected{}, nil
	case 0x8B06A6B7BB50B12B:
		return &Node_IsTraverseLinkEnabled{}, nil
	case 0x0492A60321D4502A:
		return &Node_GetEquipmentInCharacterSlot{}, nil
	case 0xE1EB3646F8125129:
		return &Node_GetListenerPositionInfo{}, nil
	case 0xD5729238CE147829:
		return &Node_LootCondition{}, nil
	case 0x3019B268A85CDD29:
		return &Node_MapMarkerHasAllFlags{}, nil
	case 0xC8D2110D854D172A:
		return &Node_IsWeatherCategoryActive{}, nil
	case 0x60D02BF79E7A562A:
		return &Node_Max{}, nil
	case 0x64E850E126946F2A:
		return &Node_GetContextActionFromBlackboard{}, nil
	case 0xF365BC694CD33C2B:
		return &Node_GetWolfSledInertia{}, nil
	case 0xE2A49D74E9AF732B:
		return &Node_AsConcussion{}, nil
	case 0x7A2E56CF1416602D:
		return &Node_IsPlatform{}, nil
	case 0xD24FAD9AC019C32B:
		return &Node_AsNavCurve{}, nil
	case 0x737E8ACFD9D8FB2B:
		return &Node_IsRealmLocked{}, nil
	case 0xFDA150728F35972C:
		return &Node_GetCameraOrbit{}, nil
	case 0xB8C25393BE0B212D:
		return &Node_AsWwiseSwitchGroup{}, nil
	case 0x4FDD3EF36DD8402E:
		return &Node_AsSoundEmitter{}, nil
	case 0x0FEE9E195929AF2E:
		return &Node_GetPlayersCurrentArea{}, nil
	case 0xEE8370865260642F:
		return &Node_GetScriptOwningWad{}, nil
	case 0xF9A854988056AE37:
		return &Node_Crank_GetCurrentC{}, nil
	case 0x1C8ED91C2E90C233:
		return &Node_GetAvailabilityState{}, nil
	case 0xDA6EFAEFA5BA8B32:
		return &Node_IsRouteToOtherRealmsViaFastTravel{}, nil
	case 0x590600725C75FB2F:
		return &Node_SortIterator{}, nil
	case 0x0E2A52DDE8961E30:
		return &Node_AsInteractZoneTemplate{}, nil
	case 0x2D5B2216780C8D30:
		return &Node_GetTimerFromBlackboard{}, nil
	case 0x430644D86593D131:
		return &Node_GatewayMarker{}, nil
	case 0x090EEFC2F2FE1633:
		return &Node_HasPositionalFlag{}, nil
	case 0x10B8A1D632732533:
		n := &Node_VectorLength{}
		n.vector, err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x96AC2853E2FF5F33:
		return &Node_AreNavCurvesEnabledForObjects{}, nil
	case 0xC88C515AC2B99433:
		return &Node_GetCompassTargetPositionLerpSpeed{}, nil
	case 0x8B7837ECFBDE4B35:
		return &Node_HitFlagsTo_String{}, nil
	case 0x359F14D86ADB4834:
		n := &Node_Add{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x0E91834533D78A34:
		return &Node_MemoryAvailableToSpawnCreature{}, nil
	case 0x5E5476DACFA62735:
		return &Node_HasAnimationComponent{}, nil
	case 0xF5F3D48E34812835:
		return &Node_GetVelocity{}, nil
	case 0xF81FEE9567519E35:
		return &Node_Branch{}, nil
	case 0x7D9B9F4AC84B2A36:
		return &Node_IsPlayGoDownloaded{}, nil
	case 0x629DD5298A06A036:
		return &Node_ChooseByInt{}, nil
	case 0xB951B756C2B29437:
		return &Node_GetKratos{}, nil
	case 0x16CDEBE8FBC38F3C:
		return &Node_IsPickupSlotOnCooldown{}, nil
	case 0x62F04BFDB01FAB3A:
		return &Node_GetCompanionReticleHitData{}, nil
	case 0x832B6E7F4A3BCC37:
		return &Node_GetCurrentOptionIndexForSetting{}, nil
	case 0xF926409DB8A42838:
		return &Node_AreBehaviorsEnabled{}, nil
	case 0x5F7A24CB8F4A8738:
		return &Node_GetEquipmentInWalletWithFlags{}, nil
	case 0xB240B78A528DBD38:
		return &Node_Sin{}, nil
	case 0x296D5FA5EDCFD43A:
		return &Node_GetBreakableHitPoints{}, nil
	case 0x389529A0B866E73A:
		return &Node_GetGameObjectFromBlackboard{}, nil
	case 0x4B05B1DF22693A3B:
		return &Node_GetGroundLevel{}, nil
	case 0xEA26BA1CBCCB583C:
		return &Node_GetVisibility{}, nil
	case 0x72226F5F3F62AA3D:
		return &Node_RecipeHasFlags{}, nil
	case 0xB0808B8786A0F13C:
		return &Node_GetPathDirection{}, nil
	case 0xA6B7FF47FDAC773D:
		return &Node_GetBestNavmeshPosInAOO{}, nil
	case 0xF88EF926B1E7993D:
		return &Node_GetFightPosition{}, nil
	case 0x82F5BE7ED96C9F3D:
		return &Node_GetTimeSinceBanterPlayed{}, nil
	case 0x5F06BFA64EA0E23D:
		return &Node_ShuffleIterator{}, nil
	case 0x53C1A8302A92AE3E:
		return &Node_IsButtonDown{}, nil
	case 0x5F47207ACCE6C23E:
		return &Node_AreContextActionsEnabled{}, nil
	case 0x6F360CF22AD4A461:
		return &Node_CreatureHasAOO{}, nil
	case 0xCD17C3E5A0EEDA50:
		return &Node_AsCreature{}, nil
	case 0x02CEEBE3D95E5A46:
		return &Node_GetStringLength{}, nil
	case 0x06559D1FE9AA1E41:
		return &Node_LootConditionSet{}, nil
	case 0x57D2EFC714904540:
		return &Node_SoundPortal{}, nil
	case 0x47758251DEC08E3F:
		return &Node_GetClosestFastTravelMarkerPos{}, nil
	case 0x7F27A7F5A1AEEE3F:
		return &Node_GetQuestName{}, nil
	case 0x1AA9354D22941A40:
		return &Node_Weapon{}, nil
	case 0x4E70A52DF80D1F40:
		return &Node_AsEquipment{}, nil
	case 0xB83A779AA5CEC540:
		return &Node_GetWadName{}, nil
	case 0xD039DA99EC66EC40:
		return &Node_GetSetting{}, nil
	case 0xE7BC749F62140F41:
		return &Node_GetSpawnFacing{}, nil
	case 0x90F135F0B3731241:
		n := &Node_Or{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xA496A11065C04B44:
		return &Node_GetPickupStageCount{}, nil
	case 0xF316CC9E82359341:
		return &Node_ScreenShake{}, nil
	case 0x4F5A9564A3073E43:
		return &Node_IsPickupSlotActive{}, nil
	case 0xE3DA123C99CC3F43:
		return &Node_GetChildren{}, nil
	case 0xDD8BEDFA5C28A043:
		return &Node_MapRealm{}, nil
	case 0x2524992DB41E8F44:
		return &Node_CompassPathDistance{}, nil
	case 0x909250A5DBAFA444:
		return &Node_InteractZoneLocked{}, nil
	case 0x01DE231A03FF3246:
		return &Node_WwiseSwitchState{}, nil
	case 0x54A2056ABB5E5146:
		return &Node_AsMapArea{}, nil
	case 0xD6760CAB6E13744C:
		return &Node_IsPositionOnScreen{}, nil
	case 0x8730F1E200545C49:
		return &Node_EquipmentHasFlags{}, nil
	case 0x3B6175BB1E297F46:
		return &Node_GetPlayerBoat{}, nil
	case 0xE0A740A87F7F2C47:
		return &Node_GetQuestProgressAndGoal{}, nil
	case 0x7DF7E88CA9F83F47:
		return &Node_GetObjectsWithFlagsInRadius{}, nil
	case 0x8FDD6E92AD8F2F49:
		return &Node_GetMotionDebuggerRecordLocation{}, nil
	case 0x3F19FD355521B249:
		n := &Node_GetVariable{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		variableOffset := int64(0)
		err = binary.Read(file, binary.LittleEndian, &variableOffset)
		if err != nil {
			return nil, err
		}
		if variableOffset>>0x3e == 1 {
			variableOffset = (variableOffset << 1) / 2
		}
		_, err = file.Seek(int64(offset+0x10)+variableOffset, 0)
		if err != nil {
			return nil, err
		}
		strLength := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &strLength)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x48)+variableOffset, 0)
		if err != nil {
			return nil, err
		}
		strData := make([]byte, strLength)
		_, err = file.Read(strData)
		if err != nil {
			return nil, err
		}
		n.name = string(strData)
		return n, nil
	case 0x29AA1F002C21BD49:
		return &Node_AsAnimation{}, nil
	case 0x7A22EB2D80162B4A:
		return &Node_GetVFSFloat{}, nil
	case 0x6EB6D1128093194B:
		return &Node_GetCreatureGroun{}, nil
	case 0x5DC16FB7F65C0750:
		n := &Node_Numeric{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.val)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x41B277201D11AA4C:
		return &Node_IterateFixedSizeArray{}, nil
	case 0xB912D0ADB49B364D:
		return &Node_GetCurrentOptionJoysticForSetting{}, nil
	case 0x285649A3EF25BD4E:
		return &Node_GetExistingResourcesInWallet{}, nil
	case 0xCEEB0BD5841CE64F:
		return &Node_GetSpawnPosition{}, nil
	case 0x1E533D3CC3573550:
		return &Node_GetBoolValueFromElement{}, nil
	case 0xFA79944964216550:
		return &Node_GetTargetCreature{}, nil
	case 0x3188431D706FCF50:
		return &Node_AsQuest{}, nil
	case 0x25725C0309FA4C57:
		return &Node_GetAngleBetweenVectors{}, nil
	case 0xDFDBD89284F95D53:
		return &Node_IsNavAssistCompassMarker{}, nil
	case 0xF04679388D8E5B52:
		return &Node_IsPlayerDebugMovement{}, nil
	case 0x9F51B67FA6E3DF50:
		return &Node_DidCreatureRecentlyHaveFlag{}, nil
	case 0x64F5CCB6B5E23351:
		return &Node_GetPendulumAngle{}, nil
	case 0x7EB1931EF4B98151:
		return &Node_Prop{}, nil
	case 0x38FF1281E4B0B151:
		return &Node_AsSoundProximityTrigger{}, nil
	case 0x296D624319FE9752:
		return &Node_GetMaxSpeedOverride{}, nil
	case 0x817F4A9FA9E9B252:
		return &Node_IsInValhalla{}, nil
	case 0x81613FB023BFCA52:
		return &Node_GetWeaponGameObject{}, nil
	case 0x8ACA12277C431753:
		return &Node_GameObject{}, nil
	case 0x58C352A97FFA1C55:
		return &Node_GetCurrentSpline{}, nil
	case 0x53370E89D3090C54:
		return &Node_GetAnimationTime{}, nil
	case 0xF5516FD8F533A654:
		return &Node_AsMapRealm{}, nil
	case 0xE68599F5C63BC254:
		return &Node_IsSummaryCategoryInMapRegion{}, nil
	case 0x67A9E17EF62DFE54:
		return &Node_GetMapMarkersInArea{}, nil
	case 0x3A7FDFDF1BE57355:
		return &Node_GetCollisionRadius{}, nil
	case 0xEEA72C87452D6156:
		return &Node_GetMPIconObjectByName{}, nil
	case 0x0D9F978FC8B76B56:
		return &Node_GetName{}, nil
	case 0xE207026532A1FC56:
		return &Node_InstancedNavCurve{}, nil
	case 0x424A385421D7B85B:
		return &Node_AsCamera{}, nil
	case 0x7BEED5AF1D6C6458:
		return &Node_GetWolfSledHarnessVisibility{}, nil
	case 0x68D7787552A5C957:
		n := &Node_HasBlackboardVariable{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x14), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.param)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x48B6B23A1686D157:
		return &Node_IsPickupSlotFree{}, nil
	case 0xD7A644434EDCDE57:
		return &Node_GetAllSoundEmittersOnObject{}, nil
	case 0x6A0FDC973A484F58:
		return &Node_AsTweakConstant{}, nil
	case 0xBD75554DAE3D155A:
		return &Node_InstancedInteractZone{}, nil
	case 0x52753D5851D03E5A:
		return &Node_IsInsideVolume{}, nil
	case 0x0E946842DE6BEB5A:
		return &Node_GetCompanion{}, nil
	case 0x7CCFE6321089EB5A:
		return &Node_GetCreatureUnderReticle{}, nil
	case 0x6B2C1C2EBD44DA5D:
		return &Node_GetQuestState{}, nil
	case 0xDA57C4E5FC9EFA5C:
		return &Node_VectorScale{}, nil
	case 0x66F2867B59960D5D:
		return &Node_VectorSubtract{}, nil
	case 0x0C46EAEC68BF6D5D:
		return &Node_ArePhysicsEnabled{}, nil
	case 0x9B9D492BD58C865D:
		return &Node_GetArrowOwner{}, nil
	case 0xE0DDB57F9A80035E:
		return &Node_GetCurrentOptionControlDownForSetting{}, nil
	case 0x5F946D06C162CD5E:
		return &Node_ContextActionSequencerStage{}, nil
	case 0x548DB112F4929960:
		return &Node_GetPendulumSpeed{}, nil
	case 0x6E7A5956333D8172:
		return &Node_IsValidReference{}, nil
	case 0xECEC3A6FFED31B6A:
		return &Node_GetBranchValueFromElement{}, nil
	case 0xD0B482CC18289467:
		return &Node_IsTargetObstructed{}, nil
	case 0x31EBD975352DE763:
		return &Node_AsRecipe{}, nil
	case 0xCA5D71C1AEF52862:
		return &Node_IsPickupActive{}, nil
	case 0x3FA7C4E1990ED462:
		return &Node_GetPlayerDeathsOnCurrentCheckpoint{}, nil
	case 0xDB2571EF4F5E3663:
		n := &Node_CastToInt32{}
		n.param, err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x73FEE5BB22D38E63:
		return &Node_GetWolfSledRopeVisibility{}, nil
	case 0x5172760831106664:
		return &Node_CanAim{}, nil
	case 0xDF47AA8D47943565:
		return &Node_Floor{}, nil
	case 0x371BB06DBB9F1466:
		return &Node_BitwiseAnd{}, nil
	case 0x05340FB9C978D166:
		return &Node_GetAllRecipesInWallet{}, nil
	case 0xD89AD93F11DC6068:
		return &Node_SoundEmitter{}, nil
	case 0x91112EB0D18C9967:
		return &Node_ContextActionSequencerStageData{}, nil
	case 0x4C3B8555AF109E67:
		return &Node_GetWeaponThrowStatus{}, nil
	case 0xF82D881709B01368:
		return &Node_AsWwiseSwitchState{}, nil
	case 0xE8FB5449C6845868:
		n := &Node_IsPlayingMove{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x53A03EE076770569:
		return &Node_GetMappedButton{}, nil
	case 0xD3C9C0F0B6665A69:
		return &Node_IsWeatherSystemActive{}, nil
	case 0xDC94522861C46769:
		return &Node_MapMarkerHasAnyFlags{}, nil
	case 0xA167F4F521BACD69:
		return &Node_HasBeenUsed{}, nil
	case 0xBA395A6FAE49776C:
		return &Node_IsOnNavMesh{}, nil
	case 0xF511929C12B9606B:
		return &Node_GetUniqueEquipmentByName{}, nil
	case 0xB4A03C1D27573B6A:
		return &Node_Sqrt{}, nil
	case 0x87F2A2ED0F07936A:
		return &Node_IsAnyBanterPlaying{}, nil
	case 0xD71817507057166B:
		return &Node_GetAnimationFrame{}, nil
	case 0x48691295C08A3E6B:
		return &Node_GetRandomPositionInVolume{}, nil
	case 0x2AB7D87BF8CA786B:
		return &Node_GetSplinePosition{}, nil
	case 0x4DDEDF5ECAE7F36B:
		return &Node_EncounterWaveData{}, nil
	case 0x2A0829F19FCC146C:
		return &Node_GetCameraOrbitForward{}, nil
	case 0x01F0BD04F14F196C:
		n := &Node_Bool{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		bytes := make([]byte, 1)
		_, err = file.Read(bytes)
		if err != nil {
			return nil, err
		}
		n.val = bytes[0] != 0
		return n, nil
	case 0xBF9FC5F418875D70:
		return &Node_Abs{}, nil
	case 0xE0E82EC16CF5B56D:
		return &Node_GetSpawnedObjectVariable{}, nil
	case 0xFEBD20D79F2CB86E:
		n := &Node_GetTweakFloatConstant{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x5D01376A7A7F866F:
		return &Node_GetRegionForWad{}, nil
	case 0x3261F73DA77F4270:
		return &Node_AsWildlife{}, nil
	case 0x52757504C8607C70:
		return &Node_AsPickup{}, nil
	case 0x17E1652A4CAE2071:
		return &Node_AsAnimJoint{}, nil
	case 0x66AF364938109171:
		return &Node_IsPlayerOnWolfSled{}, nil
	case 0xF3DD8834B3F29777:
		return &Node_Wallet{}, nil
	case 0x2F5543725EFB2F75:
		return &Node_QuestHasAllFlags{}, nil
	case 0x0B0F5B516621CB73:
		return &Node_GetDepthOfField{}, nil
	case 0xDF387335D5B2A072:
		return &Node_GetFriendlyCreatures{}, nil
	case 0xD87B24DEAFE7D772:
		return &Node_AsColor{}, nil
	case 0x600A509FDEFD2473:
		return &Node_IsMarkerLocked{}, nil
	case 0x62C6D8B2E4486973:
		n := &Node_Divide{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xB183BC7D15D37074:
		return &Node_PointOnPath{}, nil
	case 0x866D92FD752FDE74:
		return &Node_GetBitFromVariable{}, nil
	case 0xF8DA23BAFF35F074:
		return &Node_IsHighlightDisableOverrideActive{}, nil
	case 0xDB6C47BB54310C75:
		return &Node_AsProp{}, nil
	case 0x379986296E5A0776:
		n := &Node_Pickup{}
		_, err := file.Seek(int64(offset+0x20), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.namespaceHash)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x30), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.nameHash)
		if err != nil {
			return nil, err
		}
		if n.namespaceHash == 0 && n.nameHash == 0 {
			n.name = ""
			n.namespace = ""
			return n, nil
		}
		_, err = file.Seek(int64(offset+0x48), 0)
		if err != nil {
			return nil, err
		}
		strLength := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &strLength)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x58), 0)
		if err != nil {
			return nil, err
		}
		strData := make([]byte, strLength)
		_, err = file.Read(strData)
		if err != nil {
			return nil, err
		}
		n.namespace = string(strData)
		_, err = file.Seek(int64(offset+0x38), 0)
		if err != nil {
			return nil, err
		}
		nameOffset := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &nameOffset)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x38+nameOffset), 0)
		if err != nil {
			return nil, err
		}
		strLength = uint32(0)
		err = binary.Read(file, binary.LittleEndian, &strLength)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x48+nameOffset), 0)
		if err != nil {
			return nil, err
		}
		strData = make([]byte, strLength)
		_, err = file.Read(strData)
		if err != nil {
			return nil, err
		}
		n.name = string(strData)
		return n, nil
	case 0x65941400E52C4E75:
		n := &Node_Vector{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.x)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x14), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.y)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x18), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.z)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x713BB0F002E84E75:
		return &Node_AsRumble{}, nil
	case 0x3B7CACFA2AB58E75:
		return &Node_GetIntFromBlackboard{}, nil
	case 0xBC9F575C8C84B175:
		return &Node_IsAreaOfOperationEnabled{}, nil
	case 0x71688B18653EB276:
		return &Node_IsPickupAvailable{}, nil
	case 0x99733C6F0EC3CC76:
		return &Node_AsEnum{}, nil
	case 0xE2E68DA846182377:
		return &Node_GetLootObjectResult{}, nil
	case 0x399B421B678F8877:
		return &Node_IsWeaponEmbedded{}, nil
	case 0xFB29C935F9425E7A:
		return &Node_GetLootConditionByName{}, nil
	case 0x6C2E3D80AF897379:
		return &Node_GetAngrboda{}, nil
	case 0x1B9B09F8AD74B777:
		return &Node_InstancedSoundEmitter{}, nil
	case 0xBCC4F6E9DCD1C678:
		return &Node_AsGatewayMarker{}, nil
	case 0xE693829E4EB7FF78:
		return &Node_IsAOOAssignmentTypeEqualTo{}, nil
	case 0x3B9ED4FD51822A79:
		return &Node_GetCurrentContextAction{}, nil
	case 0xDF980BF9BB3D9179:
		return &Node_IsSpawnEnemiesEnabled{}, nil
	case 0x49E297A263D60C7A:
		return &Node_ATan{}, nil
	case 0xC81FFF2816BB227A:
		n := &Node_CastToFloat{}
		n.param, err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x01C63D2F87242E7A:
		return &Node_GetClosestPositionInVolume{}, nil
	case 0x161D1ED9AC60C67C:
		return &Node_IsDoingSyncMove{}, nil
	case 0xFA326E83B480777A:
		return &Node_ShouldPointToFastTravelMarker{}, nil
	case 0x7F107D84E6CA837A:
		return &Node_AreContextActionsEnabledForObjects{}, nil
	case 0x502257241BE8C17A:
		return &Node_FindPath{}, nil
	case 0xD3AD0C02C6BD677B:
		return &Node_GetAtreus{}, nil
	case 0x9151FC57AA7FDB7C:
		n := &Node_CastToUInt64{}
		n.param, err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x0CFBFEBAF7DEE47C:
		return &Node_GetPlayerLockOnTarget{}, nil
	case 0x88752874E4D1D07D:
		return &Node_IsPlayerInBoat{}, nil
	case 0x7F61F9A8EFB262C1:
		return &Node_Min{}, nil
	case 0x9E31788953E2759F:
		return &Node_WasEncounterRunning{}, nil
	case 0x229D87CD9FECB28F:
		return &Node_CanEquipToCharacterSlot{}, nil
	case 0x7456B1DC7113F583:
		n := &Node_GetCreaturesInRadius{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x7B7EA83A406A8781:
		return &Node_GetAllSkillsInTree{}, nil
	case 0xF8C8E79ACA29BB7F:
		n := &Node_IsWeaponActive{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x9624AED0F6312C7E:
		return &Node_GetCurrentVehicleCreature{}, nil
	case 0x42146239F457CE7E:
		return &Node_AsFullScreenEffect{}, nil
	case 0x688A0263E7A5267F:
		return &Node_GetMapMarkersInRegion{}, nil
	case 0xD565F65DB8F09C7F:
		return &Node_GetClosestFastTravelMarker{}, nil
	case 0x1E0C7BA7D5562380:
		return &Node_GetGIBlend{}, nil
	case 0x18ACEFF4DF05B080:
		return &Node_GetSideQuestRealm{}, nil
	case 0x43044349DCC42F81:
		return &Node_IsJumping{}, nil
	case 0xD17577A116198381:
		return &Node_IsWolfSledSpawned{}, nil
	case 0x886AC326F68F9082:
		return &Node_AsVector{}, nil
	case 0xB7BEE12B6609C681:
		return &Node_GetFirstEquippableEquipmentSlot{}, nil
	case 0x4DA747B96E332782:
		return &Node_GetEquipmentGeneratorFromHandle{}, nil
	case 0x0752554762683C82:
		return &Node_GetLookAtTarget{}, nil
	case 0x6A7B1B6023C45282:
		return &Node_IsPlayerInsideLeashZone{}, nil
	case 0x73252C8B0585B682:
		return &Node_GetNavCurveLength{}, nil
	case 0x6E7476D289B8DC82:
		return &Node_AsGameObject{}, nil
	case 0x3CD62D69A2EC7383:
		return &Node_Rumble{}, nil
	case 0xA2153E07B938BF83:
		return &Node_IsAnyOtherSoundPlaying{}, nil
	case 0xEA49A8F3CFF5F78A:
		n := &Node_CastToBool{}
		n.param, err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xF45CEC5CEC084788:
		return &Node_IsAvailabilitySourceRequesting{}, nil
	case 0x85001198E0113284:
		return &Node_WalletHasRecipe{}, nil
	case 0x2665688A6083BD85:
		return &Node_BreakVector{}, nil
	case 0x0C9DFAA307379986:
		return &Node_GetMainQuestRealm{}, nil
	case 0xEA52A31ED9223087:
		return &Node_GetVFSTokenFromPath{}, nil
	case 0xD4B0603E13F26388:
		return &Node_GetButtonIndex{}, nil
	case 0x86154A1358F01389:
		return &Node_IsModelEnabled{}, nil
	case 0xEC7E4E1F46C54289:
		return &Node_GetOptionJoysticForSetting{}, nil
	case 0x313283CDD133948A:
		return &Node_GetPickupStage{}, nil
	case 0x0D82A05038A2148E:
		return &Node_GetWolfSledIsInDrift{}, nil
	case 0x6871ED3EF297578B:
		return &Node_GetOptionControlDownForSetting{}, nil
	case 0x0C818064BBE55C8C:
		return &Node_GetBooleanFromBlackboard{}, nil
	case 0xD808964ED96BF48C:
		return &Node_AsSoundPortal{}, nil
	case 0x55055BAA233BF98D:
		return &Node_CameraRecenter{}, nil
	case 0xEE053C4B3B89168E:
		n := &Node_And{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xFDF5A13C60A8D08E:
		return &Node_CreatureOptionKeyValuePairData{}, nil
	case 0x18D62C0E1139368F:
		return &Node_AreLightsEnabled{}, nil
	case 0x55B391042B8F9199:
		n := &Node_StringHash{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.hash)
		if err != nil {
			return nil, err
		}
		if n.hash == 0 {
			n.str = ""
			return n, nil
		}
		_, err = file.Seek(int64(offset+0x28), 0)
		if err != nil {
			return nil, err
		}
		strLength := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &strLength)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x38), 0)
		if err != nil {
			return nil, err
		}
		strData := make([]byte, strLength)
		_, err = file.Read(strData)
		if err != nil {
			return nil, err
		}
		n.str = string(strData)
		return n, nil
	case 0x7496C50263153294:
		return &Node_Round{}, nil
	case 0x6F95E7B322A76492:
		return &Node_ResourceHasFlags{}, nil
	case 0xE6618B4E12E2FA8F:
		return &Node_GetCameraRender{}, nil
	case 0x2A50EEC9FC722C90:
		return &Node_AsLootCondition{}, nil
	case 0x426F8EEDDB208790:
		return &Node_GetVFSBool{}, nil
	case 0x1ED5AA8BD2EA2B91:
		return &Node_AsMapRegion{}, nil
	case 0xC4DF3851C8388992:
		return &Node_IsInteracting{}, nil
	case 0x92A9AA86DD4AF292:
		return &Node_IsPositionInCreatureAOO{}, nil
	case 0xADEC9B8C3BCF7593:
		return &Node_IsInContextAction{}, nil
	case 0x7D9AB7B48DA3E593:
		return &Node_MapAreaHasAnyFlags{}, nil
	case 0x832146796071F695:
		return &Node_GetWolfSummoningPoints{}, nil
	case 0xA70EBA6B08AEC094:
		return &Node_BitwiseOr{}, nil
	case 0xA1FD87CAE6B41695:
		return &Node_IsSideQuestPathfindingNonContiguous{}, nil
	case 0xDC56F77D99994895:
		return &Node_GetMapRegionsInRealm{}, nil
	case 0xC83F0A500667AF95:
		return &Node_GetElement{}, nil
	case 0x49B4DE5874EE4496:
		return &Node_Clamp{}, nil
	case 0x5DD2C814B4B10E98:
		return &Node_IsDying{}, nil
	case 0x1C6EAE6464D5FB98:
		n := &Node_Float{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.val)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xCCAE45D57C4D3B99:
		return &Node_GetSkillTreeTokenIn{}, nil
	case 0x717948816C00289D:
		return &Node_GetMusicIntensity{}, nil
	case 0x0E07073B5B47379B:
		return &Node_CanAcquireSkill{}, nil
	case 0xBF4420A11121B999:
		return &Node_GetCreatureByName{}, nil
	case 0xBDE036C40307D699:
		return &Node_GetEnemiesInRadius{}, nil
	case 0xAA1CEBFAF632069A:
		return &Node_NavCurve{}, nil
	case 0xDF0A4CA6DED6019B:
		return &Node_AreNavCurvesEnabled{}, nil
	case 0x0E722736994A389B:
		return &Node_GetViewport{}, nil
	case 0xB396F09C2F24479B:
		return &Node_PartFlagsTo_String{}, nil
	case 0xC63D04A31890649C:
		return &Node_Resource{}, nil
	case 0xD3CCB41E23B5119D:
		return &Node_GetSplineProgression{}, nil
	case 0xBD9E097279CC6C9E:
		return &Node_CreatureHasBehaviorTree{}, nil
	case 0xDA7DC9B7184C5D9D:
		return &Node_GetVariableInfo{}, nil
	case 0x63DF4CBFE1C47C9D:
		return &Node_Wildlife{}, nil
	case 0x5E8E6105EC5D889D:
		return &Node_IsWadLoaded{}, nil
	case 0xA9244278D9E9AC9D:
		return &Node_HashString{}, nil
	case 0x281472EF290AB89E:
		return &Node_CreatureOptionData{}, nil
	case 0xBF74532C78553B9F:
		return &Node_IsInsideLeashZone{}, nil
	case 0xC405E1E70CEB629F:
		return &Node_GameMap_IsItemStateV2{}, nil
	case 0xF2609AA50A0D81AF:
		return &Node_GetMaterialEntryValue{}, nil
	case 0x8358A69A1957F3A4:
		return &Node_GetAllInteractZones{}, nil
	case 0xB086AE0D1C053AA3:
		return &Node_IsInPlaytest{}, nil
	case 0x477EE3F0251656A1:
		return &Node_GetYakWaterHeight{}, nil
	case 0xC37E656C704E42A0:
		return &Node_GetFirstChild{}, nil
	case 0x330620E6553D5AA0:
		return &Node_GetIntForContextActionEnum{}, nil
	case 0xAFEFF0AEE067D5A0:
		return &Node_Haptic{}, nil
	case 0xDF830E2FF3BA38A1:
		return &Node_MapRegion{}, nil
	case 0xAD105E0C2CCE0BA2:
		return &Node_GetBanterDuration{}, nil
	case 0x620B22C1F08347A2:
		return &Node_GetControlIndex{}, nil
	case 0xD7DD89E55389A0A2:
		return &Node_Recipe{}, nil
	case 0x34C3C285D8C5DAA2:
		return &Node_ATan2{}, nil
	case 0xAD17799B143419A4:
		return &Node_IsBreakableBroken{}, nil
	case 0x8772407649B990A3:
		return &Node_MapRegionHasAnyFlags{}, nil
	case 0x34438E5351AFB3A3:
		return &Node_VectorDistance{}, nil
	case 0x2E0C717DF51309A4:
		return &Node_InstancedSoundProximityTrigger{}, nil
	case 0xD9E36B38CB4D0BA4:
		return &Node_GetCurrentOptionControlUpForSetting{}, nil
	case 0xDDF15EBFE9101EA4:
		return &Node_GetAssociatedGameObject{}, nil
	case 0x1D501527749595A4:
		n := &Node_GreaterThan{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x702750DD62D3BFA4:
		return &Node_GetAllResourcesInWallet{}, nil
	case 0x55C6E968B78EE3A4:
		return &Node_GetAngleToTarget{}, nil
	case 0x6159ABFC63F817A9:
		n := &Node_Multiply{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xDFE3C7BB430545A7:
		return &Node_GetHitPoints{}, nil
	case 0x54DBD189813355A6:
		return &Node_IsInteractCandidateSetActive{}, nil
	case 0x33475C6802A3CCA6:
		return &Node_DebugLams{}, nil
	case 0x505178C3079000A7:
		return &Node_AsDecision{}, nil
	case 0xBF77EC68816803A7:
		return &Node_GetCurrentOptionEventForSetting{}, nil
	case 0x1615318E67B54DA7:
		return &Node_GetCurrentOptionEventModForSetting{}, nil
	case 0xC6AFEDA51C3CB0A7:
		return &Node_GetOptionEventForSetting{}, nil
	case 0xC369570D1C63E2A7:
		n := &Node_String{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		strLength := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &strLength)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x28), 0)
		if err != nil {
			return nil, err
		}
		strData := make([]byte, strLength)
		_, err = file.Read(strData)
		if err != nil {
			return nil, err
		}
		n.str = string(strData)
		return n, nil
	case 0x86C0E290F1AD4DA8:
		return &Node_InteractZone{}, nil
	case 0x5FBABEBFC1A2E2AC:
		return &Node_HasCombatSetFlag{}, nil
	case 0xDFA4510CDE5355A9:
		return &Node_HasMeter{}, nil
	case 0x6BB256AC5B70C9AA:
		return &Node_AsLams{}, nil
	case 0xE31CAACDEE14DBAB:
		return &Node_GetObjectWithUniqueFlag{}, nil
	case 0xAAFDC116F94E44AC:
		return &Node_BehaviorTreeSubtree{}, nil
	case 0xE5770991F2D626AD:
		return &Node_Lams{}, nil
	case 0x9B8F2BB4456648AE:
		return &Node_GetProgressionFact{}, nil
	case 0x1AC1E51081DF9FAE:
		return &Node_GetPlayer{}, nil
	case 0x787A99E56BB03CB8:
		return &Node_GetStringHashFromBlackboard{}, nil
	case 0x578B34D8B6E724B4:
		return &Node_AsContextAction{}, nil
	case 0x7E2056F668B1A4B1:
		return &Node_IsMonitorLookingAt{}, nil
	case 0xF72F16207A6215B0:
		return &Node_BitwiseNot{}, nil
	case 0xD226FE538301C5B0:
		return &Node_IsTraversePathEnabled{}, nil
	case 0x81462280D132D6B0:
		return &Node_HasAnySaveGames{}, nil
	case 0xF100FA5C837567B1:
		return &Node_IsJointVisible{}, nil
	case 0xE961E431F4ED29B3:
		return &Node_AsFixedSizeArray{}, nil
	case 0x17E33942C0EE2EB3:
		return &Node_GetSoundEmitterSplineTV{}, nil
	case 0xC41E8A29CCFFB6B3:
		return &Node_Meter{}, nil
	case 0x76D9EE85BEAE10B4:
		return &Node_HasHitOrPartFlags{}, nil
	case 0xB8F0233145DD7CB6:
		n := &Node_Not{}
		n.param, err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xFC9C3531599C31B4:
		return &Node_IsOnGround{}, nil
	case 0xD0CBBCA6788843B4:
		return &Node_GetRecipesInWalletWithFlags{}, nil
	case 0xAD349BD0B7555FB5:
		return &Node_AsBehaviorTreeRoot{}, nil
	case 0x93E78A7B43A101B6:
		return &Node_GetArrayLength{}, nil
	case 0x4032A55A4D21EEB6:
		return &Node_DoesCreatureHaveLookAtTarget{}, nil
	case 0x92A05E33F2D3B0B7:
		return &Node_AsHaptic{}, nil
	case 0x2E8918D499F5B3B7:
		return &Node_IsSoundPlaying{}, nil
	case 0xAB65E4688479DEB7:
		n := &Node_GetPosition{}
		n.param, err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x0226303867D3F9BB:
		return &Node_GetMeterMin{}, nil
	case 0xEE4D32DA6322E8B9:
		return &Node_GetGlobalVariable{}, nil
	case 0xEEECB26DF1A255B8:
		return &Node_IsGoldBuild{}, nil
	case 0x6456B4DFFFBF64B8:
		return &Node_AsStringHash{}, nil
	case 0x3342095F3FD484B8:
		return &Node_GetDistanceToTargetCapsule{}, nil
	case 0x770741C5BE8201B9:
		return &Node_Animation{}, nil
	case 0x93BEC81F6320F9BA:
		return &Node_AsArrow{}, nil
	case 0x2F87F985625358BB:
		return &Node_IsEncounterRunning{}, nil
	case 0x87C06047A5C4AFBB:
		return &Node_AsCameraRecenter{}, nil
	case 0x4911B8AB8B0DECBB:
		return &Node_AreCreaturesOnSameTeam{}, nil
	case 0x63C122CC9AF800C0:
		return &Node_GetResourceAmountInWallet{}, nil
	case 0x053FEDFE0AE084BC:
		return &Node_AreTraverseGraphsEnabledForObjects{}, nil
	case 0xF4B78004ECD8E3BC:
		n := &Node_HasFlag{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xA76B7D032A8E32BE:
		return &Node_BreakVector4{}, nil
	case 0x7880289BA5E9C8BF:
		return &Node_Vector4{}, nil
	case 0xD189CE56A4830EC0:
		return &Node_GetCallAndResponseObjectsInRadius{}, nil
	case 0xF432A630B5C843C0:
		return &Node_IsOnTraverseLink{}, nil
	case 0x46659A3A806247C0:
		return &Node_MapArea{}, nil
	case 0x9B0CDE5C5E5799E3:
		return &Node_GetMapAreasInRegion{}, nil
	case 0x39635B54512FEED1:
		return &Node_EquipmentGeneratorHasFlags{}, nil
	case 0x495E1EB3063A82C9:
		return &Node_GetPickupInSlot{}, nil
	case 0x95FCA6571219A9C5:
		return &Node_RotationFromEuler{}, nil
	case 0x8BB6CF4ADF19FFC3:
		return &Node_AsTraverseLink{}, nil
	case 0xD17C2FA6806A87C1:
		return &Node_GetSelectedCreature{}, nil
	case 0x4C55E233FC9124C2:
		return &Node_PlayFX{}, nil
	case 0x7B15884DA37D2CC3:
		return &Node_GetCreatureStatsValue{}, nil
	case 0xFD68FEF59C136AC3:
		return &Node_IsDead{}, nil
	case 0x99FF9FA945EE7EC4:
		return &Node_Modulo{}, nil
	case 0xB7CD252DC758D4C4:
		return &Node_IsUnderSystemCon{}, nil
	case 0x76315154735FEEC4:
		return &Node_LootDistributor{}, nil
	case 0x10EA2556B0F501C5:
		return &Node_CreatureSpawnOptions{}, nil
	case 0x672B8FA3E239F8C6:
		return &Node_IsMusicPlayingAnything{}, nil
	case 0x37EC670CE44437C6:
		return &Node_GetPlayersCurrentRealm{}, nil
	case 0x04349256F7BF5CC6:
		return &Node_IsBanterPlayingOnCharacter{}, nil
	case 0x37EEBE2B4FF27FC6:
		n := &Node_Array{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		numElements := uint16(0)
		err = binary.Read(file, binary.LittleEndian, &numElements)
		if err != nil {
			return nil, err
		}
		n.elems = make([]Node, numElements)
		for i := uint32(0); i < uint32(numElements); i++ {
			n.elems[i], err = ReadNode(file, stack, offset+0x28+i*4)
			if err != nil {
				return nil, err
			}
		}
		return n, nil
	case 0x2704A419565AD7C6:
		return &Node_GetMeterMax{}, nil
	case 0xD29B600A6F9187C7:
		return &Node_AreTraverseLinksEnabledForObjects{}, nil
	case 0xC37CA2B248ADC5C7:
		return &Node_CanEquipTokenOfTypeTo{}, nil
	case 0xD4A31B29754CEEC8:
		return &Node_InstancedSoundPortal{}, nil
	case 0x1B3CABBE455981C9:
		return &Node_GetFirstEquippableCharacterSlot{}, nil
	case 0x368558C4ADBC16CE:
		return &Node_GetContextActionTraversalData{}, nil
	case 0x285B0213E0F582CB:
		return &Node_Camera{}, nil
	case 0xFB02F8C5E190EAC9:
		return &Node_IsAvailableForFollow{}, nil
	case 0x69872491E73AF4C9:
		return &Node_GetLootConditionSetByName{}, nil
	case 0xB6A56E5C21D55CCA:
		return &Node_IsCreature{}, nil
	case 0x9C5D660BC9DB7CCB:
		return &Node_GetCreaturesSpeedSettingFor{}, nil
	case 0xBDC90EF314FDA9CB:
		return &Node_MapMarker{}, nil
	case 0xCB3886EF25E415CC:
		return &Node_GetNumAvailableGameObjects{}, nil
	case 0x3F1C46148B738BCC:
		return &Node_IsAvailableForSplines{}, nil
	case 0x249CFF2245FDA1CD:
		return &Node_AsIcon{}, nil
	case 0x3508216B54C220D1:
		return &Node_IsMainQuestInAnotherRealm{}, nil
	case 0xDAAD760A7E4AE3CE:
		return &Node_GetClosestPositionOnNavMesh{}, nil
	case 0x178C304720B937CF:
		return &Node_GetPointOnFX{}, nil
	case 0xD2B6FB0F5611CACF:
		return &Node_Color{}, nil
	case 0x4A9129D8D9082BD0:
		return &Node_AsArrowEmitter{}, nil
	case 0x647E815E224B39D1:
		return &Node_GetEnemyFromEncounter{}, nil
	case 0xB49B90D88BEA54D1:
		return &Node_GetTimeScale{}, nil
	case 0x1E12DB985CACC9D1:
		return &Node_AsBitfield{}, nil
	case 0x4DEB6AFF92B2C4D9:
		return &Node_IsNavObstacleEnabled{}, nil
	case 0x81B24747F7AC40D7:
		return &Node_IsEncounterCompleted{}, nil
	case 0x2F4F2CDF82D589D5:
		return &Node_GetCenterOfScreenEnemy{}, nil
	case 0xAF0FD36574803CD2:
		return &Node_InstancedTraverseLink{}, nil
	case 0xC6B7F5822F0B9BD2:
		return &Node_IsAnyWeatherSystemActive{}, nil
	case 0x97560C0AF399B4D3:
		return &Node_GetParticleMonsterAliveTotemCount{}, nil
	case 0xF40E9C2B8C4429D4:
		return &Node_LootResultHasEntry{}, nil
	case 0xB01451D181A4A6D5:
		return &Node_TraversePath{}, nil
	case 0x6F75AB11FCBBB7D5:
		return &Node_AsScreenShake{}, nil
	case 0xEB250729C71019D6:
		return &Node_IsFocalZoneLockInEnabled{}, nil
	case 0xB79D93C8B88E19D6:
		return &Node_SimpleStateMachine_GetState{}, nil
	case 0x03F9F3FDCA190FD8:
		n := &Node_LessThan{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x046CD45CAE0644D7:
		return &Node_GetPlayingBanter{}, nil
	case 0xA93B70BB950485D7:
		return &Node_IsFightKnowledgeEnabled{}, nil
	case 0x96D0FA1D7BAFBED7:
		return &Node_MapRegionHasAllFlags{}, nil
	case 0x64BA34F93ACFFDD7:
		return &Node_IsContextActionTraversalBehaviorInitialized{}, nil
	case 0xA4CF9664A7D11ED8:
		n := &Node_WwiseEvent{}
		_, err = file.Seek(int64(offset+0x18), 0)
		if err != nil {
			return nil, err
		}
		strLength := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &strLength)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x30), 0)
		if err != nil {
			return nil, err
		}
		strData := make([]byte, strLength)
		_, err = file.Read(strData)
		if err != nil {
			return nil, err
		}
		n.name = string(strData)
		return n, nil
	case 0x7D379EE135DF9AD9:
		return &Node_IsSettingEnabled{}, nil
	case 0xE33553F71FA59FD9:
		return &Node_IsInNavigationMove{}, nil
	case 0x9FF0429CF84CBDD9:
		return &Node_IsAvailableForCombat{}, nil
	case 0x907BA069DD6033E1:
		return &Node_IsControlDisabled{}, nil
	case 0x5B9A8608258C96DC:
		return &Node_GetLastAttacker{}, nil
	case 0x11641DAC004780DA:
		return &Node_CanBanterPlay{}, nil
	case 0xFDE9D25003E0E3DA:
		return &Node_IsQuestInState{}, nil
	case 0xB17A3F4EF1DF05DB:
		n := &Node_TweakConstant{}
		_, err = file.Seek(int64(offset+0x30), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.val)
		_, err = file.Seek(int64(offset+0x48), 0)
		if err != nil {
			return nil, err
		}
		strLength := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &strLength)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x58), 0)
		if err != nil {
			return nil, err
		}
		strData := make([]byte, strLength)
		_, err = file.Read(strData)
		if err != nil {
			return nil, err
		}
		n.name = string(strData)
		return n, nil
	case 0x8A72A8C41BB8A1DB:
		return &Node_IsMainQuestPathfindingNonContiguous{}, nil
	case 0xB21F68D00245B7DC:
		n := &Node_GetTweakIntegerConstant{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x89F027CD83D021DE:
		return &Node_AsWeapon{}, nil
	case 0x6BB77C3F1D0313E0:
		return &Node_IsAvailableForBanter{}, nil
	case 0x237B2246A8DAFDE0:
		return &Node_MapAreaHasAllFlags{}, nil
	case 0xF9BFE3551C6692E2:
		return &Node_GetFixedSizeArrayCapacity{}, nil
	case 0x74C2F79FF80CA1E1:
		return &Node_InteractZoneEnabled{}, nil
	case 0x8DA19E6D6441D9E1:
		return &Node_IsWithinAngle{}, nil
	case 0x1367F30432F642E2:
		return &Node_PickupSlot{}, nil
	case 0xBD65B3FCDD2C53E2:
		return &Node_IsOnActiveTraversePath{}, nil
	case 0xFFB95FFE2187D9E2:
		return &Node_IsPaused{}, nil
	case 0x736B88FCE725EEE2:
		n := &Node_IsPickupAcquired{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		n.params[2], err = ReadNode(file, stack, offset+0x18)
		if err != nil {
			return nil, err
		}
		numParams := uint32(0)
		_, err = file.Seek(int64(offset+0x20), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &numParams)
		if err != nil {
			return nil, err
		}
		n.params2 = make([]Node, numParams)
		for i := uint32(0); i < numParams; i++ {
			n.params2[i], err = ReadNode(file, stack, offset+0x40+i*4)
			if err != nil {
				return nil, err
			}
		}
		return n, nil
	case 0x8C39C23A821601E3:
		return &Node_AsWad{}, nil
	case 0x84F7EDDA93CF63F2:
		return &Node_AnimJoint{}, nil
	case 0x57BDA2D4B0C09DE9:
		return &Node_GetActiveWeatherSystems{}, nil
	case 0x5538B58B256474E7:
		return &Node_GetQuest{}, nil
	case 0x4AC136F672A591E6:
		return &Node_GetSpawnedObjectOwningObject{}, nil
	case 0xD0321ADD1D0AFAE4:
		return &Node_GetCurrentAnimation{}, nil
	case 0xA3732CA01D2403E5:
		return &Node_GetCombatStatus{}, nil
	case 0xC4B873BA32340CE5:
		return &Node_Tan{}, nil
	case 0x8F5C142799B7B2E5:
		return &Node_GetWolfSledIsDriftAllowed{}, nil
	case 0x33676F753F6CC9E6:
		return &Node_GetFloatFromBlackboard{}, nil
	case 0xAA3CE495FE1BD8E6:
		return &Node_GetResourcesInWalletWithFlags{}, nil
	case 0xDC600FCEB385E1E6:
		return &Node_GetCreatureAttributeValue{}, nil
	case 0xD88F5C7C1F760EE7:
		return &Node_Banter{}, nil
	case 0x4DDDF19C603A6AE8:
		return &Node_GetNavCurveInformation{}, nil
	case 0xE6B13578701776E7:
		return &Node_FindRandomPositionOnNavMeshInRectangle{}, nil
	case 0x0E34D69ACD4DA5E7:
		return &Node_GetEquipmentInEquipmentSlot{}, nil
	case 0xCCF200B1F6E5F9E7:
		return &Node_GetAnimationPlayRate{}, nil
	case 0xC67859E53DA50CE8:
		return &Node_GetValhallaComplete{}, nil
	case 0x36792104A9317EE8:
		return &Node_GetMiniGameplaySkipped{}, nil
	case 0x0B2CC0C2456EB9E8:
		return &Node_IsWaveRunning{}, nil
	case 0xAC14AC78C1C85EE9:
		return &Node_AsBeamAttack{}, nil
	case 0x020912693E7C73E9:
		return &Node_IsAvailableForSync{}, nil
	case 0xBECDF5899D7A22EF:
		return &Node_GetNumPointsOnFX{}, nil
	case 0x1FA2A974C1A752EB:
		return &Node_GetPickupSlotName{}, nil
	case 0x9356985B357FF9E9:
		return &Node_IsParticleEmitterEnabled{}, nil
	case 0x8708331F910095EA:
		return &Node_GetValuesFromStageData{}, nil
	case 0x4DF5397C12DCA3EA:
		return &Node_Arrow{}, nil
	case 0xA0CED481D2A9D6EA:
		return &Node_AsVector4{}, nil
	case 0xF6CF4842280431EC:
		return &Node_GetFlawlessZeusAvailable{}, nil
	case 0x9AB2DAEC074AD8EC:
		return &Node_GetNearestSoundPortal{}, nil
	case 0xB47E0FF794B39EED:
		return &Node_IsMusicPlaying{}, nil
	case 0x0D4826EF61A9B4EE:
		return &Node_BehaviorTreeRoot{}, nil
	case 0x5506D83EAD1E4DF0:
		return &Node_GetIntValueFromElement{}, nil
	case 0x5DA75980D97930EF:
		return &Node_AsWwiseRTPC{}, nil
	case 0x02C33FBD7D6A63EF:
		return &Node_IsVisualScriptEnabled{}, nil
	case 0x3874D2177E6CE0EF:
		return &Node_GetGameObjectValueFromElement{}, nil
	case 0x5EB00DEA065407F0:
		return &Node_GetGameTime{}, nil
	case 0x303ABE3F2C0074F0:
		return &Node_AsBehaviorTreeSubtree{}, nil
	case 0x60602ECBB26BC0F0:
		return &Node_GetParent{}, nil
	case 0x9BB436CDE2E232F2:
		return &Node_GetCreatureContext{}, nil
	case 0x27CBE85B9C93DEFA:
		n := &Node_Int32{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.val)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xC433C5E8EBB7CEF7:
		return &Node_GetGameObjectWad{}, nil
	case 0x161F6BECFDA7EFF5:
		return &Node_WwiseSwitchGroup{}, nil
	case 0x46395421ED4DEAF3:
		return &Node_AsTraversePath{}, nil
	case 0x89E8E463C819CDF4:
		return &Node_SoundProximityTrigger{}, nil
	case 0xD59DCCEF0AE031F5:
		return &Node_ChooseByValue{}, nil
	case 0x3AA614F52D3A53F5:
		return &Node_GetNumBreakableStages{}, nil
	case 0x0899E3885F1A5EF6:
		return &Node_GetSpawnElementUID{}, nil
	case 0x970195BB56316BF6:
		return &Node_Wad{}, nil
	case 0x1CD81E378151C6F7:
		return &Node_ConcatenateAndHashStrings{}, nil
	case 0x202C4ADD3B39CBF7:
		return &Node_GetBitFromGlobalVariable{}, nil
	case 0x0F94C4CD796462F9:
		return &Node_AsResource{}, nil
	case 0x0D3A2D374D0FF6F7:
		return &Node_Creature{}, nil
	case 0xC16C77586BDF35F8:
		return &Node_GetOptionEventModForSetting{}, nil
	case 0x7059642CA7FB9FF8:
		return &Node_AreParticlesEnabled{}, nil
	case 0x0E0B6301D1621BF9:
		return &Node_Decision{}, nil
	case 0x2A48C16366F9DEF9:
		return &Node_IsQuestTracked{}, nil
	case 0x1D75EDF5FBE304FA:
		return &Node_GetSpawnedObjectHierarchyRoot{}, nil
	case 0x9638C8EE21BE49FA:
		return &Node_IsSummaryCategoryInMapArea{}, nil
	case 0xD8264A83A27FD2FA:
		return &Node_VectorNormalize{}, nil
	case 0x5CFE5D60A01723FE:
		return &Node_AreInteractZonesEnabledForObjects{}, nil
	case 0x13C70F875EAA77FC:
		n := &Node_UInt64{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.val)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xEBFF6DA57D36E0FA:
		return &Node_Cos{}, nil
	case 0x6929CF4286D1FCFA:
		return &Node_GetMeterValue{}, nil
	case 0x72DF2EFD1FD1D2FB:
		return &Node_AsLootConditionSet{}, nil
	case 0x63A1BB2AE5B8DFFB:
		return &Node_Concussion{}, nil
	case 0x630D325A8182DFFC:
		return &Node_GetQuestPrimaryCompletionFact{}, nil
	case 0x1057AC44F526E0FC:
		n := &Node_Equals{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x14)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xE27D8DA7F380F0FC:
		return &Node_Quest{}, nil
	case 0x07F8BD2057B179FD:
		return &Node_IsSkillAcquired{}, nil
	case 0x8EC3504D30015AFF:
		return &Node_CheckProgressionFactGreaterEquals{}, nil
	case 0x6FD0959F311324FE:
		return &Node_GetNameHash{}, nil
	case 0xDF83E783F97B52FE:
		return &Node_RotateVector{}, nil
	case 0x115D2C6DFA0D6AFE:
		return &Node_GetRotation{}, nil
	case 0x157ADF5C41BB4BFF:
		return &Node_AsMapMarker{}, nil
	case 0xB843D9F7DFB461FF:
		return &Node_IsInAir{}, nil
	case 0xE63B6D13A85A6AFF:
		return &Node_ResourceHasFlag{}, nil
	case 0x5540C13E096EE5FF:
		return &Node_AsInteractZone{}, nil
	case 0xDA5B10BB7B14EF07:
		return &Node_ControlClothSim{}, nil
	case 0x7727A8CE67F30804:
		return &Node_DrawPoint{}, nil
	case 0x3E251066BD213901:
		return &Node_MarkResurrectionStoneUsed{}, nil
	case 0x71DDE16CD8DF3600:
		return &Node_ClearWeatherCategory{}, nil
	case 0xC3534134634A4F00:
		return &Node_SetAnimationPlayRate{}, nil
	case 0xCABC7C93B62E5200:
		return &Node_SetDisableAllInteracts{}, nil
	case 0x78E02FAFE4467D00:
		return &Node_AddExcludeTraverseLinkFilterToCreature{}, nil
	case 0xC96987A11B93A400:
		return &Node_PlayCombatConcussion{}, nil
	case 0xDE783CC0462E3901:
		return &Node_SetHapticInstanceParameter{}, nil
	case 0x62B858BE1916AC01:
		return &Node_RequestCinematicModeEnabled{}, nil
	case 0xBEE4A5313024E102:
		return &Node_Sluice_StartWater{}, nil
	case 0x4F74958321C7F602:
		return &Node_StopRotatingObjectTowardsTarget{}, nil
	case 0x7857067582334203:
		return &Node_RegisterOnEncounterFinished{}, nil
	case 0x10C4F2AD364BB005:
		return &Node_SetTraversePathFlagEnabled{}, nil
	case 0x387D17C39FDDA304:
		return &Node_SetSplineLeashing{}, nil
	case 0xE49568386456F104:
		return &Node_DrawLine2D{}, nil
	case 0x92ED96A77591FD04:
		return &Node_SetWolfSledAutoRunSplineAngle{}, nil
	case 0xB4C390B0403B4705:
		return &Node_Crank_Enable{}, nil
	case 0x0CBDFB84C3719505:
		return &Node_CacheValuesOnStack{}, nil
	case 0x73AE4293D8116406:
		return &Node_SetAggressivePriorityOverride{}, nil
	case 0x8F07AB3AFF3F6806:
		return &Node_Crank_UpdateDrivenObjectAnim{}, nil
	case 0xC52D14E2556EA606:
		return &Node_LootPot_Reroll{}, nil
	case 0x288B37F314A28C07:
		return &Node_DisableBoatForceTurnAroundControlMode{}, nil
	case 0xB9CBE07D9FA2BA07:
		return &Node_AddRecipeToWallet{}, nil
	case 0x80CE4CB46A3DCB0D:
		return &Node_RemoveTraverseLinkFilter{}, nil
	case 0x4929DDCD848B420B:
		return &Node_StopAllSounds{}, nil
	case 0x2989409736FB5709:
		return &Node_SimpleStateMachine_ClearStateMac{}, nil
	case 0x55FAF3E7AF7AFC09:
		return &Node_SendTelemetryEvent{}, nil
	case 0x1C7C9A2235650D0A:
		return &Node_SetSplineFollowDistances{}, nil
	case 0xA8228CF5B40F5B0A:
		return &Node_ClearFinishedEncounterData{}, nil
	case 0xA042A1D17C3D910A:
		return &Node_ClearPlayerNavCurve{}, nil
	case 0x368617853E18670B:
		return &Node_SetNavObstacleEnabled{}, nil
	case 0x122C2762943FEF0B:
		return &Node_SetPreventSoftSave{}, nil
	case 0x4CE027774001FF0B:
		return &Node_FindInteractZoneWithName{}, nil
	case 0x8FC95FA334416B0D:
		return &Node_SubtractFromGlobalVariable{}, nil
	case 0xBBF4D2C15527900D:
		return &Node_CheckCineModeIntegrity{}, nil
	case 0xE718AEFB3E001A11:
		return &Node_InterruptBanterOnCharacter{}, nil
	case 0xF819B981F8233A0E:
		return &Node_SetContextActionApproachData{}, nil
	case 0x4A87EDA43D69F50E:
		return &Node_ForceCompanionTraverseLink{}, nil
	case 0xF303A75F7ECB150F:
		return &Node_SendImmediateEvent{}, nil
	case 0xA4EA6BD9694FB30F:
		return &Node_SpawnNPC{}, nil
	case 0x268F34910A62FC0F:
		return &Node_SetBehaviorTreeUpdatePriority{}, nil
	case 0xB26F48EDDBC01A11:
		return &Node_CreateVFSExec{}, nil
	case 0xA4386EE237F14211:
		return &Node_RemoveRecipeFromWallet{}, nil
	case 0x9A97E399977CAB11:
		return &Node_EndCreatureInteract{}, nil
	case 0xCF0B837AF11DE711:
		return &Node_SendTelemetryEventDLC{}, nil
	case 0x9D267AB056B1B318:
		return &Node_SetNavCurveReversal{}, nil
	case 0x75862A3D2C034015:
		return &Node_SetJointHighlightCategory{}, nil
	case 0x5DFCC8477BED2713:
		return &Node_AlertWave{}, nil
	case 0xB62F4B594F105212:
		return &Node_DrawArc{}, nil
	case 0xBBA840868A768D12:
		return &Node_DismissEffectsSubtitle{}, nil
	case 0xB0C45756AD00F212:
		return &Node_ModifyFocalZoneCameraAngleDeactivationThreshold{}, nil
	case 0xF769E83282D0F712:
		return &Node_SetChainVisible{}, nil
	case 0xF45FA65471CC1D13:
		return &Node_AddFlag{}, nil
	case 0xDBECE7DA9D4F5C13:
		return &Node_DistributeLootResult{}, nil
	case 0x5139FD4DB7F0B013:
		return &Node_ClearFocalZones{}, nil
	case 0x65F39DF8E8A5FC13:
		return &Node_SetPhysicsEnabled{}, nil
	case 0xFF5247A36B39C014:
		return &Node_UnmapTouchpadSwipe{}, nil
	case 0x1A9B2EC813932A15:
		return &Node_StartYakEnter{}, nil
	case 0x6738DA20AFE3C716:
		return &Node_GetClosestElementToPlayer{}, nil
	case 0x4ABD4E55056A8015:
		return &Node_TransferWalletContents{}, nil
	case 0x569008FA65B19E15:
		return &Node_SetVFSBool{}, nil
	case 0x538FCED6126B4816:
		return &Node_Switch{}, nil
	case 0x5C715BBAFD107416:
		return &Node_DebugHighlightObject{}, nil
	case 0xCC66989424F8A216:
		return &Node_SendBehaviorTreeEvent{}, nil
	case 0x4CF5F1E0086D5617:
		return &Node_RemoveByValue{}, nil
	case 0x48A65FA02D407E17:
		return &Node_SetEnableCounterWind{}, nil
	case 0xEAEBD054CA33F617:
		return &Node_GrantLootCondition{}, nil
	case 0x5E4E9F67AACB1C18:
		return &Node_SetCombatStatus{}, nil
	case 0xF9230FF35FA39B18:
		return &Node_ClearAllAnimationCallbacks{}, nil
	case 0xF0BE424E415C451C:
		return &Node_MarkCurrentWeaponModeForPlayer{}, nil
	case 0xCB1C0A56AB52201A:
		return &Node_SetBehaviorsEnabled{}, nil
	case 0x043473E2CF40F118:
		return &Node_ResetListenerPosition{}, nil
	case 0x67546100F1C53719:
		return &Node_SetAOOAssignment{}, nil
	case 0x9028FA4F28D7DE19:
		return &Node_CreateEquipment{}, nil
	case 0x97171F3FCF41021A:
		return &Node_MapTouchpadSwipe{}, nil
	case 0xE55AE1D8A55B051A:
		return &Node_ActivateWeatherCategory{}, nil
	case 0x4D99CD1C4078861A:
		return &Node_TriggerBlenderEffect{}, nil
	case 0x884336645708881A:
		return &Node_SetCompassTargetPositionLerpSpeed{}, nil
	case 0x9593DC7E02A2001B:
		return &Node_SetUDSActivityAvailable{}, nil
	case 0x55CE4F200039191B:
		return &Node_SetJointVisible{}, nil
	case 0x7A343491407F3A1B:
		return &Node_Crank_Unlock{}, nil
	case 0x2DE0C45B3EE6611E:
		return &Node_Sluice_WaterPourStopped{}, nil
	case 0x4952916514C9881C:
		return &Node_SetContextActionMentalStateTag{}, nil
	case 0x8500B62F5ECA041D:
		return &Node_SetCombatIndicator{}, nil
	case 0xE4C12FBA9F8A611D:
		return &Node_ForceCompanionTraversePath{}, nil
	case 0xEA5893450B92B81D:
		return &Node_SetOverrideSpeedControlData{}, nil
	case 0x1E0380F5E66D041E:
		return &Node_ActivateBoatNavCurveFunctionalityFunneling{}, nil
	case 0x25DCB8EE9DF86C1E:
		n := &Node_ClearVariable{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		variableOffset := int64(0)
		err = binary.Read(file, binary.LittleEndian, &variableOffset)
		if err != nil {
			return nil, err
		}
		if variableOffset>>0x3e == 1 {
			variableOffset = (variableOffset << 1) / 2
		}
		_, err = file.Seek(int64(offset+0x10)+variableOffset, 0)
		if err != nil {
			return nil, err
		}
		strLength := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &strLength)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x48)+variableOffset, 0)
		if err != nil {
			return nil, err
		}
		strData := make([]byte, strLength)
		_, err = file.Read(strData)
		if err != nil {
			return nil, err
		}
		n.name = string(strData)
		n.next, err = ReadNode(file, stack, offset+0x1a)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xCE6E69C96A68A01E:
		return &Node_ToggleBitOnVariable{}, nil
	case 0x4C78F7F902B0F21E:
		return &Node_StartCheckpointedMusic{}, nil
	case 0xD26E0E945DDF3C1F:
		return &Node_ClearPreventSoftSave{}, nil
	case 0x7EDE4757E12E332E:
		return &Node_SetNavAssistCompassMarker{}, nil
	case 0x6A2DC71F4AA63E28:
		return &Node_ClearAllDisabledControls{}, nil
	case 0xF344DF491B074225:
		return &Node_SetMusicCaption{}, nil
	case 0x33B54CB0351C8622:
		if seen {
			return &Node_Variable{str: "timerId"}, nil
		}
		n := &Node_StartTimer{}
		n.params[0], err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x1e)
		if err != nil {
			return nil, err
		}
		n.next, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xE733EC043D008A1F:
		return &Node_SetSplineAvoidDistances{}, nil
	case 0xA08EAD9CE2E4E31F:
		return &Node_IncrementMeterValue{}, nil
	case 0x7DE0ABAD54201420:
		return &Node_OverridePowerLevelForEncounterElement{}, nil
	case 0x684437F2DE3E3220:
		return &Node_ClearWeather{}, nil
	case 0x735C2E7B7435F221:
		return &Node_StartMusicSting{}, nil
	case 0x92410A8B10829B22:
		return &Node_GrantLootConditionSet{}, nil
	case 0x2F77F227C0A8B622:
		return &Node_SetInteractZoneLocked{}, nil
	case 0x3B35636149B0DE22:
		return &Node_CreatePendulum{}, nil
	case 0xE2A7F695910FE922:
		return &Node_SetBoatIntoTriggeredLaunchState{}, nil
	case 0x55E3F9F4AF766823:
		return &Node_LootPot_SetLootType{}, nil
	case 0xD11C88B1D544B926:
		return &Node_StopRestartableSoundLoop{}, nil
	case 0xCF00A48430315825:
		return &Node_SetOverridePositionOnDock{}, nil
	case 0x0B01ABD4C7AC5F25:
		return &Node_SetAnimDriverValue{}, nil
	case 0xF9B23D1E7CA38525:
		return &Node_ClearForceLookAt{}, nil
	case 0x6B15C03CC49D0A26:
		return &Node_SetDynamicNavmeshEnabled{}, nil
	case 0x7E1BFAB8D15D2A26:
		return &Node_SetWolfSledAutoRunAcceleration{}, nil
	case 0x2A2A172FBBA43327:
		return &Node_SetAIUseTraverseLinksWithoutPath{}, nil
	case 0x308C019DBE0F4C27:
		return &Node_StartCreatureInteract{}, nil
	case 0x2447D8E93F628E27:
		return &Node_SetContextActionFactionTag{}, nil
	case 0xBA9B425DF41E2628:
		return &Node_AcquireSkill{}, nil
	case 0xC0D85C334AB42A28:
		return &Node_RemoveCreatureAvailability{}, nil
	case 0x9744665A8284182C:
		return &Node_SetTraversePathsEnabled{}, nil
	case 0x039F1CF39FA7312A:
		return &Node_RecenterToFacing{}, nil
	case 0xCC0FB878A20C5B28:
		n := &Node_StopTimer{}
		n.param, err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xEB5418804497DB28:
		return &Node_ResetCompassPathfinding{}, nil
	case 0xF423BAFE6A03F528:
		return &Node_ActivatePickup{}, nil
	case 0xE656CE862BC50029:
		return &Node_KillPlayer{}, nil
	case 0xDFC961A37572E529:
		return &Node_EndPaperboatSync{}, nil
	case 0x0D6EBBAC1B3E322A:
		return &Node_WolfSledSplineSectionEnter{}, nil
	case 0x4C371668E474AB2A:
		return &Node_SetTraverseLinkOverridePrompt{}, nil
	case 0x029B09FE5129F92A:
		return &Node_RevertAggressivePriorityOverride{}, nil
	case 0x1C386E87024A052B:
		return &Node_ClearVisualVariantOnDock{}, nil
	case 0xF5F6FCD63EBD6E2B:
		return &Node_PrepareDynamicSyncedSequence{}, nil
	case 0xFBED9C71BD25382D:
		return &Node_OverrideCreatureLootDistributor{}, nil
	case 0x018577AA8FB94C2C:
		n := &Node_Sequence{}
		_, err := file.Seek(int64(offset+0x18), 0)
		if err != nil {
			return nil, err
		}
		numElements := uint16(0)
		err = binary.Read(file, binary.LittleEndian, &numElements)
		if err != nil {
			return nil, err
		}
		n.nodes = make([]Node, numElements)
		for i := uint32(0); i < uint32(numElements); i++ {
			n.nodes[i], err = ReadNode(file, stack, offset+0x28+i*4)
			if err != nil {
				return nil, err
			}
		}
		return n, nil
	case 0x2EAA9A11A072652C:
		return &Node_SetBreakableEnabled{}, nil
	case 0x5C09202FCF23712C:
		return &Node_DrawCircle{}, nil
	case 0xFF3593492281F42C:
		return &Node_RequestDespawnDynamicCharacter{}, nil
	case 0xCD34986ACC68252D:
		return &Node_AddCharacterFX{}, nil
	case 0x8EE2A416D4FB762D:
		return &Node_ClearOverrideSpeedControlData{}, nil
	case 0x9BB4603768D6922D:
		return &Node_SetWaitGateEnabledState{}, nil
	case 0x42C0CCCFEB51BF2D:
		return &Node_SimpleStateMachine_SwitchState{}, nil
	case 0x658FCCA94838FB2D:
		return &Node_SetContextActionTraversalCompanionLead{}, nil
	case 0x6C8A7E2A067B9838:
		return &Node_SetClonePositionRotation{}, nil
	case 0x51E44AA5E9634133:
		return &Node_SetPaperBoatOffsetSpringValues{}, nil
	case 0x86C3A625D03C7D2F:
		return &Node_EndSyncedSequenceByActor{}, nil
	case 0x6E6E6D6D6D1A4A2E:
		return &Node_ProfileBegin{}, nil
	case 0x93544193F56E882E:
		return &Node_SwitchOnPinType{}, nil
	case 0x72F45AB36D00D82E:
		return &Node_StopAndSettlePendulum{}, nil
	case 0x52971ED46F9B5B2F:
		return &Node_SetWolfSledAutoRunStrafeMultiplier{}, nil
	case 0xB2C4A0B134191B30:
		return &Node_DrawZone{}, nil
	case 0xAF485758559C5C30:
		return &Node_DeactivateBoatNavCurveFunctionality{}, nil
	case 0x2F863749050CCA30:
		return &Node_SetSoundProximityTriggerEnabled{}, nil
	case 0x1102C7F05B954632:
		return &Node_DebugSpawner{}, nil
	case 0x3E4593F64110FA32:
		return &Node_SetContextActionTraversalSegmentDiffThreshold{}, nil
	case 0x156F11A23DD01236:
		return &Node_NotifySluiceIceChanged{}, nil
	case 0xE93B1DEA0E7BF933:
		return &Node_HideUIMessage{}, nil
	case 0xD10679CE47A61934:
		return &Node_AddAreaOfOperation{}, nil
	case 0x76C490AC76EC8934:
		return &Node_SendContainerRewardAcquiredTelemetry{}, nil
	case 0x0A90C8256932AF35:
		return &Node_ShowUIMessageWithTokensBase{}, nil
	case 0xDDED51C91AB6C735:
		return &Node_SetInfiniteSpawning{}, nil
	case 0x0CF8E56F4BC06436:
		return &Node_FixBreakable{}, nil
	case 0xF13CFC8409357536:
		return &Node_SetMusicWwiseSwitch{}, nil
	case 0x717FC3314922FE36:
		return &Node_OverridePowerLevelForEncounter{}, nil
	case 0x6FA5583940A59937:
		return &Node_CreateVFSInt{}, nil
	case 0x06BAE7F60494E83B:
		return &Node_WildlifeDestroy{}, nil
	case 0xC29F860E66F9813A:
		return &Node_StartCamera{}, nil
	case 0x2A00FD10ED97AB38:
		return &Node_SetPendulumPosition{}, nil
	case 0x161B8E733316BA38:
		return &Node_ClearParticleMonsterCheckpointing{}, nil
	case 0x9F053B4780A20639:
		return &Node_SetEnableAreaOfOperation{}, nil
	case 0x720FFE691266A639:
		return &Node_Sluice_RemoveManualBlocker{}, nil
	case 0xD582C0B1D425D439:
		return &Node_RegisterSpawnedObjectForFrameUpdate{}, nil
	case 0x2EC26095F10CDF3A:
		return &Node_RemovePropToSyncWithCAOwner{}, nil
	case 0xE05AAABFD680183B:
		return &Node_RegisterForDistanceEvent{}, nil
	case 0xE9344B811A781C3B:
		return &Node_AddInteractZoneTag{}, nil
	case 0x1B08D02274C3A23B:
		return &Node_RemoveIncludeTraverseLinkFilterFromCreature{}, nil
	case 0x7AEC7BA79A00CF3B:
		return &Node_SetContextActionInitialized{}, nil
	case 0x6855CD0B8F6E323E:
		return &Node_TeleportWolfSled{}, nil
	case 0x62EC3086F7EE0B3D:
		return &Node_ModifyNavCurveProgressionLimit{}, nil
	case 0x4692F93A886E523D:
		return &Node_SetWolfSledAutoRunSpeedModifier{}, nil
	case 0xED30CA2022C5803D:
		return &Node_GrantCreatureLoot{}, nil
	case 0x09AD7B015DB9D83D:
		return &Node_QuitAndJumpToWad{}, nil
	case 0x3E0CB4D9C9232B3E:
		return &Node_HideAllUIMessages{}, nil
	case 0x12DC70E7837B773E:
		return &Node_AddCreatureToCustomSplineGroup{}, nil
	case 0x465A37C8D78DB63E:
		return &Node_SetContextActionTraversalDistanceThreshold{}, nil
	case 0xFB20B46C7EB91E3F:
		return &Node_SetPendulumRotation{}, nil
	case 0x16630BEEEC20213F:
		return &Node_ClearOverridePositionOnDock{}, nil
	case 0xD53108F82C22A75C:
		return &Node_RemoveContextActionCreatureTag{}, nil
	case 0x219FE8B841DF424D:
		return &Node_RegisterBoatDock{}, nil
	case 0xEF58D3E8FD14CE45:
		return &Node_SetWolfSledHarnessVisibility{}, nil
	case 0xDD6BD2434A0F1742:
		return &Node_CreateVFSBool{}, nil
	case 0x301C88564F331841:
		return &Node_DEBUG_ONLY_SetQuestStateAndBackpropagate{}, nil
	case 0xE31E6CFA256BB13F:
		return &Node_For{}, nil
	case 0x943DEDD1D9244240:
		return &Node_ModifyFocalZoneLockInEnabled{}, nil
	case 0xB584358FD7A14840:
		return &Node_SendEnvironment{}, nil
	case 0xDCAEF4A1AA828940:
		return &Node_ResetEncounter{}, nil
	case 0x3D601DE6260A0F41:
		return &Node_CrankSetup{}, nil
	case 0xCB73C182D3DD7D41:
		return &Node_SetPosition{}, nil
	case 0x635521765585AD41:
		return &Node_Crank_ForcePlayerDetach{}, nil
	case 0x5232D8F045C1BE41:
		return &Node_SetCheckpointedMusic{}, nil
	case 0xB2D67099EB54E341:
		return &Node_AttachZiplineToSpear{}, nil
	case 0xF205B0231C85F041:
		return &Node_SetSplineAttributes{}, nil
	case 0xD961E23297EF5244:
		return &Node_PauseEncounter{}, nil
	case 0x51611A1B530F2A42:
		return &Node_SetEncounterLODRange{}, nil
	case 0x7BEBB046B6249242:
		return &Node_RegisterWaitGateSetup{}, nil
	case 0x0A15143B381CD742:
		return &Node_SummonWolfSled{}, nil
	case 0x1C8E4548473BE942:
		n := &Node_SetPickupStage{}
		n.params[0], err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x1a)
		if err != nil {
			return nil, err
		}
		n.params[2], err = ReadNode(file, stack, offset+0x1e)
		if err != nil {
			return nil, err
		}
		n.params[3], err = ReadNode(file, stack, offset+0x22)
		if err != nil {
			return nil, err
		}
		n.next, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xC4B96DDB9A25B743:
		return &Node_RestoreMarkedWeaponModeForPlayer{}, nil
	case 0xA095CC228C1F7A44:
		return &Node_SpawnVehicle{}, nil
	case 0x1F11942C6A599544:
		return &Node_DestroyBeam{}, nil
	case 0x5939F09851630B45:
		return &Node_SetMaterialEntryValue{}, nil
	case 0x4B734DF45781A045:
		return &Node_AnimationSync{}, nil
	case 0x5D4E56346388B645:
		return &Node_RequestSpawnDynamicCharacter{}, nil
	case 0xF7EAC2A1C1938949:
		return &Node_SetControlEnabledWithType{}, nil
	case 0xDFB194EAB199CD47:
		return &Node_SetControlEnabled{}, nil
	case 0xFE00B48403CEFE45:
		return &Node_PauseCurrentAnimation{}, nil
	case 0xDBB12B2F19DF3946:
		return &Node_GetTriggerIndex{}, nil
	case 0x9340C63E01F6D946:
		return &Node_ClearForcedSpline{}, nil
	case 0x31DE27603A865647:
		return &Node_UpdateGameObjectSpline{}, nil
	case 0xA8117AF7EE3F8347:
		return &Node_SetDynamicCharacterSlot{}, nil
	case 0x938CA81FE7CEE447:
		return &Node_ModifyFocalZonePriority{}, nil
	case 0x2D6DD310504C2448:
		return &Node_SetInteractZoneHintSubPrompt{}, nil
	case 0x587984230F023248:
		n := &Node_StopSound{}
		n.params[0], err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x1a)
		if err != nil {
			return nil, err
		}
		n.params[2], err = ReadNode(file, stack, offset+0x1e)
		if err != nil {
			return nil, err
		}
		n.params[3], err = ReadNode(file, stack, offset+0x22)
		if err != nil {
			return nil, err
		}
		n.params[4], err = ReadNode(file, stack, offset+0x26)
		if err != nil {
			return nil, err
		}
		n.params[5], err = ReadNode(file, stack, offset+0x2a)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xD74EA71FFF1C4249:
		return &Node_RaycastCollision{}, nil
	case 0xC604727C4A8C7049:
		return &Node_ClearOverrideEnterBranchesOnDock{}, nil
	case 0xC5D0358F1CDA234C:
		return &Node_UnpauseEncounter{}, nil
	case 0x5C12FA50FABBF249:
		return &Node_Sluice_StopWater{}, nil
	case 0xA7ACED2793D0244B:
		return &Node_Crank_Lock{}, nil
	case 0x9C4CBC2538FB444B:
		return &Node_EndSyncedSequenceByName{}, nil
	case 0xF0292FE82CF7814B:
		return &Node_EndCinematicSkip{}, nil
	case 0xCFE7E5B82F9BDC4B:
		return &Node_StartHapticVibration{}, nil
	case 0xC53764E25ACD244C:
		return &Node_EnableSetting{}, nil
	case 0x4BA5514C4D86484C:
		return &Node_DiscoverSummaryCategoryInAllRegions{}, nil
	case 0x314915088211AB4C:
		return &Node_AddEmbedWeaponToGameObject{}, nil
	case 0x38179EC010AEE14C:
		return &Node_DisableBoatForceDirectionControlMode{}, nil
	case 0xD0F5BE15EDFF5756:
		return &Node_ShowUIMessageBase{}, nil
	case 0xAE3FDD64913A4F51:
		return &Node_ModifyResourceInWallet{}, nil
	case 0x5B086678509EC34E:
		return &Node_Crank_API{}, nil
	case 0x699B5FEBED41A84D:
		return &Node_DrawSphere{}, nil
	case 0x129D7D8D2E42AF4D:
		return &Node_TransitionBetweenFocalZones{}, nil
	case 0x15B065324E27CE4D:
		return &Node_MarkUsed{}, nil
	case 0xD058FED0CF54E84D:
		return &Node_RemoveFocalZone{}, nil
	case 0x151970BED78DAA4E:
		return &Node_ModifyFocalZoneMovementPriorityOverrideWhenLockedIn{}, nil
	case 0x1D6725E75DE77F50:
		return &Node_SetOverrideVisualVariantOnDock{}, nil
	case 0x3856FB635BA39E50:
		return &Node_EnableLoadGate{}, nil
	case 0x6B29F2BBAD06FF50:
		return &Node_SetCompassGatewayMarkerIsOpen{}, nil
	case 0x28EA6ED49B1B2251:
		return &Node_OverrideCreatureLootConditionSet{}, nil
	case 0x9292A0CB50C04151:
		return &Node_ActivateBoatNavCurveFunctionality{}, nil
	case 0x8948BE000FA91954:
		return &Node_ResetHeightField{}, nil
	case 0x5213DE1DA03AA651:
		return &Node_GetInteractLerpObject{}, nil
	case 0xEFFA2D420310BA51:
		return &Node_PushPendulum{}, nil
	case 0x5995AFF0B0F17252:
		return &Node_UnregisterTraverseGraphEntryEvent{}, nil
	case 0x43E658AE7D8DCB53:
		return &Node_SetWolfSledIsDriftAllowed{}, nil
	case 0xE183E938A952DC53:
		return &Node_RemoveThrowableTarget{}, nil
	case 0xEA58B431F33A4454:
		return &Node_ApplyStatusMeterDamage{}, nil
	case 0x5669F3B305DF4454:
		return &Node_FreezePendulum{}, nil
	case 0x7B2C555E4DA90D55:
		return &Node_DiscoverSingleRegionSummaryCategory{}, nil
	case 0x07CC6B26748F8E55:
		return &Node_SetRouteToOtherRealmsViaFastTravel{}, nil
	case 0x563C3B673122C558:
		return &Node_CreateBeam{}, nil
	case 0x397DB182DDC6AF57:
		return &Node_RemoveAreaOfOperation{}, nil
	case 0xA85EE96210E98156:
		return &Node_BreakBreakable{}, nil
	case 0xD1EC526ADCB5A156:
		return &Node_PlayBanter{}, nil
	case 0x03E167B884B5BC56:
		return &Node_ModifyFocalZoneEnablePitchOnHeadTracking{}, nil
	case 0xD2F9C5FED111E556:
		return &Node_UnregisterMinibossCheckpoint{}, nil
	case 0xC389F20F91B02757:
		return &Node_TriggerPendulumSlowdown{}, nil
	case 0xAB5D4ADF86D41A58:
		return &Node_StartNewDLCTelemetryRun{}, nil
	case 0x4701198492FF3B58:
		return &Node_SetMeterValue{}, nil
	case 0x31D075C18EA34058:
		return &Node_SetManualSaveEnabled{}, nil
	case 0x3DD30475EA985E58:
		return &Node_MarkSkipAIUpdate{}, nil
	case 0x04ECF96E34EF7D58:
		return &Node_SetTraverseLinkPromptOffset{}, nil
	case 0xB576B7E88896745A:
		return &Node_AddCombatCollision{}, nil
	case 0x386125DA5A6F1759:
		return &Node_EquipToEquipmentSlot{}, nil
	case 0xAEA860A240B9AD59:
		return &Node_SetupPaperboatSync{}, nil
	case 0x096D3084EE48D059:
		return &Node_UsePlayerHeightForFloating{}, nil
	case 0x69610626A2D8D259:
		return &Node_SendNamedEventBase{}, nil
	case 0x71E735FBCB73EB59:
		return &Node_CreateEmptyLootResult{}, nil
	case 0xF5B6C5F0DA93025B:
		return &Node_RegisterForBreakableBrokenEvent{}, nil
	case 0xE0A1F327FF70095C:
		return &Node_ActivateContextAction{}, nil
	case 0xBE22F8C8B47E905C:
		return &Node_SetScale{}, nil
	case 0xBED58696F973945C:
		return &Node_Crank_JamCrank{}, nil
	case 0x314C06281E34876D:
		return &Node_SetPreventCinematicSkip{}, nil
	case 0x43AD68C88FB8B663:
		return &Node_SetContextActionMultiCreatureData{}, nil
	case 0xA8CB958D0EECBF5F:
		return &Node_SetGlobalVariable{}, nil
	case 0x3FBE712E3A79DB5D:
		return &Node_DestroyVFSEntry{}, nil
	case 0x73DA311E272BD25C:
		return &Node_SetTraversePathPromptEnabled{}, nil
	case 0x8A1284A602512F5D:
		return &Node_OverrideCreatureCullingFadeDistance{}, nil
	case 0x229A3402AAFA765D:
		return &Node_SetFightKnowledgeEnabled{}, nil
	case 0x6898571C88FCBC5D:
		return &Node_WolfSledForceEnter{}, nil
	case 0x642D85E5B06FDA5D:
		return &Node_ClearWolfSledSpeedControlOverride{}, nil
	case 0x68F3CC91971D845E:
		return &Node_SetWolfSledRopeVisibility{}, nil
	case 0x4B444193D8F1D35E:
		return &Node_SendNamedEvent{}, nil
	case 0xB335C58CB151125F:
		return &Node_PositionObjectsForSyncedSequence{}, nil
	case 0xA4C1A9659B9C285F:
		return &Node_UnequipFromEquipmentSlot{}, nil
	case 0x4416805AC540515F:
		return &Node_DespawnVehicle{}, nil
	case 0x9A667F95ADAC2B62:
		return &Node_SimpleStateMachine_Clea{}, nil
	case 0xF1D8FEFFE1926560:
		return &Node_RegisterTraverseGraphEntryEvent{}, nil
	case 0x4226E02CC656E960:
		return &Node_RemoveChild{}, nil
	case 0x05FC42C1A0EB7261:
		return &Node_DrawDebugTable{}, nil
	case 0xB8240C85F4D68861:
		return &Node_Crank_Disable{}, nil
	case 0x11518B3408BF7862:
		return &Node_ActivateWeather{}, nil
	case 0x5EFA7E8894FF7E62:
		return &Node_AddGameObjectToProfileGroup{}, nil
	case 0x730E89A86343D062:
		return &Node_SetZiplineSlack{}, nil
	case 0xAF807E205C2F1963:
		return &Node_TriggerQuestManualActiveBehaviors{}, nil
	case 0xA7D0B75CCEDF8A63:
		return &Node_SetPaperBoatFadeIn{}, nil
	case 0x3EE00202E4E1BE67:
		n := &Node_Print{}
		n.param, err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		n.next, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x61DBF01B41B1B665:
		return &Node_SetWolfSledAutoRun{}, nil
	case 0xE82FBA575A5EC363:
		return &Node_ResetAccesibilityToggleState{}, nil
	case 0x65374919F2D2E563:
		return &Node_SpawnVehicleAtPosition{}, nil
	case 0x2C9D8FDCEF568164:
		return &Node_OverrideCreatureMotionParams{}, nil
	case 0xED1919F454FEED64:
		return &Node_SetMarkerLocked{}, nil
	case 0x59DA6CA9BDF5A365:
		return &Node_PrestreamRequest{}, nil
	case 0x20EE6BA57EFB1C66:
		return &Node_TriggerEncounterEnd{}, nil
	case 0xD137C56A2AAA4667:
		return &Node_BreakLoop{}, nil
	case 0xC7C4F34846BA6367:
		return &Node_AddSpawnToEncounter{}, nil
	case 0x0773A21B30F76D67:
		return &Node_SetSplineSpeed{}, nil
	case 0xDF05DE4B98AFB167:
		return &Node_AddPointToFX{}, nil
	case 0xE06414F529DCCC6B:
		return &Node_ClearCallAndResponseObject{}, nil
	case 0x061CEFE38BEF3069:
		return &Node_OpenSoundPortal{}, nil
	case 0xF4D93323EBEFCF69:
		return &Node_SetPlayerControlledCreature{}, nil
	case 0x2BE57FD32FABAF6A:
		return &Node_SetWwiseRTPCValue{}, nil
	case 0x2D255CBD71900F6B:
		return &Node_GetBranchStartPositionAndOrientation{}, nil
	case 0xBB23CE23A3C67E6B:
		return &Node_SetWolfSledAutoRunSlopeBoostMultiplier{}, nil
	case 0xA7ADEFCC5C47286C:
		return &Node_SetSoundPortalOpen{}, nil
	case 0xB1C43C68ABD6306C:
		return &Node_StartPaperBoatBounceFadeOut{}, nil
	case 0x2C5404F8CB91096D:
		return &Node_SetContextActionEncounterMarkerIDTag{}, nil
	case 0xD79BB5A9C1F9696D:
		return &Node_DrawRect2D{}, nil
	case 0x66CA00E9A7A90173:
		return &Node_SetInfluenceConeAttributes{}, nil
	case 0xB073BF6B158CBE70:
		return &Node_SetPlayerNavCurve{}, nil
	case 0x36BC76908E0A8F6E:
		return &Node_DestroyEquipment{}, nil
	case 0x9F92B70E4931D86D:
		return &Node_SetFlawlessZeusAvailable{}, nil
	case 0xE6D6F3771F35EC6D:
		return &Node_OverrideLookAtConversationSettings{}, nil
	case 0x3A6A03ADAE14F66D:
		return &Node_SetTextOnTextObject{}, nil
	case 0xD3865F197FEF626E:
		return &Node_SetYakInteractOptions{}, nil
	case 0x3BC222CC06E9696E:
		return &Node_GameMap_SetItemStateV2{}, nil
	case 0x6E464014AC1CED6E:
		return &Node_TriggerCallAndResponse{}, nil
	case 0x59BF1AB246EB3D6F:
		return &Node_CheckFeature{}, nil
	case 0x434314B7011C586F:
		return &Node_CoilChainedObject{}, nil
	case 0x420FEB9D3DBFBB6F:
		return &Node_ModifyFocalZoneFacingAngleDeactivationThreshold{}, nil
	case 0x1B304DDD6C3D9A70:
		return &Node_StartWolfSledExit{}, nil
	case 0x4DDA11FE18701672:
		return &Node_SetSlowdown{}, nil
	case 0xD821C10DA2C13B71:
		return &Node_InterruptAllBanter{}, nil
	case 0x8B1762741EA69B71:
		return &Node_Sluice_AddManualBlocker{}, nil
	case 0x775021CFA11CD471:
		return &Node_StopCamera{}, nil
	case 0x130035D4E3EAE271:
		return &Node_SetWolfSledAutoRunRequiresInput{}, nil
	case 0xF056416533200E72:
		return &Node_RegisterForControlMashEvent{}, nil
	case 0xB30F62EEE87BA072:
		return &Node_UnpauseCurrentAnimation{}, nil
	case 0xD9DC1F71DE85A472:
		return &Node_SetCreatureAvailabilityRequest{}, nil
	case 0x68FDE7C168CDB472:
		return &Node_SetCreatureFocus{}, nil
	case 0x5087AC8E5960BE72:
		return &Node_ShowHiddenWads{}, nil
	case 0x5E189DCCC2890B77:
		return &Node_OverrideCreatureLootCategory{}, nil
	case 0xA0469122B8A38F74:
		return &Node_ClearAllWeaponsCinematicMode{}, nil
	case 0xD586105632799173:
		return &Node_ModifyFocalZoneAttachObject{}, nil
	case 0xC3FDBEC4E5C09173:
		return &Node_SetBitOnGlobalVariable{}, nil
	case 0xAC4762B8E9EC0674:
		return &Node_ClearAndHideAllNavAssistCompassMarkers{}, nil
	case 0xD7E9B8996A9A5D74:
		return &Node_ClearCreatureFocus{}, nil
	case 0xE7D6946E7F566274:
		return &Node_RemoveSecondaryActorFromSequence{}, nil
	case 0xCB414DE50803DD74:
		return &Node_RegisterSpawnedObjectForVariables{}, nil
	case 0xA4333446FD611075:
		return &Node_PrepareSyncedSequence{}, nil
	case 0x7D838CB70CA44875:
		return &Node_SetBoatDockInteractRadius{}, nil
	case 0x52B3BAA2A5300376:
		return &Node_RegisterAsChainedObject{}, nil
	case 0x11A796139E109076:
		return &Node_EnableContextAction{}, nil
	case 0x48A8AF7C45DBD079:
		return &Node_SimpleStateMachine_SwitchStateUsingEnumPin{}, nil
	case 0x740024F1E7B57377:
		if seen {
			return &Node_Variable{str: "itr"}, nil
		}
		n := &Node_ForEach{}
		n.param, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x18), 0)
		if err != nil {
			return nil, err
		}
		numParam2 := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &numParam2)
		if err != nil {
			return nil, err
		}
		if numParam2 > 0 {
			_, err = file.Seek(int64(offset+0x20), 0)
			if err != nil {
				return nil, err
			}
			paramOffset := int64(0)
			err = binary.Read(file, binary.LittleEndian, &paramOffset)
			if err != nil {
				return nil, err
			}
			if paramOffset>>0x3e == 1 {
				paramOffset = (paramOffset << 1) / 2
			}
			n.params2 = make([]Node, numParam2)
			for i := uint32(0); i < numParam2; i++ {
				n.params2[i], err = ReadNode(file, stack, offset+0x20+i*4)
				if err != nil {
					return nil, err
				}
			}
		}
		n.operation, err = ReadNode(file, stack, offset+0x28)
		if err != nil {
			return nil, err
		}
		n.next, err = ReadNode(file, stack, offset+0x40)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x0B0F251FC92A7F79:
		return &Node_SetBanterFact{}, nil
	case 0x6D9F38641D279779:
		return &Node_SetWeaponMode{}, nil
	case 0x346DC94BAFA2BD79:
		return &Node_SetWolfSledAutoRunControlRange{}, nil
	case 0xD5B3A3CAF4ABC179:
		return &Node_TransferLootResultEntry{}, nil
	case 0x778C5529632BD679:
		return &Node_SwitchIntRange{}, nil
	case 0x1E5A282AF598267A:
		return &Node_DropThrowable{}, nil
	case 0xDFFA41428F01EB7A:
		return &Node_RegisterForMeterEvent{}, nil
	case 0x333E378B9262FE7A:
		return &Node_SoundEmitterEnabled{}, nil
	case 0xD91BE5EC6BE840BB:
		return &Node_SetFlawlessAresAvailable{}, nil
	case 0x6E7430D3CF19CC9B:
		return &Node_SetContextActionBasicData{}, nil
	case 0x84C244632D21FF8C:
		return &Node_FakeInteraction{}, nil
	case 0x936167F32AA52D84:
		return &Node_SetSpawnedObjectVariable{}, nil
	case 0xE6D818053070497E:
		return &Node_StartBoatSyncMove{}, nil
	case 0x299BC62F52B8867C:
		return &Node_SpawnWildlife{}, nil
	case 0xC401E5F40ACF787B:
		return &Node_SetPointOnFX{}, nil
	case 0x8811A6E37059027C:
		return &Node_GetNextArrayValue{}, nil
	case 0x0636C4F6B6F80D7C:
		return &Node_ProcessInteractByCandidateSet{}, nil
	case 0xA0A7D75092F2807C:
		return &Node_RegisterForBlackboardTimerExpired{}, nil
	case 0x4AE5F0FD2520867C:
		return &Node_ClearHideHealthBar{}, nil
	case 0x4D3F6C2D2363937C:
		return &Node_SetTraversePathPromptLOSEnabled{}, nil
	case 0xC0078F22522ACD7C:
		return &Node_TriggerCamera{}, nil
	case 0xD7DEAFD6478AD17C:
		return &Node_CacheCallAndResponseObject{}, nil
	case 0xF82A8C3D8E593D7D:
		return &Node_SetPropEnabled{}, nil
	case 0xF17FE57DE7713F7D:
		return &Node_ResetStickDamping{}, nil
	case 0xE32E2634F0F9D081:
		return &Node_SetResourceDebugAddMode{}, nil
	case 0x0B35998ECD02947F:
		return &Node_SetUDSActivityStart{}, nil
	case 0x7716D3FF11C6FC7F:
		return &Node_AddNewLookAtTarget{}, nil
	case 0x1C8ED930A00CC280:
		return &Node_RegisterAsZipline{}, nil
	case 0x02F42A0926C06D81:
		return &Node_SetAccelerationOverride{}, nil
	case 0x5A0CF0D59BD7C081:
		return &Node_CustomSortIterator{}, nil
	case 0x98A120AC62923982:
		return &Node_SetSplineAvoidAvoidancePreferences{}, nil
	case 0xA75D59680BEE4782:
		return &Node_UnregisterForMonitorEvent{}, nil
	case 0x16206A5B4EA99482:
		return &Node_RegisterForMovePlayEvent{}, nil
	case 0x0BC4A5B61373B982:
		return &Node_AddContextActionCreatureTag{}, nil
	case 0x2D54924FDA4B1D84:
		return &Node_ClearPreventCheckpoint{}, nil
	case 0x5D9BB9E76E013A89:
		return &Node_SetValhallaComplete{}, nil
	case 0xD7F610BFD75EDE87:
		return &Node_SimpleCombatSync{}, nil
	case 0xE56A267342096F84:
		n := &Node_SetBlackboardTimer{}
		n.params[0], err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x1c)
		if err != nil {
			return nil, err
		}
		n.params[2], err = ReadNode(file, stack, offset+0x20)
		if err != nil {
			return nil, err
		}
		n.params[3], err = ReadNode(file, stack, offset+0x24)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x1a), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.param)
		if err != nil {
			return nil, err
		}
		n.next, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x6674CE0272981885:
		return &Node_ParticleMonsterCheckpointTotemState{}, nil
	case 0xA0DC62E57BBDA485:
		return &Node_WildlifeFlee{}, nil
	case 0x64D5602381F8D485:
		return &Node_OverrideListenerPosition{}, nil
	case 0x1CB11C7D4FD6A986:
		return &Node_SetRaceBackHomeEnableSledControl{}, nil
	case 0xEB6F214E9473E187:
		return &Node_SetForceOddNumTraverseLinkConnections{}, nil
	case 0xE128B61ECBAE0788:
		return &Node_StartYakExit{}, nil
	case 0x2F8EDA94CF125588:
		return &Node_ForceEngageContextAction{}, nil
	case 0xB0C8F5C04B67C388:
		return &Node_SetCombatTarget{}, nil
	case 0x640F100B5C9F1A89:
		return &Node_SetPaperBoatSpeed{}, nil
	case 0x668E069FF8A0048B:
		return &Node_RegisterTraverseGraphEnteredEvent{}, nil
	case 0x865F09779951F889:
		return &Node_ClearBoatDockInteractRadius{}, nil
	case 0xB4F8CFBF24D9778A:
		return &Node_SetVisibility{}, nil
	case 0xC4A39C82A587808A:
		return &Node_GetAdditionalGravityOnPendulum{}, nil
	case 0x391742579FF1B98A:
		return &Node_SetObjectHighlightCategory{}, nil
	case 0xD2BBE09E1C5B028B:
		return &Node_OverrideCreatureCullingRadius{}, nil
	case 0x03D28F3C9639468B:
		return &Node_SendUIEvent{}, nil
	case 0xFC7C9BF83551D38B:
		return &Node_ToggleBitOnGlobalVariable{}, nil
	case 0xE6C4FED4DD85FA8B:
		return &Node_RemovePositionalFlag{}, nil
	case 0x38C9CDDE5C664B8C:
		return &Node_UntrackAllSideQuests{}, nil
	case 0x704E59DE641A4794:
		return &Node_ClearAllDisabledControlsWithType{}, nil
	case 0x75FAC3DCC24AF38F:
		return &Node_AddThrowableTarget{}, nil
	case 0xC8A4564AF1AA628E:
		return &Node_SetInteractZoneEnabled{}, nil
	case 0xED0F9F98867B0C8D:
		return &Node_PrepareRegisteredSyncedSequence{}, nil
	case 0xE97CEFEBEECA398D:
		return &Node_SetBitOnVariable{}, nil
	case 0xFC108195473A878D:
		return &Node_OverridePowerLevelForEncounterWave{}, nil
	case 0x2212589FBE33C98D:
		return &Node_SetTraversePathDirectionEnabled{}, nil
	case 0xA8AD5D7DF30D398E:
		return &Node_SetDockEnabled{}, nil
	case 0xA2450476F9099F8E:
		return &Node_SetCharacterFX{}, nil
	case 0x2D2A949C7ECD298F:
		return &Node_PlayMaterialAnim{}, nil
	case 0x4989D6ABC788378F:
		return &Node_SetSubtitleDistanceThresholds{}, nil
	case 0xE929EAA15401A18F:
		return &Node_WolfSledSplineSectionExit{}, nil
	case 0x4EB946151140A68F:
		return &Node_SetFlightVolumeEnabledByType{}, nil
	case 0x46351222B9005E92:
		return &Node_CreateEncounterData{}, nil
	case 0x48AE1D67FFD63F90:
		return &Node_SetSubtitleDistanceThresholdsAlternate{}, nil
	case 0xBD2977F2C15C5790:
		return &Node_ClearEquipmentPreview{}, nil
	case 0x7B44093CB0869F90:
		return &Node_ClearMaxSpeedOverride{}, nil
	case 0x6AA211B734E6DE90:
		return &Node_SetCreatureHighlightCategoryOverride{}, nil
	case 0x1B637553DB144792:
		return &Node_ForceZeroTime{}, nil
	case 0x932CDBB618E06A92:
		return &Node_TriggerQuestManualCompleteBehaviors{}, nil
	case 0x58B8F950D55D7392:
		return &Node_DrawCone{}, nil
	case 0x24B5FD371D619692:
		return &Node_Warp{}, nil
	case 0x45980D6E32BA0093:
		return &Node_ForceLookAt{}, nil
	case 0x49978BF425648593:
		return &Node_SetElement{}, nil
	case 0x34EAA3B4230BE198:
		return &Node_EnableStopToAvoid{}, nil
	case 0x38186362C8A14A95:
		return &Node_EmitArrowWithEmitter{}, nil
	case 0xB979CEBE13C2BD94:
		return &Node_IncrementBanterFact{}, nil
	case 0x5D16F611413BEC94:
		return &Node_ExecuteVFS{}, nil
	case 0x108FD597EDB2FD94:
		return &Node_DespawnDynamicCharacter{}, nil
	case 0xED858B153B2F1B95:
		return &Node_EnableStoryTime{}, nil
	case 0x699A313A4A7F3795:
		return &Node_Crank_SetComplete{}, nil
	case 0xA70317F760F04E95:
		return &Node_EvaluateLoadZones{}, nil
	case 0xE54443A46FED9B95:
		return &Node_TriggerMoveEvent{}, nil
	case 0x8E7AFA908FD5E896:
		return &Node_DrawTextInWorld{}, nil
	case 0xFCC8A7AE429DF597:
		return &Node_ModifyFocalZoneEnabled{}, nil
	case 0x460CE4F00675AC98:
		return &Node_RemoveExcludeTraverseLinkFilterToCreature{}, nil
	case 0x0E74F748793AE199:
		return &Node_Pop{}, nil
	case 0xA8BD4E51ADF01499:
		return &Node_LockPlayerOnWolfsled{}, nil
	case 0x959BA82796F32999:
		return &Node_DrawText2D{}, nil
	case 0xB213B3202D415399:
		return &Node_SetPreventCheckpoint{}, nil
	case 0xFDB50F7AA4E88599:
		return &Node_Crank_FreeCrank{}, nil
	case 0xF380AF556AE79A99:
		return &Node_SetInteractZonesEnabled{}, nil
	case 0x30B3C9BD624CFC99:
		return &Node_RemoveCombatCollision{}, nil
	case 0x7FD48C222279B99A:
		return &Node_SetContextActionRangeData{}, nil
	case 0xF33CF886B66AD69A:
		return &Node_RegisterForLookAtEvent{}, nil
	case 0xE54754C1C7A55C9B:
		return &Node_SetWolfSledInteractOptions{}, nil
	case 0x3CF06B296D68F9AC:
		return &Node_SetInteractZoneTemplate{}, nil
	case 0xB58B12285ABC5DA3:
		n := &Node_PlaySound{}
		n.params[0], err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x1a)
		if err != nil {
			return nil, err
		}
		n.params[2], err = ReadNode(file, stack, offset+0x1e)
		if err != nil {
			return nil, err
		}
		n.params[3], err = ReadNode(file, stack, offset+0x22)
		if err != nil {
			return nil, err
		}
		n.params[4], err = ReadNode(file, stack, offset+0x26)
		if err != nil {
			return nil, err
		}
		n.params[5], err = ReadNode(file, stack, offset+0x2a)
		if err != nil {
			return nil, err
		}
		n.params[6], err = ReadNode(file, stack, offset+0x32)
		if err != nil {
			return nil, err
		}
		n.params[7], err = ReadNode(file, stack, offset+0x2e)
		if err != nil {
			return nil, err
		}
		n.params[8], err = ReadNode(file, stack, offset+0x36)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x06ED874C600BFA9F:
		return &Node_ThawPendulum{}, nil
	case 0x47C289B71BA5ED9D:
		return &Node_DestroyAllCombatConcussions{}, nil
	case 0x45A6AA43C708039C:
		return &Node_SetSoundEmitterSplineCap{}, nil
	case 0xED086487612CB69C:
		return &Node_ShowCompassMarker{}, nil
	case 0x2BD3C99FC77A029D:
		return &Node_SetPendulumRestrictionPlane{}, nil
	case 0x0B95A4D7AAD0089D:
		return &Node_CreateVFSFloat{}, nil
	case 0x0BADEF43F0BF459D:
		return &Node_ModifyFocalZoneObstacle{}, nil
	case 0x6E6D631004FB2A9E:
		return &Node_SetAdditionalGravityOnPendulum{}, nil
	case 0xFC8B572B55812F9E:
		return &Node_PlayUISound{}, nil
	case 0x91E07862FAB91C9F:
		return &Node_SetGIBlend{}, nil
	case 0xEE595C442C98349F:
		return &Node_SetSplineEnabled{}, nil
	case 0xFE58556DE495709F:
		return &Node_ClearWeaponCinematicMode{}, nil
	case 0xBDB26859BD0CA7A1:
		return &Node_SluiceSetup{}, nil
	case 0xE06930B6A95E2BA0:
		return &Node_SetSingleContextActionEnabled{}, nil
	case 0x154E1258828695A0:
		return &Node_SetPlayerNavCurveAutorun{}, nil
	case 0x71CD4766342DA7A0:
		return &Node_SetRotation{}, nil
	case 0xAEE9054F63B253A1:
		return &Node_RegisterOnEncounterStart{}, nil
	case 0xF21D95B0478585A1:
		return &Node_CreatePendulumWithAnchorJoint{}, nil
	case 0xD1CACD0FDD7DC5A1:
		return &Node_SetWolfSledSpeedControlOverride{}, nil
	case 0xAEA7CBEDA61807A2:
		return &Node_AddIncludeTraverseLinkFilterToCreature{}, nil
	case 0xE23288D20C8CE3A2:
		return &Node_AddPropToSyncWithCAOwner{}, nil
	case 0x814A6A0F5FF9EBA2:
		return &Node_RegisterForControlPressEvent{}, nil
	case 0xC589C70C1A5009A3:
		return &Node_SetSplineSpeedType{}, nil
	case 0xE907ED7BBD9878A7:
		return &Node_PlayAnimation{}, nil
	case 0xF99907F6D50948A5:
		return &Node_SetCreatureExitedWaitGate{}, nil
	case 0xE9B8521F3B7B01A4:
		return &Node_RemoveEmbedWeaponFromGameObject{}, nil
	case 0x570A24C77A612DA4:
		return &Node_ChangeBehaviorTree{}, nil
	case 0xB8915EA08DBD2FA4:
		return &Node_StopMusic{}, nil
	case 0x6BD3A840ABC9F3A4:
		return &Node_SetWwiseState{}, nil
	case 0x8B89BD2B78DD29A5:
		return &Node_ActivateBoatNavCurveFunctionalityAutoMovement{}, nil
	case 0x9BD48B594AFA54A5:
		return &Node_SendSystemicBanterResponse{}, nil
	case 0x020806227BCF7CA5:
		return &Node_PlayRestartableSoundLoop{}, nil
	case 0x7313E9BD6CF27FA5:
		return &Node_SetTraversePathCustomEntry{}, nil
	case 0x0C4F78EA31F1F6A5:
		return &Node_RecenterToFramePosition{}, nil
	case 0xD7756AF46F618BA6:
		return &Node_SetZiplineSpeedOverride{}, nil
	case 0x8F98E6CF544B4BAA:
		return &Node_DetermineSplineSpeedAndInputLocks{}, nil
	case 0xB8920460E2B483A7:
		return &Node_SetVFSInt{}, nil
	case 0xC1A2F2BC5950C3A7:
		return &Node_CopyUpgradeEquipState{}, nil
	case 0x9E7338CD4D3227A8:
		return &Node_SetMaxSpeedOverride{}, nil
	case 0xCA4B548CAC3847A8:
		return &Node_AddFocalZone{}, nil
	case 0xF3FF6D257BEDB4A8:
		return &Node_SetRealmLocked{}, nil
	case 0x2B63F1F9A07591AA:
		return &Node_AddPositionalFlag{}, nil
	case 0x4B003F6849C0B2AB:
		return &Node_DisableContextAction{}, nil
	case 0xC003C543D34D07AC:
		return &Node_ProfileEnd{}, nil
	case 0x0D2B060CAB7C40AC:
		return &Node_AddToGlobalVariable{}, nil
	case 0x7C73674E0DED74B3:
		return &Node_SetCameraTargetEnabled{}, nil
	case 0xE6890CA7D51AD5B1:
		return &Node_SendLuaHook{}, nil
	case 0xA42E1444D1908DAF:
		return &Node_GetArrayElement{}, nil
	case 0x89BE5980598A09AD:
		return &Node_RemoveCreatureFromCustomSplineGroup{}, nil
	case 0xDB0944412F63BFAD:
		return &Node_TransferFullLootResult{}, nil
	case 0x27ED992A8A91E0AD:
		return &Node_OverrideStickDamping{}, nil
	case 0x89CFAA38D5DF19AE:
		return &Node_SetHighlightEnabledOverride{}, nil
	case 0xCEE4429F322AA9AE:
		return &Node_SetMaxSpeedOverrideWithType{}, nil
	case 0xBBB3B7601DCDE1AF:
		return &Node_RegisterOnEncounterReset{}, nil
	case 0x6F714798792DECAF:
		return &Node_RegisterForDeathEvent{}, nil
	case 0x3DCE422724D3CFB0:
		return &Node_RestoreDefaultConversationSettings{}, nil
	case 0xC1800676C4A84DB1:
		return &Node_DespawnEncounter{}, nil
	case 0xC77FE43B811988B1:
		return &Node_SetIsDeathAllowed{}, nil
	case 0x5536CDAD8E4308B3:
		return &Node_ClearMaxSpeedOverrideWithType{}, nil
	case 0xE82753E110B6D7B1:
		return &Node_SkipWave{}, nil
	case 0xA5BC7B841A46EAB1:
		return &Node_ClearAllMarkUsed{}, nil
	case 0x93F317BBFFBA20B2:
		return &Node_EnableBoatForceDirectionControlMode{}, nil
	case 0x73FD14DB205D77B2:
		return &Node_EnableAllInteractTags{}, nil
	case 0xBEBB357D4BE2D4B2:
		return &Node_RegisterForHealthChangeEvent{}, nil
	case 0x586FBE26D77818B3:
		n := &Node_AcquirePickup{}
		n.params[0], err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x1a)
		if err != nil {
			return nil, err
		}
		n.params[2], err = ReadNode(file, stack, offset+0x1e)
		if err != nil {
			return nil, err
		}
		n.params[3], err = ReadNode(file, stack, offset+0x38)
		if err != nil {
			return nil, err
		}
		numParams := uint32(0)
		_, err = file.Seek(int64(offset+0x28), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &numParams)
		if err != nil {
			return nil, err
		}
		n.params2 = make([]Node, numParams)
		for i := uint32(0); i < numParams; i++ {
			n.params2[i], err = ReadNode(file, stack, offset+0x40+i*4)
			if err != nil {
				return nil, err
			}
		}
		return n, nil
	case 0x703FB7ADC34639B3:
		return &Node_StoreCheckpointWithPositionOverride{}, nil
	case 0x19B9CEED4CF96AB3:
		return &Node_HideCompassMarker{}, nil
	case 0x17E2EF454E236FB3:
		return &Node_CheckTriggerWithI{}, nil
	case 0x9B53B26CDC5F70B7:
		return &Node_BucketEquipmentByTrait{}, nil
	case 0xCF0F8A1ABC0E2BB5:
		return &Node_CloseSoundPortal{}, nil
	case 0x1428B5B813E68EB3:
		return &Node_HideWads{}, nil
	case 0x8354A048658DCCB3:
		return &Node_ApplyEquipmentJointsToObject{}, nil
	case 0xA47E97F13A6278B4:
		return &Node_HideHealthBar{}, nil
	case 0xEEC20F7878088DB4:
		return &Node_Remove{}, nil
	case 0x70D71059AA5CE3B4:
		return &Node_SetTraverseLinkEnabled{}, nil
	case 0xF6DCADF3E65E82B5:
		return &Node_RemoveFlag{}, nil
	case 0x72D321C22620FCB5:
		return &Node_ModifyFocalZone{}, nil
	case 0x800379829F5402B6:
		return &Node_FunctionCall{}, nil
		n := &Node_FunctionCall{}
		n.node, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		numParams := uint32(0)
		_, err = file.Seek(int64(offset+0x20), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &numParams)
		if err != nil {
			return nil, err
		}
		n.params = make([]Node, numParams)
		for i := uint32(0); i < numParams; i++ {
			n.params[i], err = ReadNode(file, stack, offset+0x60+i*4)
			if err != nil {
				return nil, err
			}
		}
		_, err = file.Seek(int64(offset+0x30), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &numParams)
		if err != nil {
			return nil, err
		}
		n.params2 = make([]Node, numParams)
		for i := uint32(0); i < numParams; i++ {
			n.params2[i], err = ReadNode(file, stack, offset+0x60+(uint32(len(n.params))+i)*4)
			if err != nil {
				return nil, err
			}
		}
		n.next, err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xE620D9C5E701D9B6:
		return &Node_SuspendCreatureCulling{}, nil
	case 0xD42C7509A63A00B7:
		return &Node_SetContextActionCustomData{}, nil
	case 0xE7DBC0432EA296B9:
		return &Node_Spawn{}, nil
	case 0x9C774F258C4097B7:
		return &Node_SetInteractZoneMainPrompts{}, nil
	case 0x7F719B5AFF7C01B8:
		return &Node_FilterIterator{}, nil
	case 0x9EFDE2F647C016B8:
		return &Node_AddLookAtPriorityOverride{}, nil
	case 0xFB9211B734B116B9:
		return &Node_SetInteractZoneSubPrompt{}, nil
	case 0x50A0D02A440589B9:
		return &Node_DeactivatePickup{}, nil
	case 0xC09EEC09B6972BBA:
		return &Node_TriggerPlayFX{}, nil
	case 0xFE5ACCE5749F5ABA:
		return &Node_SetTeam{}, nil
	case 0x1F06E7DA5EDC8CBA:
		return &Node_ClonePose{}, nil
	case 0xA32FF958880623BB:
		return &Node_ClearSoundEmitterSplineCap{}, nil
	case 0x33DE5CD422F239E0:
		return &Node_RemoveCombatIndicator{}, nil
	case 0x160858D1E26567CB:
		return &Node_ConfigureLogicLoadGroup{}, nil
	case 0xBD6B9D367BBBD0C3:
		return &Node_RegisterForEnemySpawn{}, nil
	case 0x1CA6830FB3AF14C0:
		return &Node_AssignCreatureToEncounter{}, nil
	case 0x6CB2FB587A8EF4BD:
		return &Node_SetSkillTreeToken{}, nil
	case 0xED799B7E6D9717BC:
		return &Node_FinishHapticVibration{}, nil
	case 0xF1D367B310F491BC:
		return &Node_SetParticlesEnabled{}, nil
	case 0x3A103AA04E69B1BC:
		return &Node_RegisterForAnimFrameEvent{}, nil
	case 0x2AF1E3155106B5BC:
		return &Node_SetThrowable{}, nil
	case 0x99FBF37F4C5E91BD:
		n := &Node_SetBlackboardVariable{}
		n.params[0], err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x1c)
		if err != nil {
			return nil, err
		}
		n.params[2], err = ReadNode(file, stack, offset+0x20)
		if err != nil {
			return nil, err
		}
		n.params[3], err = ReadNode(file, stack, offset+0x24)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x1a), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.param)
		if err != nil {
			return nil, err
		}
		n.next, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x3EA9812733C284BE:
		n := &Node_If{}
		n.Condition, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		n.True, err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		n.False, err = ReadNode(file, stack, offset+0x1a)
		if err != nil {
			return nil, err
		}
		if n.True == nil && n.False == nil {
			//fmt.Println("Warning: Both True and False branches are nil")
			return &Node_Dummy{}, nil
		}
		return n, nil
	case 0xCEA196DA307787BE:
		return &Node_ResetPadLightColor{}, nil
	case 0x7D8CF66F8FA13BBF:
		return &Node_SetWindMotorsEnabled{}, nil
	case 0x7492B9D10A2996BF:
		return &Node_MeterAdjustPauseDelay{}, nil
	case 0xC577213B3DAB0AC0:
		return &Node_ZoomSnap{}, nil
	case 0x5686F27FBD621BC2:
		return &Node_RelinquishPickup{}, nil
	case 0x93456DA09F6173C1:
		return &Node_SetWolfSledAutoRunSplineSpeed{}, nil
	case 0x23CF75E3714C80C1:
		return &Node_SetPaperBoatDistanceFromSplineClamp{}, nil
	case 0x88444EF88BAE90C1:
		return &Node_RemoveLookAtTarget{}, nil
	case 0x4C9F1A5A7BADADC1:
		return &Node_DrawCircle2D{}, nil
	case 0xD260E07CAF060EC2:
		return &Node_SetPendulumRestrictionAngle{}, nil
	case 0x9A2E9FED1B762DC2:
		return &Node_RegisterForFrameUpdateEvent{}, nil
	case 0xF7578140C5A266C2:
		return &Node_ClearGlobalVariable{}, nil
	case 0x0CBDEEFA3218BAC2:
		return &Node_RemoveLookAtPriorityOverride{}, nil
	case 0x104CBA15132F22C3:
		return &Node_SetZiplineStartAndEnd{}, nil
	case 0x75FFBF388E3C2AC3:
		return &Node_MusicFadeOutAndLogVolume{}, nil
	case 0x7B790CFD067064C8:
		return &Node_StartMainMusicTrack{}, nil
	case 0xD86A7DE7A38B3BC7:
		return &Node_SetPadLightColor{}, nil
	case 0x17030A7A0E6462C4:
		return &Node_EmitArrow{}, nil
	case 0x006615B249EA39C5:
		return &Node_RegisterForEventQueueProcessedEvent{}, nil
	case 0x0E669863732A85C6:
		return &Node_SetProgressionFact{}, nil
	case 0x89C2B1718383CFC6:
		return &Node_ResetTraverseLinkFilterForCreature{}, nil
	case 0x11EEF6D23D45F7C6:
		return &Node_CheckPersistentCreaturesEnabled{}, nil
	case 0x7DAC2138270FB2C7:
		return &Node_DisableLoadGate{}, nil
	case 0xDAE24153E3C2C1C7:
		return &Node_RemoveFX{}, nil
	case 0xC17719B48B2708C8:
		return &Node_RemoveArrows{}, nil
	case 0xB62D431D723214C8:
		return &Node_ResetTraverseLinkFilter{}, nil
	case 0x0766B9A6D53C2BC8:
		return &Node_FreeRequestedPrestreams{}, nil
	case 0x2715F8E04AB12ECA:
		return &Node_RecenterPendulum{}, nil
	case 0xCD8402C769416CC8:
		return &Node_SetWolfSledAutoRunEnableSteering{}, nil
	case 0x5C3EE4310A07ACC8:
		return &Node_SetSingleNavCurveEnabled{}, nil
	case 0x078AA15F42E331C9:
		return &Node_SetWeaponCinematicMode{}, nil
	case 0xA4594F86CDEB6BC9:
		return &Node_SendContextActionStim{}, nil
	case 0x13EC9DD37EBF17CA:
		return &Node_AddEquipmentPreview{}, nil
	case 0xF4552D883BDD64CA:
		return &Node_EquipToCharacterSlot{}, nil
	case 0xD28B3FC779BF1DCB:
		return &Node_StoreSoftSave{}, nil
	case 0x6890F438920F49CB:
		return &Node_ForceInterruptContextAction{}, nil
	case 0x590A375878C159CB:
		return &Node_TriggerEncounter{}, nil
	case 0x8F34F197FA833AD3:
		return &Node_SetCreaturePersistent{}, nil
	case 0x58DC3020AFE4CBCF:
		return &Node_SetSoundGeometryEnabled{}, nil
	case 0x911500A33D3E1BCD:
		return &Node_RegisterForButtonPressEvent{}, nil
	case 0x1181913D2D39E5CB:
		return &Node_SetHapticGlobalParameter{}, nil
	case 0x46493595ED2900CC:
		return &Node_ForceCompanionStriderPath{}, nil
	case 0xE8034B50CDC8AFCC:
		return &Node_Push{}, nil
	case 0x3E409BF05CDE14CD:
		return &Node_Sluice_WaterPourHit{}, nil
	case 0x4D5CD6B23FEB43CD:
		return &Node_GetParticleMonsterTotemCheckpointState{}, nil
	case 0xF5A02EDEBA664ACD:
		return &Node_AddTraverseLinkFilter{}, nil
	case 0x0859ED9B8EA393CD:
		return &Node_SaveYakWaterSpeedType{}, nil
	case 0x09814617B42945CE:
		return &Node_UpdateGameObjectSplineCheap{}, nil
	case 0xAAA4902BA47498CE:
		return &Node_SetInfluenceCircleAttributes{}, nil
	case 0x829CE0E9FB2E8FD1:
		return &Node_WildlifeEvent{}, nil
	case 0x62B1E65A1D890AD0:
		return &Node_RegisterOnEncounterCreated{}, nil
	case 0xE04618BA38E132D0:
		return &Node_SpawnDynamicCharacter{}, nil
	case 0xDDBFD135EE960ED1:
		return &Node_SendAwarenessStimWithPa{}, nil
	case 0x306FC73442ED64D1:
		n := &Node_SetVariable{}
		_, err := file.Seek(int64(offset+0x10), 0)
		if err != nil {
			return nil, err
		}
		variableOffset := int64(0)
		err = binary.Read(file, binary.LittleEndian, &variableOffset)
		if err != nil {
			return nil, err
		}
		if variableOffset>>0x3e == 1 {
			variableOffset = (variableOffset << 1) / 2
		}
		_, err = file.Seek(int64(offset+0x10)+variableOffset, 0)
		if err != nil {
			return nil, err
		}
		strLength := uint32(0)
		err = binary.Read(file, binary.LittleEndian, &strLength)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x48)+variableOffset, 0)
		if err != nil {
			return nil, err
		}
		strData := make([]byte, strLength)
		_, err = file.Read(strData)
		if err != nil {
			return nil, err
		}
		n.name = string(strData)
		n.val, err = ReadNode(file, stack, offset+0x1a)
		if err != nil {
			return nil, err
		}
		n.next, err = ReadNode(file, stack, offset+0x1e)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xFB7D4D63C03F86D1:
		return &Node_SetTraversePathSegmentEnabled{}, nil
	case 0x02EA73472B150DD2:
		return &Node_SetCreatureEnteredWaitGate{}, nil
	case 0x4A798C01AFB922D2:
		return &Node_ClearCreatureMotionParamsOverride{}, nil
	case 0x52AE2B1E2F5384D2:
		return &Node_RegisterForReticleTargetInvalid{}, nil
	case 0x9DD616B30C7DC0D2:
		return &Node_SetCompanionFollowStriderPath{}, nil
	case 0x2D156651A99FF8D2:
		return &Node_SetCompassEnabled{}, nil
	case 0x3D87A333EBE35DDA:
		return &Node_AddChild{}, nil
	case 0xEBB7567C0E0B03D6:
		return &Node_ForceUseSpline{}, nil
	case 0xC60095F026B241D3:
		return &Node_SetModelEnabled{}, nil
	case 0xC8F0D35E1DA069D3:
		return &Node_SendAwareness{}, nil
	case 0x8AF0389F805AA5D4:
		return &Node_SetVisualScriptEnabled{}, nil
	case 0xDA0A5BCFA82514D5:
		return &Node_UnregisterTraverseGraphEnteredEvent{}, nil
	case 0xEE1AD52FED2AFAD5:
		return &Node_CompleteWave{}, nil
	case 0x94F21B6443784FD6:
		return &Node_SetWolfSledAutoRunResetSettings{}, nil
	case 0xF5353AA0246B0ED7:
		return &Node_SetVFSFloat{}, nil
	case 0x9C7A8F0CBE5E38D7:
		return &Node_SetAnimFloatChannelDriverValue{}, nil
	case 0xD716A9489B8ABBD8:
		return &Node_UpdateStaticWindClientPositions{}, nil
	case 0xC08A6EF4FE48C3D8:
		return &Node_RotateObjectTowardsTarget{}, nil
	case 0xF56486D47F7AC7DD:
		return &Node_SetPaperBoatSpringValues{}, nil
	case 0x4BC17693B7199CDB:
		return &Node_CancelLockOn{}, nil
	case 0xBEEA97399149A3DC:
		return &Node_StartMusic{}, nil
	case 0x06D2FB09AC380DDD:
		return &Node_RegisterForUIButtonPressEvent{}, nil
	case 0x0790A5250A294DDD:
		return &Node_StopBlenderEffect{}, nil
	case 0x83F841C3053D7EDD:
		return &Node_SetForceCompanionTraverseLink{}, nil
	case 0x46BF5D0FA7360FDE:
		return &Node_SetDecalsEnabled{}, nil
	case 0x387B9CAAB2DE38DE:
		return &Node_ClearRage{}, nil
	case 0x4727B2C4817C3ADE:
		return &Node_GetInteractObject{}, nil
	case 0x3DB70B15D52A48DF:
		return &Node_SetSavedNavCurveAttributes{}, nil
	case 0xDADE3DCF0F2D8BF2:
		return &Node_SetSplineLeadDistances{}, nil
	case 0x514178653F5CD5E9:
		return &Node_ModifyFocalZoneObstacleEnabled{}, nil
	case 0x545D9E1D95BD28E5:
		return &Node_SpawnFX{}, nil
	case 0x4C3EFE4DEB0BCBE1:
		return &Node_SpherecastCollision{}, nil
	case 0x391AF8B8AF3243E0:
		return &Node_SimulateSpearZipline{}, nil
	case 0xE1B98A6DC7E479E0:
		return &Node_SetAllWeaponsCinematicMode{}, nil
	case 0x4115DE9E537B7BE0:
		return &Node_IncrementProgressionFact{}, nil
	case 0x81ED7185C3790CE1:
		return &Node_RegisterOnEnemyDeath{}, nil
	case 0x59614AE27D8F2DE1:
		return &Node_UnlockCombatStatus{}, nil
	case 0xE552C038B5FED7E1:
		return &Node_CompleteWaveElement{}, nil
	case 0x2719CF7D476FFDE2:
		return &Node_SetTraverseLinksEnabled{}, nil
	case 0x666D38C8B83528E3:
		return &Node_SetCharacterConfig{}, nil
	case 0x58E222C4EC3CFCE4:
		return &Node_EnableBoatForceTurnAroundControlMode{}, nil
	case 0x1E8B6A149CB61EE5:
		return &Node_FrameDelay{}, nil
	case 0xDCAF72DFAA47ECE7:
		return &Node_SetInteractTagEnabled{}, nil
	case 0xABF8484B8E2CE9E5:
		return &Node_DestroyPendulum{}, nil
	case 0x4AD10217ADD409E7:
		return &Node_SetZiplineEndDisconnected{}, nil
	case 0x431B8AAFE6711BE7:
		return &Node_DestroyGameObject{}, nil
	case 0xF62E6CE42D9073E7:
		return &Node_SetVolumeEnabled{}, nil
	case 0x8E50C18E35B79FE7:
		return &Node_ClearAllInteractZoneTags{}, nil
	case 0x8ED4956C99540FE8:
		return &Node_AddSimpleObjectActorToSequence{}, nil
	case 0x4DC272A52D3312E9:
		return &Node_PushWeaponStateForSequence{}, nil
	case 0x95F55CD91AB52CE9:
		return &Node_ClearCustomSplineGroup{}, nil
	case 0x08ED25FC75F798E9:
		return &Node_AddEquipmentToWallet{}, nil
	case 0x00959812C143BCE9:
		return &Node_CheckDecision{}, nil
	case 0x14D68E2697B0F5ED:
		return &Node_SetContextActionLoopActionData{}, nil
	case 0x18FBC2D5EC1A86EA:
		return &Node_SetOverrideEnterBranchesOnDock{}, nil
	case 0x3DF048F7A817DAE9:
		return &Node_SetOverrideStatusEffectIconParent{}, nil
	case 0xEA40D91F29E3ECE9:
		return &Node_SetWwiseSwitch{}, nil
	case 0x384DCA43F6210EEA:
		return &Node_Error{}, nil
	case 0x81F87ACD9E2570EA:
		return &Node_RequestMaterialSwap{}, nil
	case 0x6E74FF9A5F9F7DEA:
		return &Node_SetContextActionTraversalBehavior{}, nil
	case 0x699CD6CF59262AEB:
		return &Node_BatchQueryEquipmentTraits{}, nil
	case 0x6C34C065D76515EC:
		return &Node_ApplyContextIdle{}, nil
	case 0x44141A7DF026CBEC:
		return &Node_SetContextActionBasicBtreeData{}, nil
	case 0x14A7C3903DF5D0EC:
		return &Node_SubtractFromVariable{}, nil
	case 0x828474D4FC561AED:
		return &Node_SetCompanion{}, nil
	case 0x552AEA13CB7C14F0:
		return &Node_MusicFadeIn{}, nil
	case 0xF8DB059D6EF36EEE:
		return &Node_SpawnGameObject{}, nil
	case 0x780702AA5482B9EE:
		return &Node_SendPreStartNamedEvent{}, nil
	case 0x882DEF46A5A7DDEE:
		return &Node_FindChildGameObjectWithName{}, nil
	case 0x91CA30F7678B15EF:
		return &Node_UnequipFromCharacterSlot{}, nil
	case 0x67F09185C02292EF:
		return &Node_LoadFastTravelWads{}, nil
	case 0x912D41C1D8C16CF0:
		return &Node_SetComponentsEnabled{}, nil
	case 0x115DF0963B92E2F0:
		return &Node_AddLootResultToWallet{}, nil
	case 0x4B798F5D4EEE23F2:
		return &Node_SetAttachmentConfig{}, nil
	case 0x6D29E03DF73D41F2:
		return &Node_SetPlayGoRequired{}, nil
	case 0x81E90E4EAB478AF9:
		return &Node_SetMarkerClaimed{}, nil
	case 0x871F022830B708F5:
		return &Node_StoreCheckpoint{}, nil
	case 0xB871A4D9BE9FEFF3:
		return &Node_SetPaperBoatPitchRollValues{}, nil
	case 0x52CF3FF5AFD416F3:
		return &Node_RelinquishAllPickups{}, nil
	case 0x419022E1061F2DF3:
		return &Node_ModifyFocalZoneFacingAngleActivationTolerance{}, nil
	case 0x31A9AE98E5FA34F3:
		return &Node_Insert{}, nil
	case 0xCF50D29EA21DEAF3:
		return &Node_SetCollisionEnabled{}, nil
	case 0x002B6D0560A10BF4:
		return &Node_IncrementPickupStage{}, nil
	case 0xB3FA7525ED7588F4:
		return &Node_RemoveInteractZoneTag{}, nil
	case 0x89BC7EAAEBC2CFF4:
		return &Node_SetUDSActivityEnd{}, nil
	case 0x246174CD6CB7D2F4:
		return &Node_SetLightsEnabled{}, nil
	case 0x73A1AA47ABC3D6F4:
		return &Node_SetTraversePathEnabled{}, nil
	case 0x71B639DB54AF51F6:
		return &Node_RegisterForControlSwipeEvent{}, nil
	case 0x502EC1D347864BF5:
		return &Node_AbortCreatureInteract{}, nil
	case 0x9830AF3A232673F5:
		n := &Node_ClearBlackboardVariable{}
		n.params[0], err = ReadNode(file, stack, offset+0x16)
		if err != nil {
			return nil, err
		}
		n.params[1], err = ReadNode(file, stack, offset+0x1c)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(int64(offset+0x1a), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &n.param)
		if err != nil {
			return nil, err
		}
		n.next, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xACA2F86C9167CCF5:
		return &Node_SetParticleEmitterEnabled{}, nil
	case 0x64EFB050714EE0F5:
		return &Node_AddToVariable{}, nil
	case 0xA04DDE040CD640F6:
		return &Node_RecenterToAngle{}, nil
	case 0xED7172791C4A52F6:
		return &Node_RelinquishPickupBySlot{}, nil
	case 0x9200ECA91D4AE1F6:
		return &Node_AddCreatureActorToSequence{}, nil
	case 0x4CB79DDE2B22FBF7:
		return &Node_TriggerWave{}, nil
	case 0xADBD64976106C7F8:
		return &Node_SetUDSActivityUnavailable{}, nil
	case 0x55090118944CE2FD:
		return &Node_RegisterForWeaponThrown{}, nil
	case 0x47E12B02002160FB:
		return &Node_DestroyArrows{}, nil
	case 0x8AB1851B268709FA:
		return &Node_DecrementPickupStage{}, nil
	case 0x8AD3EF9F220936FA:
		return &Node_SetContextActionsEnabled{}, nil
	case 0x4810D4406CCE17FB:
		return &Node_ScentTracker_SetUpdateFrequency{}, nil
	case 0x4B674C79E5A41DFB:
		return &Node_SetContextActionTraversalExitSegment{}, nil
	case 0x119F4CD3CE2928FB:
		return &Node_DropChainedObject{}, nil
	case 0x84CF01EEDEEC8EFB:
		return &Node_DrawNavCurve{}, nil
	case 0x7A4232127681BFFB:
		return &Node_SetTimeScale{}, nil
	case 0x5432F51B14AD4CFC:
		return &Node_SetContextActionBranchData{}, nil
	case 0xEAA1FC9B92648BFC:
		return &Node_DrawLine{}, nil
	case 0xDA2EA6B42CA9F9FC:
		return &Node_SetInteractZonePromptEnabled{}, nil
	case 0x2C9625B907D114FF:
		return &Node_RegisterMinibossCheckpoint{}, nil
	case 0xABE9385ED79254FE:
		return &Node_RegisterObjectAsProp{}, nil
	case 0xB8FCC6DD299C97FE:
		return &Node_AlertWaveElement{}, nil
	case 0x5882181CE75BD2FE:
		n := &Node_GetCreature{}
		n.param, err = ReadNode(file, stack, offset+0x12)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0x1132BFDDC3CFD6FE:
		return &Node_LockOn{}, nil
	case 0xF5B303D764C3DCFE:
		return &Node_SetShowPlaceholderBoatOnDockDisable{}, nil
	case 0xC0B46C3F28F02DFF:
		return &Node_SetCompanionWaitOnTraversePath{}, nil
	case 0x1D32CC53BBAC71FF:
		return &Node_CineOnly_SwapCreatureToExistingObject{}, nil
	case 0x5C0F3DE76B8BBEFF:
		return &Node_LoadCheck{}, nil
	case 0xE345CC797D57D5FF:
		return &Node_ResetCompass{}, nil
	case 0xcc60de7cd3b6e663:
		return &Node_Event{}, nil
	case 0x269b61b515610c3b:
		return &Node_RegisteredEvent{}, nil
	case 0x7cc678873d79f660:
		return &Node_Data{}, nil
	case 0xe94fc2743690e481:
		return &Node_Triggerable{}, nil
	case 0x89f5f4d75ff26610:
		return &Node_ConstantDisplay{}, nil
	case 0xfea23d13757d05c5:
		return &Node_Constant{}, nil
	case 0x856bc88dce81cfd1:
		return &Node_EditorAssetReference{}, nil
	case 0xc5e296be1dd6ced9:
		return &Node_Struct{}, nil
	case 0xb5e4f22a23062ff8:
		return &Node_EnumReference{}, nil
	case 0x52f178c8c09070cf:
		return &Node_WadId{}, nil
	case 0x2c9fd77556119222:
		return &Node_RawUniqueId{}, nil
	case 0x9bdca2fde055f622:
		return &Node_RawStringHash{}, nil
	case 0x9df4beb458171ca0:
		return &Node_GenericGuidPathReference{}, nil
	case 0x0849b07af0eb138e:
		return &Node_GameObjectReference{}, nil
	case 0xfd79eb0f8d223a5b:
		return &Node_InstancedObjectReference{}, nil
	case 0x0d05dbd3dc6c4f76:
		return &Node_InteractZoneReference{}, nil
	case 0xc02f579701d823bc:
		return &Node_ArrowId{}, nil
	case 0xa6171fe17d54329e:
		return &Node_ArrowEmitterId{}, nil
	case 0x40c6b7e9778f4ac5:
		return &Node_BeamAttackId{}, nil
	case 0x95105785b3c257d9:
		return &Node_CameraId{}, nil
	case 0xcf228e78da60a4a5:
		return &Node_CameraRecenterId{}, nil
	case 0xe9de8c61895cdaf0:
		return &Node_GatewayMarkerId{}, nil
	case 0x4c49a8c78c504caa:
		return &Node_ConcussionId{}, nil
	case 0x29e7e9b37903ac59:
		return &Node_DecisionId{}, nil
	case 0xb49c7bea838afe04:
		return &Node_EquipmentId{}, nil
	case 0x88f073d90a625908:
		return &Node_IconId{}, nil
	case 0x79a0e934acbd6dfc:
		return &Node_MeterId{}, nil
	case 0xd62d4572caf7926d:
		return &Node_PickupId{}, nil
	case 0xd16793f5f58bb9b5:
		return &Node_PickupSlotId{}, nil
	case 0xfdaad40094d8a839:
		return &Node_QuestId{}, nil
	case 0xa478db9cd3a09ff6:
		return &Node_RumbleId{}, nil
	case 0x2e76833856f86838:
		return &Node_ScreenShakeId{}, nil
	case 0x622b77cd8bfdbfb4:
		return &Node_FullScreenEffectId{}, nil
	case 0x5a56ed75136d76eb:
		return &Node_BranchId{}, nil
	case 0xd2b018e91c1a93fa:
		return &Node_MapRealmId{}, nil
	case 0xae18acf3bfdcf152:
		return &Node_MapRegionId{}, nil
	case 0x178dfc16926e0844:
		return &Node_MapMarkerId{}, nil
	case 0xd1470e0ce80bc0d8:
		return &Node_MapAreaId{}, nil
	case 0xcde34ffb67f87d97:
		return &Node_PropId{}, nil
	case 0x048dca824cdd6c1d:
		return &Node_TweakConstantId{}, nil
	case 0xb50eceffb8abd291:
		return &Node_PlayFXId{}, nil
	case 0x78852150a54e3f86:
		return &Node_InteractZoneTemplateId{}, nil
	case 0xefc78cf842cb91c1:
		return &Node_AnimationId{}, nil
	case 0xba3cf0f381de832f:
		return &Node_AnimJointReference{}, nil
	case 0xc61a0e105014e038:
		return &Node_TraversePathReference{}, nil
	case 0x492efafdec6d9a69:
		return &Node_TraverseLinkReference{}, nil
	case 0x86644edfd60344c5:
		return &Node_ContextActionReference{}, nil
	case 0xade3a3c7d5ae113c:
		return &Node_NavCurveReference{}, nil
	case 0xe8ddee845605f702:
		return &Node_SoundEmitterReference{}, nil
	case 0x37615fcda854b0a4:
		return &Node_SoundPortalReference{}, nil
	case 0x6b36b9a958046d4c:
		return &Node_SoundProximityTriggerReference{}, nil
	case 0x36a9670bf45a9087:
		return &Node_BehaviorTreeRootReference{}, nil
	case 0x4885673f08024735:
		return &Node_BehaviorTreeSubtreeReference{}, nil
	case 0x23dcc4734469660b:
		return &Node_WwiseEventId{}, nil
	case 0xa72bb875f329502d:
		return &Node_WwiseSwitchGroupId{}, nil
	case 0x78d91938944b6264:
		return &Node_WwiseSwitchStateId{}, nil
	case 0x22321a16a46541cf:
		return &Node_WwiseRTPCId{}, nil
	case 0xd82f895868cbc7d9:
		return &Node_LamsId{}, nil
	case 0x1f4bfa75e62f6476:
		return &Node_BanterId{}, nil
	case 0x25805601e8a96131:
		return &Node_RecipeId{}, nil
	case 0xcdc7f7460c8cc34f:
		return &Node_ResourceId{}, nil
	case 0xa79f65af4efb7ab4:
		return &Node_WalletId{}, nil
	case 0x7cf65792f2480290:
		return &Node_WeaponId{}, nil
	case 0x47bdaa2fe5e117ba:
		return &Node_WildlifeId{}, nil
	case 0x663726d070490221:
		return &Node_HapticId{}, nil
	case 0x03920e3a93c6e400:
		return &Node_LootConditionId{}, nil
	case 0x0ceeaba96ca34bc7:
		return &Node_LootConditionSetId{}, nil
	case 0xe41d8772c589f070:
		return &Node_LootDistributorId{}, nil
	case 0x4a0ca99b7bf90834:
		return &Node_RelativeReference{}, nil
	case 0x220e63265fd2c566:
		return &Node_RelativeGameObjectReference{}, nil
	case 0x398229c1bf072da3:
		return &Node_RelativeInteractZoneReference{}, nil
	case 0x42c59fcb07f5e71b:
		return &Node_RelativeTraversePathReference{}, nil
	case 0xc11435c61ac3673b:
		return &Node_RelativeTraverseLinkReference{}, nil
	case 0x295d9e74f165fba5:
		return &Node_RelativeContextActionReference{}, nil
	case 0x7d7f523a1e1c8c86:
		return &Node_RelativeNavCurveReference{}, nil
	case 0x189cb06796275fbc:
		return &Node_RelativeSoundEmitterReference{}, nil
	case 0xcf1534664b91d850:
		return &Node_RelativeSoundPortalReference{}, nil
	case 0x873bd08d804f158e:
		return &Node_RelativeSoundProximityTriggerReference{}, nil
	case 0x657ccc60e8b3adb6:
		return &Node_RelativeAnimJoint{}, nil
	case 0xbc838245895da8de:
		return &Node_AsWadId{}, nil
	case 0x1414bbfb8eafb551:
		return &Node_TweaksIdAsStringHash{}, nil
	case 0x86de5d27641e2855:
		return &Node_ConcatenateAndHashRawStrings{}, nil
	case 0x04ae331e307a3420:
		return &Node_OnFastTravelWadsLoaded{}, nil
	case 0x926caeba606562e8:
		return &Node_OnBeamCreated{}, nil
	case 0x514d9afff1681639:
		if seen {
			return &Node_Variable{str: "onTimerComplete"}, nil
		}
		n := &Node_OnTimerComplete{}
		n.node, err = ReadNode(file, stack, offset+0x10)
		if err != nil {
			return nil, err
		}
		return n, nil
	case 0xcb8641b853fbcd9c:
		return &Node_OnFrameDelay{}, nil
	case 0x055359a614867cc6:
		return &Node_OnSpawnWildlifeComplete{}, nil
	case 0xa04a86f51c8e34c6:
		return &Node_OnSpawnComplete{}, nil
	case 0x90217b02c282d226:
		return &Node_OnSpawnedObjectFrameUpdate{}, nil
	case 0x7f2545f59c5ca5b8:
		return &Node_OnSpawnedObjectFrameUpdateExpired{}, nil
	case 0xa0261df565152550:
		return &Node_OnSpawnedObjectDestroyed{}, nil
	case 0x407762df190d8e8a:
		return &Node_OnSpawnedObjectRecycled{}, nil
	case 0xdbc491e75c2b829c:
		return &Node_OnSpawnFXComplete{}, nil
	case 0x5cd26a7454f3b9bb:
		return &Node_OnMaterialAnimationDone{}, nil
	case 0x231fe60f7af11389:
		return &Node_OnAnimationDone{}, nil
	case 0x8ddbc2744aa1a33f:
		return &Node_RegisterCombatIndicator{}, nil
	case 0x174bf2b7c7b7bc6a:
		return &Node_OnArrowEmitted{}, nil
	case 0xf0a27848bbaeffc2:
		return &Node_OnLoadCheckFinished{}, nil
	case 0xaf7b779367274516:
		return &Node_OnBlackboardTimerExpired{}, nil
	case 0xbf881bd3ad19d00e:
		return &Node_OnWarpComplete{}, nil
	case 0x2c7a27caaeca647b:
		return &Node_ModifyFocalZoneCameraAngleActivationTolerance{}, nil
	case 0x746ee266064e4228:
		return &Node_OnFocalZonePreActivationEnter{}, nil
	case 0xd49668ece1374231:
		return &Node_OnFocalZonePreActivationExit{}, nil
	case 0x515c5a812d44cb8c:
		return &Node_OnBanterStart{}, nil
	case 0x875031f40b6c907c:
		return &Node_OnBanterDone{}, nil
	case 0x42160ff9e068b081:
		return &Node_OnNextBanterDone{}, nil
	case 0x881ef391553740c9:
		return &Node_OnBanterFailed{}, nil
	case 0xc8da6c2869fe521f:
		return &Node_RemoveEquipmentFromWallet{}, nil
	case 0xe96efae266d22bb6:
		return &Node_OnUIMessageClosed{}, nil
	case 0x374b9370e2741624:
		return &Node_OnCallAndResponseTriggered{}, nil
	case 0x78028f7675b302d1:
		return &Node_OnCallAndResponseFinished{}, nil
	case 0x19f88e7729f15fec:
		return &Node_OnRegisteredFrameUpdate{}, nil
	case 0x8940960461f727b8:
		return &Node_OnAnimFrame{}, nil
	case 0xc52912a99e3029c2:
		return &Node_OnMovePlay{}, nil
	case 0xc94dc15319c6652a:
		return &Node_OnHealthChange{}, nil
	case 0xf3f35e11ad6680c9:
		return &Node_OnDeath{}, nil
	case 0xbdc033de21601dc5:
		return &Node_OnLookAt{}, nil
	case 0x0ea3372bcf3edaf3:
		return &Node_OnButtonPress{}, nil
	case 0x784b89e7c35a1b9d:
		return &Node_OnUIButtonPress{}, nil
	case 0x87c93b285cf1afd5:
		return &Node_OnControlPress{}, nil
	case 0x2c5d0590b77b7db1:
		return &Node_OnControlMash{}, nil
	case 0xb2efc2bf56af4552:
		return &Node_OnDistanceCheck{}, nil
	case 0xf0f5c75512f5557d:
		return &Node_OnEventQueueProcessed{}, nil
	case 0x18136aec00f11253:
		return &Node_OnRegisteredBreakableBroken{}, nil
	case 0x7b0c22f8691e872d:
		return &Node_OnRegisteredMeterChanged{}, nil
	case 0x3311c868a305e711:
		return &Node_OnControlSwipe{}, nil
	case 0x978d17b5d77e416d:
		return &Node_OnReticleTargetInvalid{}, nil
	case 0x1d0b8860ff8e4f75:
		return &Node_OnVariableChanged{}, nil
	case 0x7e41456b339534ce:
		return &Node_LoadGate{}, nil
	case 0x9a5c017fe16eb2a7:
		return &Node_OnStartGameFromThisWad{}, nil
	case 0xa6d8ecc2eaaa4512:
		return &Node_ClearAllWeather{}, nil
	case 0x078b1630af80bf08:
		return &Node_CreatureSpawnOptionsStruct{}, nil
	case 0x3247023efaa55f82:
		return &Node_OnEncounterStart{}, nil
	case 0x3e347f0a44f8fc8b:
		return &Node_OnEncounterCreated{}, nil
	case 0xa331adf5bd27ea6c:
		return &Node_OnEncounterFinished{}, nil
	case 0xe41883d29611e9ed:
		return &Node_OnEncounterReset{}, nil
	case 0xb7dcf355d3ae7665:
		return &Node_OnEnemySpawn{}, nil
	case 0x6865f5f27fcd9a8b:
		return &Node_OnDynamicCharacterSpawn{}, nil
	case 0xe273b915d4249fe1:
		return &Node_OnEnemyDeath{}, nil
	case 0xdb4e8cf9e7f209ab:
		return &Node_Crank_Callback{}, nil
	case 0xa9f5278bcde79690:
		return &Node_Sluice_Callback{}, nil
	case 0x97da1f4f120b4345:
		return &Node_Sluice_SoundCallback{}, nil
	case 0x49a41bd6dc3e2c6e:
		return &Node_OnInteractStart{}, nil
	case 0x695dc97456f1b038:
		return &Node_OnInteractFinished{}, nil
	case 0x2e8228b8c3c59c9c:
		return &Node_OnInteractAborted{}, nil
	case 0xc444784d3307e6c9:
		return &Node_OnInteractDone{}, nil
	case 0x71744f76784cc25d:
		return &Node_OnInteractAttemptedWhileOccupied{}, nil
	case 0xe20e4150a4da67c3:
		return &Node_OnTraverseGraphEntryEvent{}, nil
	case 0x14a3081dcbe7d1be:
		return &Node_OnTraverseGraphEnteredEvent{}, nil
	case 0x1923576cf1912be1:
		return &Node_OnCreatedVFSEntryChanged{}, nil
	case 0x7c088808310b8c8a:
		return &Node_OnCreatedVFSEntryExecuted{}, nil
	case 0xaefa146d7ad401fc:
		return &Node_OnConcussionHitBase{}, nil
	case 0x6e82923ca6494322:
		return &Node_OnArrowImpactBase{}, nil
	case 0x90daac765f2d2923:
		return &Node_OnCombatCollisionBase{}, nil
	case 0x0a2458fed1a8eea5:
		return &Node_OnWeaponThrown{}, nil
	case 0xe5d1ab0456da1291:
		return &Node_OnThrownWeaponBlockedBase{}, nil
	case 0xdaae217faa20e2d0:
		return &Node_NamedEventBase{}, nil
	case 0x8ccb347d8f347e2a:
		return &Node_OnImmediateEvent{}, nil
	case 0xe33ecdcf4f6c0eee:
		return &Node_SyncedSequenceActorInfo{}, nil
	case 0x0d04000a399016e0:
		return &Node_SyncedSequenceSimpleObjectActorInfo{}, nil
	case 0x8be15850bea252ec:
		return &Node_SyncedSequenceCreatureActorInfo{}, nil
	case 0xd52f5120e4ca98ea:
		return &Node_SyncedSequenceDynamicActorInfo{}, nil
	case 0x64378ab079270ae4:
		return &Node_SyncedSequenceSimpleObjectActorInfoSet{}, nil
	case 0x3c27b77d71bb6a20:
		return &Node_SyncedSequenceCreatureIdentifier{}, nil
	case 0x57ed7db35ce082ca:
		return &Node_SyncedSequenceCreatureActorInfoSet{}, nil
	case 0x3582cf8c3c09106c:
		return &Node_MotionWarpParameters{}, nil
	case 0xb5cd5b784b03b76a:
		return &Node_SyncedSequence{}, nil
	case 0x79d6757f054d54f7:
		return &Node_RegisteredSyncedSequence{}, nil
	case 0xf3ad9196ee051258:
		return &Node_DynamicSyncedSequenceEventContainer{}, nil
	case 0x7bd6c2178771e9c4:
		return &Node_SyncedSequenceEvent{}, nil
	case 0xff4aef85420d4a3f:
		return &Node_SyncedSequenceOnSkip{}, nil
	case 0x4c59ee9fe3f68832:
		return &Node_SyncedSequenceOnSkipExit{}, nil
	case 0xfa0e90b2e5be7491:
		return &Node_SyncedSequenceOnSequenceExit{}, nil
	case 0xcc9b2cfc9968d3c9:
		return &Node_SyncedSequenceOnTimeReached{}, nil
	case 0x20f7b326fadcb634:
		return &Node_SyncedSequenceOnMoveTimeReached{}, nil
	case 0x5a48e5adb73eaf17:
		return &Node_FocalZoneStrafe{}, nil
	case 0x8dd68602aba3bfff:
		return &Node_SyncedSequenceTime{}, nil
	default:
	}
	return nil, fmt.Errorf("Unknown node type: %X", nodeId)
}

// func ReadNode(file *os.File, header *Node_Header, offset uint32) (Node, error) {
// 	_, err := file.Seek(int64(offset), 0)
// 	if err != nil {
// 		return nil, err
// 	}
// 	nodeIndex := uint16(0)
// 	err = binary.Read(file, binary.LittleEndian, &nodeIndex)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if nodeIndex == 0xFFFF {
// 		return nil, nil
// 	}
// 	if index, ok := header.nodeMap[nodeIndex]; !ok {
// 		header.nodeMap[nodeIndex] = uint32(len(header.nodeMap))
// 	}
// 	return ReadNode(file, stack, nodeOffsets[nodeIndex])
// }
