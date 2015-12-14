package models

import (
	"database/sql"
	"fmt"
)

func (obj InterfaceConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS InterfaceConfig " +
		"( " +
		" Enabled bool " +
		" Description TEXT " +
		" Type TEXT " +
		" Name TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY( ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj InterfaceConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO InterfaceConfig (Enabled, Description, Type, Name, Mtu) VALUES (%v, %v, %v, %v, %v);",
		obj.Enabled, obj.Description, obj.Type, obj.Name, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj InterfaceConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from InterfaceConfig where rowid = " + objId
	fmt.Println("### DB Deleting InterfaceConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj InterfaceConfig) GetKey() (string, error) {
	return "", nil
}
func (obj InterfaceStateCounters) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS InterfaceStateCounters " +
		"( " +
		" Name TEXT " +
		" InMulticastPkts INTEGER " +
		" OutBroadcastPkts INTEGER " +
		" OperStatus INTEGER " +
		" InBroadcastPkts INTEGER " +
		" Enabled bool " +
		" OutDiscards INTEGER " +
		" Mtu INTEGER " +
		" OutOctets INTEGER " +
		" Ifindex INTEGER " +
		" InDiscards INTEGER " +
		" Type TEXT " +
		" InUnicastPkts INTEGER " +
		" LastClear TEXT " +
		" AdminStatus INTEGER " +
		" OutUnicastPkts INTEGER " +
		" Description TEXT " +
		" InOctets INTEGER " +
		" InUnknownProtos INTEGER " +
		" OutErrors INTEGER " +
		" LastChange INTEGER " +
		" InErrors INTEGER " +
		" OutMulticastPkts INTEGER " +
		"PRIMARY KEY( ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj InterfaceStateCounters) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO InterfaceStateCounters (Name, InMulticastPkts, OutBroadcastPkts, OperStatus, InBroadcastPkts, Enabled, OutDiscards, Mtu, OutOctets, Ifindex, InDiscards, Type, InUnicastPkts, LastClear, AdminStatus, OutUnicastPkts, Description, InOctets, InUnknownProtos, OutErrors, LastChange, InErrors, OutMulticastPkts) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Name, obj.InMulticastPkts, obj.OutBroadcastPkts, obj.OperStatus, obj.InBroadcastPkts, obj.Enabled, obj.OutDiscards, obj.Mtu, obj.OutOctets, obj.Ifindex, obj.InDiscards, obj.Type, obj.InUnicastPkts, obj.LastClear, obj.AdminStatus, obj.OutUnicastPkts, obj.Description, obj.InOctets, obj.InUnknownProtos, obj.OutErrors, obj.LastChange, obj.InErrors, obj.OutMulticastPkts)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj InterfaceStateCounters) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from InterfaceStateCounters where rowid = " + objId
	fmt.Println("### DB Deleting InterfaceStateCounters\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj InterfaceStateCounters) GetKey() (string, error) {
	return "", nil
}
func (obj HoldTimeConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS HoldTimeConfig " +
		"( " +
		" Down INTEGER " +
		" Description TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" Type TEXT " +
		" Up INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj HoldTimeConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO HoldTimeConfig (Down, Description, Enabled, NameKey, Type, Up, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.Down, obj.Description, obj.Enabled, obj.NameKey, obj.Type, obj.Up, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj HoldTimeConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from HoldTimeConfig where rowid = " + objId
	fmt.Println("### DB Deleting HoldTimeConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj HoldTimeConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj HoldTimeState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS HoldTimeState " +
		"( " +
		" Down INTEGER " +
		" Description TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" Type TEXT " +
		" Up INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj HoldTimeState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO HoldTimeState (Down, Description, Enabled, NameKey, Type, Up, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.Down, obj.Description, obj.Enabled, obj.NameKey, obj.Type, obj.Up, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj HoldTimeState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from HoldTimeState where rowid = " + objId
	fmt.Println("### DB Deleting HoldTimeState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj HoldTimeState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj SubinterfaceConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceConfig " +
		"( " +
		" Index INTEGER " +
		" Name TEXT " +
		" Enabled bool " +
		" Description TEXT " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY( ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceConfig (Index, Name, Enabled, Description, Unnumbered, Type, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.Index, obj.Name, obj.Enabled, obj.Description, obj.Unnumbered, obj.Type, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceConfig) GetKey() (string, error) {
	return "", nil
}
func (obj SubinterfaceStateCounters) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceStateCounters " +
		"( " +
		" Index INTEGER " +
		" Name TEXT " +
		" InMulticastPkts INTEGER " +
		" OutBroadcastPkts INTEGER " +
		" OperStatus INTEGER " +
		" InBroadcastPkts INTEGER " +
		" Enabled bool " +
		" OutDiscards INTEGER " +
		" Mtu INTEGER " +
		" OutOctets INTEGER " +
		" Ifindex INTEGER " +
		" InDiscards INTEGER " +
		" Type TEXT " +
		" InUnicastPkts INTEGER " +
		" LastClear TEXT " +
		" AdminStatus INTEGER " +
		" OutUnicastPkts INTEGER " +
		" Description TEXT " +
		" InOctets INTEGER " +
		" InUnknownProtos INTEGER " +
		" OutErrors INTEGER " +
		" LastChange INTEGER " +
		" InErrors INTEGER " +
		" OutMulticastPkts INTEGER " +
		" Unnumbered bool " +
		"PRIMARY KEY( ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceStateCounters) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceStateCounters (Index, Name, InMulticastPkts, OutBroadcastPkts, OperStatus, InBroadcastPkts, Enabled, OutDiscards, Mtu, OutOctets, Ifindex, InDiscards, Type, InUnicastPkts, LastClear, AdminStatus, OutUnicastPkts, Description, InOctets, InUnknownProtos, OutErrors, LastChange, InErrors, OutMulticastPkts, Unnumbered) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Index, obj.Name, obj.InMulticastPkts, obj.OutBroadcastPkts, obj.OperStatus, obj.InBroadcastPkts, obj.Enabled, obj.OutDiscards, obj.Mtu, obj.OutOctets, obj.Ifindex, obj.InDiscards, obj.Type, obj.InUnicastPkts, obj.LastClear, obj.AdminStatus, obj.OutUnicastPkts, obj.Description, obj.InOctets, obj.InUnknownProtos, obj.OutErrors, obj.LastChange, obj.InErrors, obj.OutMulticastPkts, obj.Unnumbered)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceStateCounters) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceStateCounters where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceStateCounters\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceStateCounters) GetKey() (string, error) {
	return "", nil
}
func (obj SubinterfaceVlanConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceVlanConfig " +
		"( " +
		" VlanId_VlanId_QinqId TEXT " +
		" IndexKey TEXT " +
		" Name TEXT " +
		" VlanId_VlanId INTEGER " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" GlobalVlanId TEXT " +
		" Description TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceVlanConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceVlanConfig (VlanId_VlanId_QinqId, IndexKey, Name, VlanId_VlanId, Enabled, Unnumbered, Type, GlobalVlanId, Description, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.VlanId_VlanId_QinqId, obj.IndexKey, obj.Name, obj.VlanId_VlanId, obj.Enabled, obj.Unnumbered, obj.Type, obj.GlobalVlanId, obj.Description, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceVlanConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceVlanConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceVlanConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceVlanConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceVlanState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceVlanState " +
		"( " +
		" VlanId_VlanId_QinqId TEXT " +
		" IndexKey TEXT " +
		" Name TEXT " +
		" VlanId_VlanId INTEGER " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" GlobalVlanId TEXT " +
		" Description TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceVlanState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceVlanState (VlanId_VlanId_QinqId, IndexKey, Name, VlanId_VlanId, Enabled, Unnumbered, Type, GlobalVlanId, Description, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.VlanId_VlanId_QinqId, obj.IndexKey, obj.Name, obj.VlanId_VlanId, obj.Enabled, obj.Unnumbered, obj.Type, obj.GlobalVlanId, obj.Description, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceVlanState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceVlanState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceVlanState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceVlanState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv4AddressConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4AddressConfig " +
		"( " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" Name TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" PrefixLength INTEGER " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4AddressConfig (IndexKey, Mtu, Name, Ip, Enabled, Unnumbered, Type, PrefixLength, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Mtu, obj.Name, obj.Ip, obj.Enabled, obj.Unnumbered, obj.Type, obj.PrefixLength, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4AddressConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4AddressConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4AddressConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv4AddressState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4AddressState " +
		"( " +
		" Origin INTEGER " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" Name TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" PrefixLength INTEGER " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4AddressState (Origin, IndexKey, Mtu, Name, Ip, Enabled, Unnumbered, Type, PrefixLength, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Origin, obj.IndexKey, obj.Mtu, obj.Name, obj.Ip, obj.Enabled, obj.Unnumbered, obj.Type, obj.PrefixLength, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4AddressState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4AddressState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4AddressState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4AddressVrrpVrrpGroupConfig " +
		"( " +
		" IndexKey TEXT " +
		" Description TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" VirtualRouterId INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" VirtualAddress TEXT " +
		" Name TEXT " +
		"PRIMARY KEY(IndexKey, IpKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressVrrpVrrpGroupConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4AddressVrrpVrrpGroupConfig (IndexKey, Description, PreemptDelay, PrefixLength, Enabled, IpKey, VirtualRouterId, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, Unnumbered, Type, VirtualAddress, Name) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Description, obj.PreemptDelay, obj.PrefixLength, obj.Enabled, obj.IpKey, obj.VirtualRouterId, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.Unnumbered, obj.Type, obj.VirtualAddress, obj.Name)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4AddressVrrpVrrpGroupConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4AddressVrrpVrrpGroupConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressVrrpVrrpGroupConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey) + "and" + "IpKey = " + string(obj.IpKey)
	return key, nil
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4AddressVrrpVrrpGroupState " +
		"( " +
		" IndexKey TEXT " +
		" Name TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" VirtualRouterId INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" CurrentPriority INTEGER " +
		" Unnumbered bool " +
		" Type TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey, IpKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressVrrpVrrpGroupState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4AddressVrrpVrrpGroupState (IndexKey, Name, PreemptDelay, PrefixLength, Enabled, IpKey, VirtualRouterId, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, CurrentPriority, Unnumbered, Type, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Name, obj.PreemptDelay, obj.PrefixLength, obj.Enabled, obj.IpKey, obj.VirtualRouterId, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.CurrentPriority, obj.Unnumbered, obj.Type, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4AddressVrrpVrrpGroupState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4AddressVrrpVrrpGroupState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressVrrpVrrpGroupState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey) + "and" + "IpKey = " + string(obj.IpKey)
	return key, nil
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig " +
		"( " +
		" IndexKey TEXT " +
		" Type TEXT " +
		" Name TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" PriorityDecrement INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" TrackInterface TEXT " +
		" Unnumbered bool " +
		" VirtualRouterIdKey TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey, IpKey, VirtualRouterIdKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig (IndexKey, Type, Name, PreemptDelay, PrefixLength, PriorityDecrement, Enabled, IpKey, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, TrackInterface, Unnumbered, VirtualRouterIdKey, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Type, obj.Name, obj.PreemptDelay, obj.PrefixLength, obj.PriorityDecrement, obj.Enabled, obj.IpKey, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.TrackInterface, obj.Unnumbered, obj.VirtualRouterIdKey, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey) + "and" + "IpKey = " + string(obj.IpKey) + "and" + "VirtualRouterIdKey = " + string(obj.VirtualRouterIdKey)
	return key, nil
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState " +
		"( " +
		" IndexKey TEXT " +
		" Type TEXT " +
		" Name TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" PriorityDecrement INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" TrackInterface TEXT " +
		" Unnumbered bool " +
		" VirtualRouterIdKey TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey, IpKey, VirtualRouterIdKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState (IndexKey, Type, Name, PreemptDelay, PrefixLength, PriorityDecrement, Enabled, IpKey, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, TrackInterface, Unnumbered, VirtualRouterIdKey, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Type, obj.Name, obj.PreemptDelay, obj.PrefixLength, obj.PriorityDecrement, obj.Enabled, obj.IpKey, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.TrackInterface, obj.Unnumbered, obj.VirtualRouterIdKey, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4AddressVrrpVrrpGroupInterfaceTrackingState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey) + "and" + "IpKey = " + string(obj.IpKey) + "and" + "VirtualRouterIdKey = " + string(obj.VirtualRouterIdKey)
	return key, nil
}

func (obj SubinterfaceIpv4NeighborConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4NeighborConfig " +
		"( " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" LinkLayerAddress TEXT " +
		" Name TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4NeighborConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4NeighborConfig (IndexKey, Mtu, LinkLayerAddress, Name, Ip, Enabled, Unnumbered, Type, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Mtu, obj.LinkLayerAddress, obj.Name, obj.Ip, obj.Enabled, obj.Unnumbered, obj.Type, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4NeighborConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4NeighborConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4NeighborConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4NeighborConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv4NeighborState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4NeighborState " +
		"( " +
		" Origin INTEGER " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" LinkLayerAddress TEXT " +
		" Name TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4NeighborState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4NeighborState (Origin, IndexKey, Mtu, LinkLayerAddress, Name, Ip, Enabled, Unnumbered, Type, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Origin, obj.IndexKey, obj.Mtu, obj.LinkLayerAddress, obj.Name, obj.Ip, obj.Enabled, obj.Unnumbered, obj.Type, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4NeighborState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4NeighborState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4NeighborState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4NeighborState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv4Config) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4Config " +
		"( " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" Name TEXT " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4Config) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4Config (IndexKey, Mtu, Name, Enabled, Unnumbered, Type, Description) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Mtu, obj.Name, obj.Enabled, obj.Unnumbered, obj.Type, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4Config) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4Config where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4Config\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4Config) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv4State) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv4State " +
		"( " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" Name TEXT " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4State) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv4State (IndexKey, Mtu, Name, Enabled, Unnumbered, Type, Description) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Mtu, obj.Name, obj.Enabled, obj.Unnumbered, obj.Type, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv4State) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv4State where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv4State\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv4State) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv6AddressConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6AddressConfig " +
		"( " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" Name TEXT " +
		" Ip TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" PrefixLength INTEGER " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6AddressConfig (IndexKey, Mtu, Name, Ip, DupAddrDetectTransmits, Enabled, Unnumbered, Type, PrefixLength, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Mtu, obj.Name, obj.Ip, obj.DupAddrDetectTransmits, obj.Enabled, obj.Unnumbered, obj.Type, obj.PrefixLength, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6AddressConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6AddressConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6AddressConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv6AddressState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6AddressState " +
		"( " +
		" Origin INTEGER " +
		" IndexKey TEXT " +
		" Status INTEGER " +
		" Name TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" PrefixLength INTEGER " +
		" DupAddrDetectTransmits INTEGER " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6AddressState (Origin, IndexKey, Status, Name, Ip, Enabled, Mtu, PrefixLength, DupAddrDetectTransmits, Unnumbered, Type, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Origin, obj.IndexKey, obj.Status, obj.Name, obj.Ip, obj.Enabled, obj.Mtu, obj.PrefixLength, obj.DupAddrDetectTransmits, obj.Unnumbered, obj.Type, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6AddressState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6AddressState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6AddressState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6AddressVrrpVrrpGroupConfig " +
		"( " +
		" IndexKey TEXT " +
		" Name TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" VirtualRouterId INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" DupAddrDetectTransmits INTEGER " +
		" VirtualLinkLocal TEXT " +
		" Unnumbered bool " +
		" Type TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey, IpKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressVrrpVrrpGroupConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6AddressVrrpVrrpGroupConfig (IndexKey, Name, PreemptDelay, PrefixLength, Enabled, IpKey, VirtualRouterId, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, DupAddrDetectTransmits, VirtualLinkLocal, Unnumbered, Type, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Name, obj.PreemptDelay, obj.PrefixLength, obj.Enabled, obj.IpKey, obj.VirtualRouterId, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.DupAddrDetectTransmits, obj.VirtualLinkLocal, obj.Unnumbered, obj.Type, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6AddressVrrpVrrpGroupConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6AddressVrrpVrrpGroupConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressVrrpVrrpGroupConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey) + "and" + "IpKey = " + string(obj.IpKey)
	return key, nil
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6AddressVrrpVrrpGroupState " +
		"( " +
		" IndexKey TEXT " +
		" Name TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" VirtualRouterId INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" CurrentPriority INTEGER " +
		" DupAddrDetectTransmits INTEGER " +
		" VirtualLinkLocal TEXT " +
		" Unnumbered bool " +
		" Type TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey, IpKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressVrrpVrrpGroupState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6AddressVrrpVrrpGroupState (IndexKey, Name, PreemptDelay, PrefixLength, Enabled, IpKey, VirtualRouterId, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, CurrentPriority, DupAddrDetectTransmits, VirtualLinkLocal, Unnumbered, Type, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Name, obj.PreemptDelay, obj.PrefixLength, obj.Enabled, obj.IpKey, obj.VirtualRouterId, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.CurrentPriority, obj.DupAddrDetectTransmits, obj.VirtualLinkLocal, obj.Unnumbered, obj.Type, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6AddressVrrpVrrpGroupState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6AddressVrrpVrrpGroupState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressVrrpVrrpGroupState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey) + "and" + "IpKey = " + string(obj.IpKey)
	return key, nil
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig " +
		"( " +
		" IndexKey TEXT " +
		" Type TEXT " +
		" TrackInterface TEXT " +
		" Name TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" PriorityDecrement INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" DupAddrDetectTransmits INTEGER " +
		" VirtualLinkLocal TEXT " +
		" Unnumbered bool " +
		" VirtualRouterIdKey TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey, IpKey, VirtualRouterIdKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig (IndexKey, Type, TrackInterface, Name, PreemptDelay, PrefixLength, PriorityDecrement, Enabled, IpKey, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, DupAddrDetectTransmits, VirtualLinkLocal, Unnumbered, VirtualRouterIdKey, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Type, obj.TrackInterface, obj.Name, obj.PreemptDelay, obj.PrefixLength, obj.PriorityDecrement, obj.Enabled, obj.IpKey, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.DupAddrDetectTransmits, obj.VirtualLinkLocal, obj.Unnumbered, obj.VirtualRouterIdKey, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey) + "and" + "IpKey = " + string(obj.IpKey) + "and" + "VirtualRouterIdKey = " + string(obj.VirtualRouterIdKey)
	return key, nil
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState " +
		"( " +
		" IndexKey TEXT " +
		" Type TEXT " +
		" TrackInterface TEXT " +
		" Name TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" PriorityDecrement INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" DupAddrDetectTransmits INTEGER " +
		" VirtualLinkLocal TEXT " +
		" Unnumbered bool " +
		" VirtualRouterIdKey TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey, IpKey, VirtualRouterIdKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState (IndexKey, Type, TrackInterface, Name, PreemptDelay, PrefixLength, PriorityDecrement, Enabled, IpKey, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, DupAddrDetectTransmits, VirtualLinkLocal, Unnumbered, VirtualRouterIdKey, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Type, obj.TrackInterface, obj.Name, obj.PreemptDelay, obj.PrefixLength, obj.PriorityDecrement, obj.Enabled, obj.IpKey, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.DupAddrDetectTransmits, obj.VirtualLinkLocal, obj.Unnumbered, obj.VirtualRouterIdKey, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AddressVrrpVrrpGroupInterfaceTrackingState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey) + "and" + "IpKey = " + string(obj.IpKey) + "and" + "VirtualRouterIdKey = " + string(obj.VirtualRouterIdKey)
	return key, nil
}

func (obj SubinterfaceIpv6NeighborConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6NeighborConfig " +
		"( " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" LinkLayerAddress TEXT " +
		" Name TEXT " +
		" Ip TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6NeighborConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6NeighborConfig (IndexKey, Mtu, LinkLayerAddress, Name, Ip, DupAddrDetectTransmits, Enabled, Unnumbered, Type, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Mtu, obj.LinkLayerAddress, obj.Name, obj.Ip, obj.DupAddrDetectTransmits, obj.Enabled, obj.Unnumbered, obj.Type, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6NeighborConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6NeighborConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6NeighborConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6NeighborConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv6NeighborState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6NeighborState " +
		"( " +
		" Origin INTEGER " +
		" IndexKey TEXT " +
		" Name TEXT " +
		" LinkLayerAddress TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" NeighborState INTEGER " +
		" IsRouter bool " +
		" DupAddrDetectTransmits INTEGER " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6NeighborState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6NeighborState (Origin, IndexKey, Name, LinkLayerAddress, Ip, Enabled, Mtu, NeighborState, IsRouter, DupAddrDetectTransmits, Unnumbered, Type, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Origin, obj.IndexKey, obj.Name, obj.LinkLayerAddress, obj.Ip, obj.Enabled, obj.Mtu, obj.NeighborState, obj.IsRouter, obj.DupAddrDetectTransmits, obj.Unnumbered, obj.Type, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6NeighborState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6NeighborState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6NeighborState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6NeighborState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv6Config) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6Config " +
		"( " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" Name TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6Config) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6Config (IndexKey, Mtu, Name, DupAddrDetectTransmits, Enabled, Unnumbered, Type, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Mtu, obj.Name, obj.DupAddrDetectTransmits, obj.Enabled, obj.Unnumbered, obj.Type, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6Config) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6Config where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6Config\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6Config) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv6State) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6State " +
		"( " +
		" IndexKey TEXT " +
		" Mtu INTEGER " +
		" Name TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" Enabled bool " +
		" Unnumbered bool " +
		" Type TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6State) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6State (IndexKey, Mtu, Name, DupAddrDetectTransmits, Enabled, Unnumbered, Type, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.Mtu, obj.Name, obj.DupAddrDetectTransmits, obj.Enabled, obj.Unnumbered, obj.Type, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6State) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6State where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6State\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6State) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv6AutoconfConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6AutoconfConfig " +
		"( " +
		" IndexKey TEXT " +
		" CreateGlobalAddresses bool " +
		" Name TEXT " +
		" CreateTemporaryAddresses bool " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" DupAddrDetectTransmits INTEGER " +
		" TemporaryPreferredLifetime INTEGER " +
		" Unnumbered bool " +
		" Type TEXT " +
		" TemporaryValidLifetime INTEGER " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AutoconfConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6AutoconfConfig (IndexKey, CreateGlobalAddresses, Name, CreateTemporaryAddresses, Enabled, Mtu, DupAddrDetectTransmits, TemporaryPreferredLifetime, Unnumbered, Type, TemporaryValidLifetime, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.CreateGlobalAddresses, obj.Name, obj.CreateTemporaryAddresses, obj.Enabled, obj.Mtu, obj.DupAddrDetectTransmits, obj.TemporaryPreferredLifetime, obj.Unnumbered, obj.Type, obj.TemporaryValidLifetime, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6AutoconfConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6AutoconfConfig where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6AutoconfConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AutoconfConfig) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj SubinterfaceIpv6AutoconfState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS SubinterfaceIpv6AutoconfState " +
		"( " +
		" IndexKey TEXT " +
		" CreateGlobalAddresses bool " +
		" Name TEXT " +
		" CreateTemporaryAddresses bool " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" DupAddrDetectTransmits INTEGER " +
		" TemporaryPreferredLifetime INTEGER " +
		" Unnumbered bool " +
		" Type TEXT " +
		" TemporaryValidLifetime INTEGER " +
		" Description TEXT " +
		"PRIMARY KEY(IndexKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AutoconfState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO SubinterfaceIpv6AutoconfState (IndexKey, CreateGlobalAddresses, Name, CreateTemporaryAddresses, Enabled, Mtu, DupAddrDetectTransmits, TemporaryPreferredLifetime, Unnumbered, Type, TemporaryValidLifetime, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.IndexKey, obj.CreateGlobalAddresses, obj.Name, obj.CreateTemporaryAddresses, obj.Enabled, obj.Mtu, obj.DupAddrDetectTransmits, obj.TemporaryPreferredLifetime, obj.Unnumbered, obj.Type, obj.TemporaryValidLifetime, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj SubinterfaceIpv6AutoconfState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from SubinterfaceIpv6AutoconfState where rowid = " + objId
	fmt.Println("### DB Deleting SubinterfaceIpv6AutoconfState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj SubinterfaceIpv6AutoconfState) GetKey() (string, error) {
	key := "IndexKey = " + string(obj.IndexKey)
	return key, nil
}

func (obj EthernetConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS EthernetConfig " +
		"( " +
		" MacAddress TEXT " +
		" Description TEXT " +
		" AggregateId TEXT " +
		" NameKey TEXT " +
		" Enabled bool " +
		" Speed TEXT " +
		" Mtu INTEGER " +
		" DuplexMode INTEGER " +
		" EnableFlowControl bool " +
		" Auto bool " +
		" Type TEXT " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO EthernetConfig (MacAddress, Description, AggregateId, NameKey, Enabled, Speed, Mtu, DuplexMode, EnableFlowControl, Auto, Type) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.MacAddress, obj.Description, obj.AggregateId, obj.NameKey, obj.Enabled, obj.Speed, obj.Mtu, obj.DuplexMode, obj.EnableFlowControl, obj.Auto, obj.Type)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj EthernetConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from EthernetConfig where rowid = " + objId
	fmt.Println("### DB Deleting EthernetConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj EthernetStateCounters) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS EthernetStateCounters " +
		"( " +
		" MacAddress TEXT " +
		" InCrcErrors INTEGER " +
		" Description TEXT " +
		" AggregateId TEXT " +
		" NameKey TEXT " +
		" InFragmentFrames INTEGER " +
		" Enabled bool " +
		" InMacControlFrames INTEGER " +
		" OutMacPauseFrames INTEGER " +
		" Mtu INTEGER " +
		" Out8021qFrames INTEGER " +
		" DuplexMode INTEGER " +
		" InMacPauseFrames INTEGER " +
		" In8021qFrames INTEGER " +
		" EnableFlowControl bool " +
		" InJabberFrames INTEGER " +
		" InOversizeFrames INTEGER " +
		" Type TEXT " +
		" OutMacControlFrames INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetStateCounters) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO EthernetStateCounters (MacAddress, InCrcErrors, Description, AggregateId, NameKey, InFragmentFrames, Enabled, InMacControlFrames, OutMacPauseFrames, Mtu, Out8021qFrames, DuplexMode, InMacPauseFrames, In8021qFrames, EnableFlowControl, InJabberFrames, InOversizeFrames, Type, OutMacControlFrames) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.MacAddress, obj.InCrcErrors, obj.Description, obj.AggregateId, obj.NameKey, obj.InFragmentFrames, obj.Enabled, obj.InMacControlFrames, obj.OutMacPauseFrames, obj.Mtu, obj.Out8021qFrames, obj.DuplexMode, obj.InMacPauseFrames, obj.In8021qFrames, obj.EnableFlowControl, obj.InJabberFrames, obj.InOversizeFrames, obj.Type, obj.OutMacControlFrames)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj EthernetStateCounters) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from EthernetStateCounters where rowid = " + objId
	fmt.Println("### DB Deleting EthernetStateCounters\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetStateCounters) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj EthernetState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS EthernetState " +
		"( " +
		" AggregateId TEXT " +
		" Description TEXT " +
		" MacAddress TEXT " +
		" Auto bool " +
		" Enabled bool " +
		" Speed TEXT " +
		" Mtu INTEGER " +
		" DuplexMode INTEGER " +
		" EnableFlowControl bool " +
		" NameKey TEXT " +
		" Type TEXT " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO EthernetState (AggregateId, Description, MacAddress, Auto, Enabled, Speed, Mtu, DuplexMode, EnableFlowControl, NameKey, Type) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.AggregateId, obj.Description, obj.MacAddress, obj.Auto, obj.Enabled, obj.Speed, obj.Mtu, obj.DuplexMode, obj.EnableFlowControl, obj.NameKey, obj.Type)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj EthernetState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from EthernetState where rowid = " + objId
	fmt.Println("### DB Deleting EthernetState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj EthernetVlanConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS EthernetVlanConfig " +
		"( " +
		" MacAddress TEXT " +
		" Description TEXT " +
		" NativeVlan_VlanId_QinqId TEXT " +
		" AggregateId TEXT " +
		" NativeVlan_VlanId INTEGER " +
		" NameKey TEXT " +
		" TrunkVlans_VlanId_VlanRange_QinqId TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" DuplexMode INTEGER " +
		" TrunkVlans_VlanId INTEGER " +
		" EnableFlowControl bool " +
		" AccessVlan_VlanId INTEGER " +
		" TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange TEXT " +
		" InterfaceMode INTEGER " +
		" Type TEXT " +
		" AccessVlan_VlanId_QinqId TEXT " +
		" TrunkVlans_VlanId_VlanRange TEXT " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetVlanConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO EthernetVlanConfig (MacAddress, Description, NativeVlan_VlanId_QinqId, AggregateId, NativeVlan_VlanId, NameKey, TrunkVlans_VlanId_VlanRange_QinqId, Enabled, Mtu, DuplexMode, TrunkVlans_VlanId, EnableFlowControl, AccessVlan_VlanId, TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange, InterfaceMode, Type, AccessVlan_VlanId_QinqId, TrunkVlans_VlanId_VlanRange) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.MacAddress, obj.Description, obj.NativeVlan_VlanId_QinqId, obj.AggregateId, obj.NativeVlan_VlanId, obj.NameKey, obj.TrunkVlans_VlanId_VlanRange_QinqId, obj.Enabled, obj.Mtu, obj.DuplexMode, obj.TrunkVlans_VlanId, obj.EnableFlowControl, obj.AccessVlan_VlanId, obj.TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange, obj.InterfaceMode, obj.Type, obj.AccessVlan_VlanId_QinqId, obj.TrunkVlans_VlanId_VlanRange)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj EthernetVlanConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from EthernetVlanConfig where rowid = " + objId
	fmt.Println("### DB Deleting EthernetVlanConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetVlanConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj EthernetVlanState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS EthernetVlanState " +
		"( " +
		" MacAddress TEXT " +
		" Description TEXT " +
		" NativeVlan_VlanId_QinqId TEXT " +
		" AggregateId TEXT " +
		" NativeVlan_VlanId INTEGER " +
		" NameKey TEXT " +
		" TrunkVlans_VlanId_VlanRange_QinqId TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" DuplexMode INTEGER " +
		" TrunkVlans_VlanId INTEGER " +
		" EnableFlowControl bool " +
		" AccessVlan_VlanId INTEGER " +
		" TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange TEXT " +
		" InterfaceMode INTEGER " +
		" Type TEXT " +
		" AccessVlan_VlanId_QinqId TEXT " +
		" TrunkVlans_VlanId_VlanRange TEXT " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetVlanState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO EthernetVlanState (MacAddress, Description, NativeVlan_VlanId_QinqId, AggregateId, NativeVlan_VlanId, NameKey, TrunkVlans_VlanId_VlanRange_QinqId, Enabled, Mtu, DuplexMode, TrunkVlans_VlanId, EnableFlowControl, AccessVlan_VlanId, TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange, InterfaceMode, Type, AccessVlan_VlanId_QinqId, TrunkVlans_VlanId_VlanRange) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.MacAddress, obj.Description, obj.NativeVlan_VlanId_QinqId, obj.AggregateId, obj.NativeVlan_VlanId, obj.NameKey, obj.TrunkVlans_VlanId_VlanRange_QinqId, obj.Enabled, obj.Mtu, obj.DuplexMode, obj.TrunkVlans_VlanId, obj.EnableFlowControl, obj.AccessVlan_VlanId, obj.TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange, obj.InterfaceMode, obj.Type, obj.AccessVlan_VlanId_QinqId, obj.TrunkVlans_VlanId_VlanRange)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj EthernetVlanState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from EthernetVlanState where rowid = " + objId
	fmt.Println("### DB Deleting EthernetVlanState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj EthernetVlanState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj AggregationConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS AggregationConfig " +
		"( " +
		" Description TEXT " +
		" LagType INTEGER " +
		" Enabled bool " +
		" NameKey TEXT " +
		" MinLinks INTEGER " +
		" Type TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO AggregationConfig (Description, LagType, Enabled, NameKey, MinLinks, Type, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.LagType, obj.Enabled, obj.NameKey, obj.MinLinks, obj.Type, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj AggregationConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from AggregationConfig where rowid = " + objId
	fmt.Println("### DB Deleting AggregationConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj AggregationState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS AggregationState " +
		"( " +
		" Description TEXT " +
		" Members TEXT " +
		" LagType INTEGER " +
		" Enabled bool " +
		" NameKey TEXT " +
		" MinLinks INTEGER " +
		" Type TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO AggregationState (Description, Members, LagType, Enabled, NameKey, MinLinks, Type, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.Members, obj.LagType, obj.Enabled, obj.NameKey, obj.MinLinks, obj.Type, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj AggregationState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from AggregationState where rowid = " + objId
	fmt.Println("### DB Deleting AggregationState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj AggregationLacpConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS AggregationLacpConfig " +
		"( " +
		" Description TEXT " +
		" MinLinks INTEGER " +
		" SystemPriority INTEGER " +
		" NameKey TEXT " +
		" Interval INTEGER " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" SystemIdMac TEXT " +
		" LagType INTEGER " +
		" Type TEXT " +
		" LacpMode INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationLacpConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO AggregationLacpConfig (Description, MinLinks, SystemPriority, NameKey, Interval, Enabled, Mtu, SystemIdMac, LagType, Type, LacpMode) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.MinLinks, obj.SystemPriority, obj.NameKey, obj.Interval, obj.Enabled, obj.Mtu, obj.SystemIdMac, obj.LagType, obj.Type, obj.LacpMode)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj AggregationLacpConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from AggregationLacpConfig where rowid = " + objId
	fmt.Println("### DB Deleting AggregationLacpConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationLacpConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj AggregationLacpState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS AggregationLacpState " +
		"( " +
		" Description TEXT " +
		" MinLinks INTEGER " +
		" SystemPriority INTEGER " +
		" NameKey TEXT " +
		" Interval INTEGER " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" SystemIdMac TEXT " +
		" LagType INTEGER " +
		" Type TEXT " +
		" LacpMode INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationLacpState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO AggregationLacpState (Description, MinLinks, SystemPriority, NameKey, Interval, Enabled, Mtu, SystemIdMac, LagType, Type, LacpMode) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.MinLinks, obj.SystemPriority, obj.NameKey, obj.Interval, obj.Enabled, obj.Mtu, obj.SystemIdMac, obj.LagType, obj.Type, obj.LacpMode)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj AggregationLacpState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from AggregationLacpState where rowid = " + objId
	fmt.Println("### DB Deleting AggregationLacpState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationLacpState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj AggregationLacpMemberStateCounters) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS AggregationLacpMemberStateCounters " +
		"( " +
		" SystemIdMac TEXT " +
		" MinLinks INTEGER " +
		" SystemPriority INTEGER " +
		" LacpUnknownErrors INTEGER " +
		" Interval INTEGER " +
		" Enabled bool " +
		" Aggregatable bool " +
		" OperKey INTEGER " +
		" Mtu INTEGER " +
		" Distributing bool " +
		" PartnerKey INTEGER " +
		" LacpErrors INTEGER " +
		" SystemId TEXT " +
		" Timeout INTEGER " +
		" Activity INTEGER " +
		" LacpRxErrors INTEGER " +
		" Type TEXT " +
		" Collecting bool " +
		" LagType INTEGER " +
		" Description TEXT " +
		" LacpTxErrors INTEGER " +
		" LacpOutPkts INTEGER " +
		" LacpInPkts INTEGER " +
		" Synchronization INTEGER " +
		" PartnerId TEXT " +
		" NameKey TEXT " +
		" Interface TEXT " +
		" LacpMode INTEGER " +
		"PRIMARY KEY(OperKey, PartnerKey, NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationLacpMemberStateCounters) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO AggregationLacpMemberStateCounters (SystemIdMac, MinLinks, SystemPriority, LacpUnknownErrors, Interval, Enabled, Aggregatable, OperKey, Mtu, Distributing, PartnerKey, LacpErrors, SystemId, Timeout, Activity, LacpRxErrors, Type, Collecting, LagType, Description, LacpTxErrors, LacpOutPkts, LacpInPkts, Synchronization, PartnerId, NameKey, Interface, LacpMode) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.SystemIdMac, obj.MinLinks, obj.SystemPriority, obj.LacpUnknownErrors, obj.Interval, obj.Enabled, obj.Aggregatable, obj.OperKey, obj.Mtu, obj.Distributing, obj.PartnerKey, obj.LacpErrors, obj.SystemId, obj.Timeout, obj.Activity, obj.LacpRxErrors, obj.Type, obj.Collecting, obj.LagType, obj.Description, obj.LacpTxErrors, obj.LacpOutPkts, obj.LacpInPkts, obj.Synchronization, obj.PartnerId, obj.NameKey, obj.Interface, obj.LacpMode)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj AggregationLacpMemberStateCounters) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from AggregationLacpMemberStateCounters where rowid = " + objId
	fmt.Println("### DB Deleting AggregationLacpMemberStateCounters\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationLacpMemberStateCounters) GetKey() (string, error) {
	key := "OperKey = " + string(obj.OperKey) + "and" + "PartnerKey = " + string(obj.PartnerKey) + "and" + "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj AggregationVlanConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS AggregationVlanConfig " +
		"( " +
		" Description TEXT " +
		" NativeVlan_VlanId_QinqId TEXT " +
		" MinLinks INTEGER " +
		" NativeVlan_VlanId INTEGER " +
		" NameKey TEXT " +
		" TrunkVlans_VlanId_VlanRange_QinqId TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" AccessVlan_VlanId INTEGER " +
		" TrunkVlans_VlanId INTEGER " +
		" LagType INTEGER " +
		" TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange TEXT " +
		" InterfaceMode INTEGER " +
		" Type TEXT " +
		" AccessVlan_VlanId_QinqId TEXT " +
		" TrunkVlans_VlanId_VlanRange TEXT " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationVlanConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO AggregationVlanConfig (Description, NativeVlan_VlanId_QinqId, MinLinks, NativeVlan_VlanId, NameKey, TrunkVlans_VlanId_VlanRange_QinqId, Enabled, Mtu, AccessVlan_VlanId, TrunkVlans_VlanId, LagType, TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange, InterfaceMode, Type, AccessVlan_VlanId_QinqId, TrunkVlans_VlanId_VlanRange) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.NativeVlan_VlanId_QinqId, obj.MinLinks, obj.NativeVlan_VlanId, obj.NameKey, obj.TrunkVlans_VlanId_VlanRange_QinqId, obj.Enabled, obj.Mtu, obj.AccessVlan_VlanId, obj.TrunkVlans_VlanId, obj.LagType, obj.TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange, obj.InterfaceMode, obj.Type, obj.AccessVlan_VlanId_QinqId, obj.TrunkVlans_VlanId_VlanRange)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj AggregationVlanConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from AggregationVlanConfig where rowid = " + objId
	fmt.Println("### DB Deleting AggregationVlanConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationVlanConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj AggregationVlanState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS AggregationVlanState " +
		"( " +
		" Description TEXT " +
		" NativeVlan_VlanId_QinqId TEXT " +
		" MinLinks INTEGER " +
		" NativeVlan_VlanId INTEGER " +
		" NameKey TEXT " +
		" TrunkVlans_VlanId_VlanRange_QinqId TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" AccessVlan_VlanId INTEGER " +
		" TrunkVlans_VlanId INTEGER " +
		" LagType INTEGER " +
		" TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange TEXT " +
		" InterfaceMode INTEGER " +
		" Type TEXT " +
		" AccessVlan_VlanId_QinqId TEXT " +
		" TrunkVlans_VlanId_VlanRange TEXT " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationVlanState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO AggregationVlanState (Description, NativeVlan_VlanId_QinqId, MinLinks, NativeVlan_VlanId, NameKey, TrunkVlans_VlanId_VlanRange_QinqId, Enabled, Mtu, AccessVlan_VlanId, TrunkVlans_VlanId, LagType, TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange, InterfaceMode, Type, AccessVlan_VlanId_QinqId, TrunkVlans_VlanId_VlanRange) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.NativeVlan_VlanId_QinqId, obj.MinLinks, obj.NativeVlan_VlanId, obj.NameKey, obj.TrunkVlans_VlanId_VlanRange_QinqId, obj.Enabled, obj.Mtu, obj.AccessVlan_VlanId, obj.TrunkVlans_VlanId, obj.LagType, obj.TrunkVlans_VlanId_VlanRange_QinqId_QinqIdRange, obj.InterfaceMode, obj.Type, obj.AccessVlan_VlanId_QinqId, obj.TrunkVlans_VlanId_VlanRange)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj AggregationVlanState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from AggregationVlanState where rowid = " + objId
	fmt.Println("### DB Deleting AggregationVlanState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj AggregationVlanState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanConfig " +
		"( " +
		" Vlan_Uint16 INTEGER " +
		" Description TEXT " +
		" Vlan_Uint16_String TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" Type TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanConfig (Vlan_Uint16, Description, Vlan_Uint16_String, Enabled, NameKey, Type, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.Vlan_Uint16, obj.Description, obj.Vlan_Uint16_String, obj.Enabled, obj.NameKey, obj.Type, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanState " +
		"( " +
		" Vlan_Uint16 INTEGER " +
		" Description TEXT " +
		" Vlan_Uint16_String TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" Type TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanState (Vlan_Uint16, Description, Vlan_Uint16_String, Enabled, NameKey, Type, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.Vlan_Uint16, obj.Description, obj.Vlan_Uint16_String, obj.Enabled, obj.NameKey, obj.Type, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv4AddressConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4AddressConfig " +
		"( " +
		" PrefixLength INTEGER " +
		" Type TEXT " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4AddressConfig (PrefixLength, Type, Description, vlan_String, Ip, Enabled, NameKey, vlan_Uint16, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.PrefixLength, obj.Type, obj.Description, obj.vlan_String, obj.Ip, obj.Enabled, obj.NameKey, obj.vlan_Uint16, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4AddressConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4AddressConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4AddressConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv4AddressState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4AddressState " +
		"( " +
		" Origin INTEGER " +
		" PrefixLength INTEGER " +
		" Type TEXT " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4AddressState (Origin, PrefixLength, Type, Description, vlan_String, Ip, Enabled, NameKey, vlan_Uint16, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Origin, obj.PrefixLength, obj.Type, obj.Description, obj.vlan_String, obj.Ip, obj.Enabled, obj.NameKey, obj.vlan_Uint16, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4AddressState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4AddressState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4AddressState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4AddressVrrpVrrpGroupConfig " +
		"( " +
		" PrefixLength INTEGER " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" PreemptDelay INTEGER " +
		" NameKey TEXT " +
		" Enabled bool " +
		" IpKey TEXT " +
		" VirtualRouterId INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" vlan_Uint16 INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" AdvertisementInterval INTEGER " +
		" Type TEXT " +
		" VirtualAddress TEXT " +
		"PRIMARY KEY(NameKey, IpKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressVrrpVrrpGroupConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4AddressVrrpVrrpGroupConfig (PrefixLength, Description, vlan_String, PreemptDelay, NameKey, Enabled, IpKey, VirtualRouterId, Mtu, Priority, vlan_Uint16, Preempt, AcceptMode, AdvertisementInterval, Type, VirtualAddress) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.PrefixLength, obj.Description, obj.vlan_String, obj.PreemptDelay, obj.NameKey, obj.Enabled, obj.IpKey, obj.VirtualRouterId, obj.Mtu, obj.Priority, obj.vlan_Uint16, obj.Preempt, obj.AcceptMode, obj.AdvertisementInterval, obj.Type, obj.VirtualAddress)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4AddressVrrpVrrpGroupConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4AddressVrrpVrrpGroupConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressVrrpVrrpGroupConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey) + "and" + "IpKey = " + string(obj.IpKey)
	return key, nil
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4AddressVrrpVrrpGroupState " +
		"( " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" VirtualRouterId INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" vlan_Uint16 INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" CurrentPriority INTEGER " +
		" NameKey TEXT " +
		" Type TEXT " +
		" VirtualAddress TEXT " +
		" AdvertisementInterval INTEGER " +
		"PRIMARY KEY(IpKey, NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressVrrpVrrpGroupState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4AddressVrrpVrrpGroupState (Description, vlan_String, PreemptDelay, PrefixLength, Enabled, IpKey, VirtualRouterId, Mtu, Priority, vlan_Uint16, Preempt, AcceptMode, CurrentPriority, NameKey, Type, VirtualAddress, AdvertisementInterval) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.vlan_String, obj.PreemptDelay, obj.PrefixLength, obj.Enabled, obj.IpKey, obj.VirtualRouterId, obj.Mtu, obj.Priority, obj.vlan_Uint16, obj.Preempt, obj.AcceptMode, obj.CurrentPriority, obj.NameKey, obj.Type, obj.VirtualAddress, obj.AdvertisementInterval)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4AddressVrrpVrrpGroupState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4AddressVrrpVrrpGroupState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressVrrpVrrpGroupState) GetKey() (string, error) {
	key := "IpKey = " + string(obj.IpKey) + "and" + "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig " +
		"( " +
		" PrefixLength INTEGER " +
		" Type TEXT " +
		" PreemptDelay INTEGER " +
		" vlan_String TEXT " +
		" vlan_Uint16 INTEGER " +
		" PriorityDecrement INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" NameKey TEXT " +
		" TrackInterface TEXT " +
		" VirtualRouterIdKey TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IpKey, NameKey, VirtualRouterIdKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig (PrefixLength, Type, PreemptDelay, vlan_String, vlan_Uint16, PriorityDecrement, Enabled, IpKey, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, NameKey, TrackInterface, VirtualRouterIdKey, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.PrefixLength, obj.Type, obj.PreemptDelay, obj.vlan_String, obj.vlan_Uint16, obj.PriorityDecrement, obj.Enabled, obj.IpKey, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.NameKey, obj.TrackInterface, obj.VirtualRouterIdKey, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingConfig) GetKey() (string, error) {
	key := "IpKey = " + string(obj.IpKey) + "and" + "NameKey = " + string(obj.NameKey) + "and" + "VirtualRouterIdKey = " + string(obj.VirtualRouterIdKey)
	return key, nil
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState " +
		"( " +
		" PrefixLength INTEGER " +
		" Type TEXT " +
		" PreemptDelay INTEGER " +
		" vlan_String TEXT " +
		" vlan_Uint16 INTEGER " +
		" PriorityDecrement INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" NameKey TEXT " +
		" TrackInterface TEXT " +
		" VirtualRouterIdKey TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(IpKey, NameKey, VirtualRouterIdKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState (PrefixLength, Type, PreemptDelay, vlan_String, vlan_Uint16, PriorityDecrement, Enabled, IpKey, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, NameKey, TrackInterface, VirtualRouterIdKey, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.PrefixLength, obj.Type, obj.PreemptDelay, obj.vlan_String, obj.vlan_Uint16, obj.PriorityDecrement, obj.Enabled, obj.IpKey, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.NameKey, obj.TrackInterface, obj.VirtualRouterIdKey, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4AddressVrrpVrrpGroupInterfaceTrackingState) GetKey() (string, error) {
	key := "IpKey = " + string(obj.IpKey) + "and" + "NameKey = " + string(obj.NameKey) + "and" + "VirtualRouterIdKey = " + string(obj.VirtualRouterIdKey)
	return key, nil
}

func (obj RoutedVlanIpv4NeighborConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4NeighborConfig " +
		"( " +
		" LinkLayerAddress TEXT " +
		" Type TEXT " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4NeighborConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4NeighborConfig (LinkLayerAddress, Type, Description, vlan_String, Ip, Enabled, NameKey, vlan_Uint16, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.LinkLayerAddress, obj.Type, obj.Description, obj.vlan_String, obj.Ip, obj.Enabled, obj.NameKey, obj.vlan_Uint16, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4NeighborConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4NeighborConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4NeighborConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4NeighborConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv4NeighborState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4NeighborState " +
		"( " +
		" Origin INTEGER " +
		" LinkLayerAddress TEXT " +
		" Type TEXT " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Ip TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4NeighborState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4NeighborState (Origin, LinkLayerAddress, Type, Description, vlan_String, Ip, Enabled, NameKey, vlan_Uint16, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Origin, obj.LinkLayerAddress, obj.Type, obj.Description, obj.vlan_String, obj.Ip, obj.Enabled, obj.NameKey, obj.vlan_Uint16, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4NeighborState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4NeighborState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4NeighborState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4NeighborState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv4Config) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4Config " +
		"( " +
		" vlan_Uint16 INTEGER " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" Type TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4Config) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4Config (vlan_Uint16, Description, vlan_String, Enabled, NameKey, Type, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.vlan_Uint16, obj.Description, obj.vlan_String, obj.Enabled, obj.NameKey, obj.Type, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4Config) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4Config where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4Config\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4Config) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv4State) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv4State " +
		"( " +
		" vlan_Uint16 INTEGER " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" Type TEXT " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4State) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv4State (vlan_Uint16, Description, vlan_String, Enabled, NameKey, Type, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v);",
		obj.vlan_Uint16, obj.Description, obj.vlan_String, obj.Enabled, obj.NameKey, obj.Type, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv4State) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv4State where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv4State\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv4State) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6AddressConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6AddressConfig " +
		"( " +
		" PrefixLength INTEGER " +
		" Type TEXT " +
		" Enabled bool " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Ip TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" NameKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6AddressConfig (PrefixLength, Type, Enabled, Description, vlan_String, Ip, DupAddrDetectTransmits, NameKey, vlan_Uint16, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.PrefixLength, obj.Type, obj.Enabled, obj.Description, obj.vlan_String, obj.Ip, obj.DupAddrDetectTransmits, obj.NameKey, obj.vlan_Uint16, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6AddressConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6AddressConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6AddressConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6AddressState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6AddressState " +
		"( " +
		" Origin INTEGER " +
		" PrefixLength INTEGER " +
		" Status INTEGER " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" NameKey TEXT " +
		" Type TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" Ip TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" vlan_Uint16 INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6AddressState (Origin, PrefixLength, Status, Description, vlan_String, NameKey, Type, Enabled, Mtu, Ip, DupAddrDetectTransmits, vlan_Uint16) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Origin, obj.PrefixLength, obj.Status, obj.Description, obj.vlan_String, obj.NameKey, obj.Type, obj.Enabled, obj.Mtu, obj.Ip, obj.DupAddrDetectTransmits, obj.vlan_Uint16)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6AddressState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6AddressState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6AddressState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6AddressVrrpVrrpGroupConfig " +
		"( " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" VirtualRouterId INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" vlan_Uint16 INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" NameKey TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" VirtualLinkLocal TEXT " +
		" Type TEXT " +
		" VirtualAddress TEXT " +
		" AdvertisementInterval INTEGER " +
		"PRIMARY KEY(IpKey, NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressVrrpVrrpGroupConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6AddressVrrpVrrpGroupConfig (Description, vlan_String, PreemptDelay, PrefixLength, Enabled, IpKey, VirtualRouterId, Mtu, Priority, vlan_Uint16, Preempt, AcceptMode, NameKey, DupAddrDetectTransmits, VirtualLinkLocal, Type, VirtualAddress, AdvertisementInterval) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.vlan_String, obj.PreemptDelay, obj.PrefixLength, obj.Enabled, obj.IpKey, obj.VirtualRouterId, obj.Mtu, obj.Priority, obj.vlan_Uint16, obj.Preempt, obj.AcceptMode, obj.NameKey, obj.DupAddrDetectTransmits, obj.VirtualLinkLocal, obj.Type, obj.VirtualAddress, obj.AdvertisementInterval)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6AddressVrrpVrrpGroupConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6AddressVrrpVrrpGroupConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressVrrpVrrpGroupConfig) GetKey() (string, error) {
	key := "IpKey = " + string(obj.IpKey) + "and" + "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6AddressVrrpVrrpGroupState " +
		"( " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" PreemptDelay INTEGER " +
		" PrefixLength INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" VirtualRouterId INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" vlan_Uint16 INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" CurrentPriority INTEGER " +
		" NameKey TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" VirtualLinkLocal TEXT " +
		" Type TEXT " +
		" VirtualAddress TEXT " +
		" AdvertisementInterval INTEGER " +
		"PRIMARY KEY(IpKey, NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressVrrpVrrpGroupState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6AddressVrrpVrrpGroupState (Description, vlan_String, PreemptDelay, PrefixLength, Enabled, IpKey, VirtualRouterId, Mtu, Priority, vlan_Uint16, Preempt, AcceptMode, CurrentPriority, NameKey, DupAddrDetectTransmits, VirtualLinkLocal, Type, VirtualAddress, AdvertisementInterval) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Description, obj.vlan_String, obj.PreemptDelay, obj.PrefixLength, obj.Enabled, obj.IpKey, obj.VirtualRouterId, obj.Mtu, obj.Priority, obj.vlan_Uint16, obj.Preempt, obj.AcceptMode, obj.CurrentPriority, obj.NameKey, obj.DupAddrDetectTransmits, obj.VirtualLinkLocal, obj.Type, obj.VirtualAddress, obj.AdvertisementInterval)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6AddressVrrpVrrpGroupState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6AddressVrrpVrrpGroupState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressVrrpVrrpGroupState) GetKey() (string, error) {
	key := "IpKey = " + string(obj.IpKey) + "and" + "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig " +
		"( " +
		" TrackInterface TEXT " +
		" Type TEXT " +
		" PrefixLength INTEGER " +
		" vlan_String TEXT " +
		" PreemptDelay INTEGER " +
		" NameKey TEXT " +
		" PriorityDecrement INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" DupAddrDetectTransmits INTEGER " +
		" VirtualLinkLocal TEXT " +
		" VirtualRouterIdKey TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(NameKey, IpKey, VirtualRouterIdKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig (TrackInterface, Type, PrefixLength, vlan_String, PreemptDelay, NameKey, PriorityDecrement, Enabled, IpKey, vlan_Uint16, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, DupAddrDetectTransmits, VirtualLinkLocal, VirtualRouterIdKey, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.TrackInterface, obj.Type, obj.PrefixLength, obj.vlan_String, obj.PreemptDelay, obj.NameKey, obj.PriorityDecrement, obj.Enabled, obj.IpKey, obj.vlan_Uint16, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.DupAddrDetectTransmits, obj.VirtualLinkLocal, obj.VirtualRouterIdKey, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey) + "and" + "IpKey = " + string(obj.IpKey) + "and" + "VirtualRouterIdKey = " + string(obj.VirtualRouterIdKey)
	return key, nil
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState " +
		"( " +
		" TrackInterface TEXT " +
		" Type TEXT " +
		" PrefixLength INTEGER " +
		" vlan_String TEXT " +
		" PreemptDelay INTEGER " +
		" NameKey TEXT " +
		" PriorityDecrement INTEGER " +
		" Enabled bool " +
		" IpKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		" Priority INTEGER " +
		" AdvertisementInterval INTEGER " +
		" Preempt bool " +
		" AcceptMode bool " +
		" DupAddrDetectTransmits INTEGER " +
		" VirtualLinkLocal TEXT " +
		" VirtualRouterIdKey TEXT " +
		" VirtualAddress TEXT " +
		" Description TEXT " +
		"PRIMARY KEY(NameKey, IpKey, VirtualRouterIdKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState (TrackInterface, Type, PrefixLength, vlan_String, PreemptDelay, NameKey, PriorityDecrement, Enabled, IpKey, vlan_Uint16, Mtu, Priority, AdvertisementInterval, Preempt, AcceptMode, DupAddrDetectTransmits, VirtualLinkLocal, VirtualRouterIdKey, VirtualAddress, Description) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.TrackInterface, obj.Type, obj.PrefixLength, obj.vlan_String, obj.PreemptDelay, obj.NameKey, obj.PriorityDecrement, obj.Enabled, obj.IpKey, obj.vlan_Uint16, obj.Mtu, obj.Priority, obj.AdvertisementInterval, obj.Preempt, obj.AcceptMode, obj.DupAddrDetectTransmits, obj.VirtualLinkLocal, obj.VirtualRouterIdKey, obj.VirtualAddress, obj.Description)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AddressVrrpVrrpGroupInterfaceTrackingState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey) + "and" + "IpKey = " + string(obj.IpKey) + "and" + "VirtualRouterIdKey = " + string(obj.VirtualRouterIdKey)
	return key, nil
}

func (obj RoutedVlanIpv6NeighborConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6NeighborConfig " +
		"( " +
		" LinkLayerAddress TEXT " +
		" Type TEXT " +
		" Enabled bool " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Ip TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" NameKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6NeighborConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6NeighborConfig (LinkLayerAddress, Type, Enabled, Description, vlan_String, Ip, DupAddrDetectTransmits, NameKey, vlan_Uint16, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.LinkLayerAddress, obj.Type, obj.Enabled, obj.Description, obj.vlan_String, obj.Ip, obj.DupAddrDetectTransmits, obj.NameKey, obj.vlan_Uint16, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6NeighborConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6NeighborConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6NeighborConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6NeighborConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6NeighborState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6NeighborState " +
		"( " +
		" Origin INTEGER " +
		" LinkLayerAddress TEXT " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" NameKey TEXT " +
		" Type TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" NeighborState INTEGER " +
		" Ip TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" vlan_Uint16 INTEGER " +
		" IsRouter bool " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6NeighborState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6NeighborState (Origin, LinkLayerAddress, Description, vlan_String, NameKey, Type, Enabled, Mtu, NeighborState, Ip, DupAddrDetectTransmits, vlan_Uint16, IsRouter) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.Origin, obj.LinkLayerAddress, obj.Description, obj.vlan_String, obj.NameKey, obj.Type, obj.Enabled, obj.Mtu, obj.NeighborState, obj.Ip, obj.DupAddrDetectTransmits, obj.vlan_Uint16, obj.IsRouter)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6NeighborState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6NeighborState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6NeighborState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6NeighborState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6Config) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6Config " +
		"( " +
		" DupAddrDetectTransmits INTEGER " +
		" Type TEXT " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6Config) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6Config (DupAddrDetectTransmits, Type, Description, vlan_String, Enabled, NameKey, vlan_Uint16, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.DupAddrDetectTransmits, obj.Type, obj.Description, obj.vlan_String, obj.Enabled, obj.NameKey, obj.vlan_Uint16, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6Config) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6Config where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6Config\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6Config) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6State) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6State " +
		"( " +
		" DupAddrDetectTransmits INTEGER " +
		" Type TEXT " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" Enabled bool " +
		" NameKey TEXT " +
		" vlan_Uint16 INTEGER " +
		" Mtu INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6State) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6State (DupAddrDetectTransmits, Type, Description, vlan_String, Enabled, NameKey, vlan_Uint16, Mtu) VALUES (%v, %v, %v, %v, %v, %v, %v, %v);",
		obj.DupAddrDetectTransmits, obj.Type, obj.Description, obj.vlan_String, obj.Enabled, obj.NameKey, obj.vlan_Uint16, obj.Mtu)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6State) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6State where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6State\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6State) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6AutoconfConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6AutoconfConfig " +
		"( " +
		" CreateGlobalAddresses bool " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" CreateTemporaryAddresses bool " +
		" Type TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" NameKey TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" vlan_Uint16 INTEGER " +
		" TemporaryValidLifetime INTEGER " +
		" TemporaryPreferredLifetime INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AutoconfConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6AutoconfConfig (CreateGlobalAddresses, Description, vlan_String, CreateTemporaryAddresses, Type, Enabled, Mtu, NameKey, DupAddrDetectTransmits, vlan_Uint16, TemporaryValidLifetime, TemporaryPreferredLifetime) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.CreateGlobalAddresses, obj.Description, obj.vlan_String, obj.CreateTemporaryAddresses, obj.Type, obj.Enabled, obj.Mtu, obj.NameKey, obj.DupAddrDetectTransmits, obj.vlan_Uint16, obj.TemporaryValidLifetime, obj.TemporaryPreferredLifetime)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6AutoconfConfig) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6AutoconfConfig where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6AutoconfConfig\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AutoconfConfig) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}

func (obj RoutedVlanIpv6AutoconfState) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS RoutedVlanIpv6AutoconfState " +
		"( " +
		" CreateGlobalAddresses bool " +
		" Description TEXT " +
		" vlan_String TEXT " +
		" CreateTemporaryAddresses bool " +
		" Type TEXT " +
		" Enabled bool " +
		" Mtu INTEGER " +
		" NameKey TEXT " +
		" DupAddrDetectTransmits INTEGER " +
		" vlan_Uint16 INTEGER " +
		" TemporaryValidLifetime INTEGER " +
		" TemporaryPreferredLifetime INTEGER " +
		"PRIMARY KEY(NameKey) ) "
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AutoconfState) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO RoutedVlanIpv6AutoconfState (CreateGlobalAddresses, Description, vlan_String, CreateTemporaryAddresses, Type, Enabled, Mtu, NameKey, DupAddrDetectTransmits, vlan_Uint16, TemporaryValidLifetime, TemporaryPreferredLifetime) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
		obj.CreateGlobalAddresses, obj.Description, obj.vlan_String, obj.CreateTemporaryAddresses, obj.Type, obj.Enabled, obj.Mtu, obj.NameKey, obj.DupAddrDetectTransmits, obj.vlan_Uint16, obj.TemporaryValidLifetime, obj.TemporaryPreferredLifetime)
	fmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}
	return objectId, err
}

func (obj RoutedVlanIpv6AutoconfState) DeleteObjectFromDb(objId string, dbHdl *sql.DB) error {
	dbCmd := "delete from RoutedVlanIpv6AutoconfState where rowid = " + objId
	fmt.Println("### DB Deleting RoutedVlanIpv6AutoconfState\n")
	_, err := ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}
func (obj RoutedVlanIpv6AutoconfState) GetKey() (string, error) {
	key := "NameKey = " + string(obj.NameKey)
	return key, nil
}
