package ffi

import (
	"io"
	"os"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/actors/runtime/proof"
	"github.com/ipfs/go-cid"
)

type SortedPrivateSectorInfo struct{ abi.SectorNumber }
type PrivateSectorInfo struct {
	abi.SectorNumber
	CacheDirPath     string
	PoStProofType    abi.RegisteredPoStProof
	SealedSectorPath string
	SectorInfo       proof.SectorInfo
}
type FallbackChallenges struct {
	Sectors    []abi.SectorNumber
	Challenges map[abi.SectorNumber][]uint64
}

func GenerateUnsealedCID(abi.RegisteredSealProof, []abi.PieceInfo) (cid.Cid, error) {
	return cid.Undef, nil
}
func GeneratePieceCIDFromFile(abi.RegisteredSealProof, io.Reader, abi.UnpaddedPieceSize) (cid.Cid, error) {
	return cid.Undef, nil
}
func UnsealRange(abi.RegisteredSealProof, string, *os.File, *os.File, abi.SectorNumber, abi.ActorID, abi.SealRandomness, cid.Cid, uint64, uint64) error {
	return nil
}
func SealPreCommitPhase1(abi.RegisteredSealProof, string, string, string, abi.SectorNumber, abi.ActorID, abi.SealRandomness, []abi.PieceInfo) ([]byte, error) {
	return nil, nil
}
func SealPreCommitPhase2([]byte, string, string) (cid.Cid, cid.Cid, error) {
	return cid.Undef, cid.Undef, nil
}
func SealCommitPhase1(abi.RegisteredSealProof, cid.Cid, cid.Cid, string, string, abi.SectorNumber, abi.ActorID, abi.SealRandomness, abi.InteractiveSealRandomness, []abi.PieceInfo) ([]byte, error) {
	return nil, nil
}
func SealCommitPhase2([]byte, abi.SectorNumber, abi.ActorID) ([]byte, error) {
	return nil, nil
}
func ClearCache(uint64, string) error {
	return nil
}
func GenerateWinningPoSt(abi.ActorID, SortedPrivateSectorInfo, abi.PoStRandomness) ([]proof.PoStProof, error) {
	return nil, nil
}
func GenerateWindowPoSt(abi.ActorID, SortedPrivateSectorInfo, abi.PoStRandomness) ([]proof.PoStProof, []abi.SectorNumber, error) {
	return nil, nil, nil
}
func NewSortedPrivateSectorInfo(...PrivateSectorInfo) SortedPrivateSectorInfo {
	return SortedPrivateSectorInfo{}
}
func VerifySeal(proof.SealVerifyInfo) (bool, error) {
	return true, nil
}
func VerifyWinningPoSt(proof.WinningPoStVerifyInfo) (bool, error) {
	return true, nil
}
func VerifyWindowPoSt(proof.WindowPoStVerifyInfo) (bool, error) {
	return true, nil
}
func GenerateWinningPoStSectorChallenge(abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error) {
	return nil, nil
}
func GeneratePoStFallbackSectorChallenges(abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, []abi.SectorNumber) (*FallbackChallenges, error) {
	return nil, nil
}
func GenerateSingleVanillaProof(PrivateSectorInfo, []uint64) ([]byte, error) {
	return nil, nil
}
func GetGPUDevices() ([]string, error) {
	return nil, nil
}
func FauxRep(abi.RegisteredSealProof, string, string) (cid.Cid, error) {
	return cid.Undef, nil
}
