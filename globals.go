package main

// This file initializes the global maps used to store data within the program.
var IdToAsset = map[uint64]Asset{}
var IdToListings = map[uint64]Listing{}
var IdToSimilarAssets = map[uint64][]uint64{}
var IdToListedIds = map[uint64][]uint64{}

var IdToCluster = map[uint64]int{}
var ClusterToAssetIds = [][]uint64{}
var ClusterToActiveAssetIds = [][]uint64{}

func AssetIdsToAssets(assetIds []uint64) []Asset {
	assets := make([]Asset, len(assetIds))
	for i, assetId := range assetIds {
		assets[i] = IdToAsset[assetId]
	}
	return assets
}

func AssetIdsToListings(assetIds []uint64) []Listing {
	listings := make([]Listing, len(assetIds))
	for i, assetId := range assetIds {
		listings[i] = IdToListings[assetId]
	}
	return listings
}
