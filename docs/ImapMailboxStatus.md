# ImapMailboxStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The mailbox name. | 
**ReadOnly** | **bool** | True if the mailbox is open in read-only mode. | 
**Items** | Pointer to **map[string]map[string]interface{}** | Results map | 
**Flags** | Pointer to **[]string** | The mailbox flags. | 
**PermanentFlags** | Pointer to **[]string** | The mailbox permanent flags. | 
**UnseenSeqNum** | **int64** | The sequence number of the first unseen message in the mailbox. | 
**Messages** | **int32** | The number of messages in this mailbox. | 
**Recent** | **int32** | The number of messages not seen since the last time the mailbox was opened. | 
**Unseen** | **int32** | The number of unread messages. | 
**UidNext** | **int64** | The next UID. | 
**UidValidity** | **int32** | Together with a UID, it is a unique identifier for a message. Must be greater than or equal to 1. | 
**AppendLimit** | Pointer to **int32** | Per-mailbox limit of message size. Set only if server supports the APPENDLIMIT extension | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


