package pgcodes

// IsSuccessfulCompletion returns true if the code equals "00000".
func IsSuccessfulCompletion(code string) bool {
	return code == SuccessfulCompletion
}

// IsWarning returns true if the code starts with "01".
func IsWarning(code string) bool {
	return len(code) >= 2 && code[:2] == "01"
}

// IsWarningDynamicResultSetsReturned returns true if the code equals "0100C".
func IsWarningDynamicResultSetsReturned(code string) bool {
	return code == WarningDynamicResultSetsReturned
}

// IsWarningImplicitZeroBitPadding returns true if the code equals "01008".
func IsWarningImplicitZeroBitPadding(code string) bool {
	return code == WarningImplicitZeroBitPadding
}

// IsWarningNullValueEliminatedInSetFunction returns true if the code equals "01003".
func IsWarningNullValueEliminatedInSetFunction(code string) bool {
	return code == WarningNullValueEliminatedInSetFunction
}

// IsWarningPrivilegeNotGranted returns true if the code equals "01007".
func IsWarningPrivilegeNotGranted(code string) bool {
	return code == WarningPrivilegeNotGranted
}

// IsWarningPrivilegeNotRevoked returns true if the code equals "01006".
func IsWarningPrivilegeNotRevoked(code string) bool {
	return code == WarningPrivilegeNotRevoked
}

// IsWarningStringDataRightTruncation returns true if the code equals "01004".
func IsWarningStringDataRightTruncation(code string) bool {
	return code == WarningStringDataRightTruncation
}

// IsNoData returns true if the code equals "02000".
func IsNoData(code string) bool {
	return code == NoData
}

// IsNoAdditionalDynamicResultSetsReturned returns true if the code equals "02001".
func IsNoAdditionalDynamicResultSetsReturned(code string) bool {
	return code == NoAdditionalDynamicResultSetsReturned
}

// IsSqlStatementNotYetComplete returns true if the code equals "03000".
func IsSqlStatementNotYetComplete(code string) bool {
	return code == SqlStatementNotYetComplete
}

// IsConnectionException returns true if the code starts with "08".
func IsConnectionException(code string) bool {
	return len(code) >= 2 && code[:2] == "08"
}

// IsConnectionDoesNotExist returns true if the code equals "08003".
func IsConnectionDoesNotExist(code string) bool {
	return code == ConnectionDoesNotExist
}

// IsConnectionFailure returns true if the code equals "08006".
func IsConnectionFailure(code string) bool {
	return code == ConnectionFailure
}

// IsSqlclientUnableToEstablishSqlconnection returns true if the code equals "08001".
func IsSqlclientUnableToEstablishSqlconnection(code string) bool {
	return code == SqlclientUnableToEstablishSqlconnection
}

// IsSqlserverRejectedEstablishmentOfSqlconnection returns true if the code equals "08004".
func IsSqlserverRejectedEstablishmentOfSqlconnection(code string) bool {
	return code == SqlserverRejectedEstablishmentOfSqlconnection
}

// IsTransactionResolutionUnknown returns true if the code equals "08007".
func IsTransactionResolutionUnknown(code string) bool {
	return code == TransactionResolutionUnknown
}

// IsProtocolViolation returns true if the code equals "08P01".
func IsProtocolViolation(code string) bool {
	return code == ProtocolViolation
}

// IsTriggeredActionException returns true if the code equals "09000".
func IsTriggeredActionException(code string) bool {
	return code == TriggeredActionException
}

// IsFeatureNotSupported returns true if the code equals "0A000".
func IsFeatureNotSupported(code string) bool {
	return code == FeatureNotSupported
}

// IsInvalidTransactionInitiation returns true if the code equals "0B000".
func IsInvalidTransactionInitiation(code string) bool {
	return code == InvalidTransactionInitiation
}

// IsLocatorException returns true if the code equals "0F000".
func IsLocatorException(code string) bool {
	return code == LocatorException
}

// IsInvalidLocatorSpecification returns true if the code equals "0F001".
func IsInvalidLocatorSpecification(code string) bool {
	return code == InvalidLocatorSpecification
}

// IsInvalidGrantor returns true if the code equals "0L000".
func IsInvalidGrantor(code string) bool {
	return code == InvalidGrantor
}

// IsInvalidGrantOperation returns true if the code equals "0LP01".
func IsInvalidGrantOperation(code string) bool {
	return code == InvalidGrantOperation
}

// IsInvalidRoleSpecification returns true if the code equals "0P000".
func IsInvalidRoleSpecification(code string) bool {
	return code == InvalidRoleSpecification
}

// IsDiagnosticsException returns true if the code equals "0Z000".
func IsDiagnosticsException(code string) bool {
	return code == DiagnosticsException
}

// IsStackOverflow returns true if the code equals "0Z001".
func IsStackOverflow(code string) bool {
	return code == StackOverflow
}

// IsStackUnderflow returns true if the code equals "0Z002".
func IsStackUnderflow(code string) bool {
	return code == StackUnderflow
}

// IsDataException returns true if the code equals "22000".
func IsDataException(code string) bool {
	return code == DataException
}

// IsNumericValueOutOfRange returns true if the code equals "22003".
func IsNumericValueOutOfRange(code string) bool {
	return code == NumericValueOutOfRange
}

// IsNullValueNotAllowed returns true if the code equals "22004".
func IsNullValueNotAllowed(code string) bool {
	return code == NullValueNotAllowed
}

// IsErrorInAssignment returns true if the code equals "22005".
func IsErrorInAssignment(code string) bool {
	return code == ErrorInAssignment
}

// IsInvalidDatetimeFormat returns true if the code equals "22007".
func IsInvalidDatetimeFormat(code string) bool {
	return code == InvalidDatetimeFormat
}

// IsDatetimeFieldOverflow returns true if the code equals "22008".
func IsDatetimeFieldOverflow(code string) bool {
	return code == DatetimeFieldOverflow
}

// IsDivisionByZero returns true if the code equals "22012".
func IsDivisionByZero(code string) bool {
	return code == DivisionByZero
}

// IsInvalidPrecedingOrFollowingSize returns true if the code equals "22013".
func IsInvalidPrecedingOrFollowingSize(code string) bool {
	return code == InvalidPrecedingOrFollowingSize
}

// IsFloatingPointException returns true if the code equals "22014".
func IsFloatingPointException(code string) bool {
	return code == FloatingPointException
}

// IsInvalidArgumentForLogarithm returns true if the code equals "22015".
func IsInvalidArgumentForLogarithm(code string) bool {
	return code == InvalidArgumentForLogarithm
}

// IsInvalidArgumentForPowerFunction returns true if the code equals "22016".
func IsInvalidArgumentForPowerFunction(code string) bool {
	return code == InvalidArgumentForPowerFunction
}

// IsInvalidArgumentForWidthBucketFunction returns true if the code equals "22017".
func IsInvalidArgumentForWidthBucketFunction(code string) bool {
	return code == InvalidArgumentForWidthBucketFunction
}

// IsInvalidCharacterValueForCast returns true if the code equals "22018".
func IsInvalidCharacterValueForCast(code string) bool {
	return code == InvalidCharacterValueForCast
}

// IsInvalidEscapeCharacter returns true if the code equals "22019".
func IsInvalidEscapeCharacter(code string) bool {
	return code == InvalidEscapeCharacter
}

// IsInvalidLimitValue returns true if the code equals "22020".
func IsInvalidLimitValue(code string) bool {
	return code == InvalidLimitValue
}

// IsNullValueNoIndicatorParameter returns true if the code equals "22021".
func IsNullValueNoIndicatorParameter(code string) bool {
	return code == NullValueNoIndicatorParameter
}

// IsUnterminatedCString returns true if the code equals "22022".
func IsUnterminatedCString(code string) bool {
	return code == UnterminatedCString
}

// IsInvalidEscapeSequence returns true if the code equals "22023".
func IsInvalidEscapeSequence(code string) bool {
	return code == InvalidEscapeSequence
}

// IsStringDataLengthMismatch returns true if the code equals "22024".
func IsStringDataLengthMismatch(code string) bool {
	return code == StringDataLengthMismatch
}

// IsStringDataRightTruncation returns true if the code equals "22025".
func IsStringDataRightTruncation(code string) bool {
	return code == StringDataRightTruncation
}

// IsSubstringError returns true if the code equals "22026".
func IsSubstringError(code string) bool {
	return code == SubstringError
}

// IsTrimError returns true if the code equals "22027".
func IsTrimError(code string) bool {
	return code == TrimError
}

// IsCharacterNotInRepertoire returns true if the code equals "22028".
func IsCharacterNotInRepertoire(code string) bool {
	return code == CharacterNotInRepertoire
}

// IsIndicatorOverflow returns true if the code equals "22029".
func IsIndicatorOverflow(code string) bool {
	return code == IndicatorOverflow
}

// IsInvalidParameterValue returns true if the code equals "22030".
func IsInvalidParameterValue(code string) bool {
	return code == InvalidParameterValue
}

// IsUntranslatableCharacter returns true if the code equals "22031".
func IsUntranslatableCharacter(code string) bool {
	return code == UntranslatableCharacter
}

// IsNonstandardUseOfEscapeCharacter returns true if the code equals "22032".
func IsNonstandardUseOfEscapeCharacter(code string) bool {
	return code == NonstandardUseOfEscapeCharacter
}

// IsIntegrityConstraintViolation returns true if the code equals "23000".
func IsIntegrityConstraintViolation(code string) bool {
	return code == IntegrityConstraintViolation
}

// IsRestrictViolation returns true if the code equals "23001".
func IsRestrictViolation(code string) bool {
	return code == RestrictViolation
}

// IsNotNullViolation returns true if the code equals "23502".
func IsNotNullViolation(code string) bool {
	return code == NotNullViolation
}

// IsForeignKeyViolation returns true if the code equals "23503".
func IsForeignKeyViolation(code string) bool {
	return code == ForeignKeyViolation
}

// IsUniqueViolation returns true if the code equals "23505".
func IsUniqueViolation(code string) bool {
	return code == UniqueViolation
}

// IsCheckViolation returns true if the code equals "23514".
func IsCheckViolation(code string) bool {
	return code == CheckViolation
}

// IsExclusionViolation returns true if the code equals "23P01".
func IsExclusionViolation(code string) bool {
	return code == ExclusionViolation
}

// IsInvalidTransactionState returns true if the code equals "25000".
func IsInvalidTransactionState(code string) bool {
	return code == InvalidTransactionState
}

// IsActiveSqlTransaction returns true if the code equals "25001".
func IsActiveSqlTransaction(code string) bool {
	return code == ActiveSqlTransaction
}

// IsBranchTransactionAlreadyActive returns true if the code equals "25002".
func IsBranchTransactionAlreadyActive(code string) bool {
	return code == BranchTransactionAlreadyActive
}

// IsHeldCursorRequiresSameIsolationLevel returns true if the code equals "25003".
func IsHeldCursorRequiresSameIsolationLevel(code string) bool {
	return code == HeldCursorRequiresSameIsolationLevel
}

// IsInappropriateAccessModeForBranchTransaction returns true if the code equals "25004".
func IsInappropriateAccessModeForBranchTransaction(code string) bool {
	return code == InappropriateAccessModeForBranchTransaction
}

// IsInappropriateIsolationLevelForBranchTransaction returns true if the code equals "25005".
func IsInappropriateIsolationLevelForBranchTransaction(code string) bool {
	return code == InappropriateIsolationLevelForBranchTransaction
}

// IsNoActiveSqlTransactionForBranchTransaction returns true if the code equals "25006".
func IsNoActiveSqlTransactionForBranchTransaction(code string) bool {
	return code == NoActiveSqlTransactionForBranchTransaction
}

// IsReadOnlySqlTransaction returns true if the code equals "25007".
func IsReadOnlySqlTransaction(code string) bool {
	return code == ReadOnlySqlTransaction
}

// IsSchemaAndDataStatementMixingNotSupported returns true if the code equals "25008".
func IsSchemaAndDataStatementMixingNotSupported(code string) bool {
	return code == SchemaAndDataStatementMixingNotSupported
}

// IsNoActiveSqlTransaction returns true if the code equals "25009".
func IsNoActiveSqlTransaction(code string) bool {
	return code == NoActiveSqlTransaction
}

// IsInFailedSqlTransaction returns true if the code equals "25010".
func IsInFailedSqlTransaction(code string) bool {
	return code == InFailedSqlTransaction
}

// IsIdleInTransactionSessionTimeout returns true if the code equals "25P01".
func IsIdleInTransactionSessionTimeout(code string) bool {
	return code == IdleInTransactionSessionTimeout
}

// IsTransactionTimeout returns true if the code equals "25P02".
func IsTransactionTimeout(code string) bool {
	return code == TransactionTimeout
}

// IsInvalidAuthorizationSpecification returns true if the code equals "28000".
func IsInvalidAuthorizationSpecification(code string) bool {
	return code == InvalidAuthorizationSpecification
}

// IsInvalidPassword returns true if the code equals "28P01".
func IsInvalidPassword(code string) bool {
	return code == InvalidPassword
}

// IsDependentPrivilegeDescriptorsStillExist returns true if the code equals "2B000".
func IsDependentPrivilegeDescriptorsStillExist(code string) bool {
	return code == DependentPrivilegeDescriptorsStillExist
}

// IsDependentObjectsStillExist returns true if the code equals "2BP01".
func IsDependentObjectsStillExist(code string) bool {
	return code == DependentObjectsStillExist
}

// IsInvalidTransactionTermination returns true if the code equals "2D000".
func IsInvalidTransactionTermination(code string) bool {
	return code == InvalidTransactionTermination
}

// IsSqlRoutineException returns true if the code equals "2F000".
func IsSqlRoutineException(code string) bool {
	return code == SqlRoutineException
}

// IsModifyingSqlDataNotPermitted returns true if the code equals "2F002".
func IsModifyingSqlDataNotPermitted(code string) bool {
	return code == ModifyingSqlDataNotPermitted
}

// IsProhibitedSqlStatementAttempted returns true if the code equals "2F003".
func IsProhibitedSqlStatementAttempted(code string) bool {
	return code == ProhibitedSqlStatementAttempted
}

// IsReadingSqlDataNotPermitted returns true if the code equals "2F004".
func IsReadingSqlDataNotPermitted(code string) bool {
	return code == ReadingSqlDataNotPermitted
}

// IsFunctionExecutedNoReturnStatement returns true if the code equals "2F005".
func IsFunctionExecutedNoReturnStatement(code string) bool {
	return code == FunctionExecutedNoReturnStatement
}

// IsInvalidCursorName returns true if the code equals "34000".
func IsInvalidCursorName(code string) bool {
	return code == InvalidCursorName
}

// IsExternalRoutineException returns true if the code equals "38000".
func IsExternalRoutineException(code string) bool {
	return code == ExternalRoutineException
}

// IsContainingSqlNotPermitted returns true if the code equals "38001".
func IsContainingSqlNotPermitted(code string) bool {
	return code == ContainingSqlNotPermitted
}

// IsModifyingSqlDataNotPermittedExternal returns true if the code equals "38002".
func IsModifyingSqlDataNotPermittedExternal(code string) bool {
	return code == ModifyingSqlDataNotPermittedExternal
}

// IsProhibitedSqlStatementAttemptedExternal returns true if the code equals "38003".
func IsProhibitedSqlStatementAttemptedExternal(code string) bool {
	return code == ProhibitedSqlStatementAttemptedExternal
}

// IsReadingSqlDataNotPermittedExternal returns true if the code equals "38004".
func IsReadingSqlDataNotPermittedExternal(code string) bool {
	return code == ReadingSqlDataNotPermittedExternal
}

// IsExternalRoutineInvocationException returns true if the code equals "39000".
func IsExternalRoutineInvocationException(code string) bool {
	return code == ExternalRoutineInvocationException
}

// IsInvalidSqlstateReturned returns true if the code equals "39001".
func IsInvalidSqlstateReturned(code string) bool {
	return code == InvalidSqlstateReturned
}

// IsNullValueNotAllowedExternal returns true if the code equals "39004".
func IsNullValueNotAllowedExternal(code string) bool {
	return code == NullValueNotAllowedExternal
}

// IsTriggerProtocolViolated returns true if the code equals "39P01".
func IsTriggerProtocolViolated(code string) bool {
	return code == TriggerProtocolViolated
}

// IsSrfProtocolViolated returns true if the code equals "39P02".
func IsSrfProtocolViolated(code string) bool {
	return code == SrfProtocolViolated
}

// IsSavepointException returns true if the code equals "3B000".
func IsSavepointException(code string) bool {
	return code == SavepointException
}

// IsInvalidSavepointSpecification returns true if the code equals "3B001".
func IsInvalidSavepointSpecification(code string) bool {
	return code == InvalidSavepointSpecification
}

// IsInvalidCatalogName returns true if the code equals "3D000".
func IsInvalidCatalogName(code string) bool {
	return code == InvalidCatalogName
}

// IsInvalidSchemaName returns true if the code equals "3F000".
func IsInvalidSchemaName(code string) bool {
	return code == InvalidSchemaName
}

// IsTransactionRollback returns true if the code equals "40000".
func IsTransactionRollback(code string) bool {
	return code == TransactionRollback
}

// IsTransactionIntegrityConstraintViolation returns true if the code equals "40001".
func IsTransactionIntegrityConstraintViolation(code string) bool {
	return code == TransactionIntegrityConstraintViolation
}

// IsSerializationFailure returns true if the code equals "40002".
func IsSerializationFailure(code string) bool {
	return code == SerializationFailure
}

// IsStatementCompletionUnknown returns true if the code equals "40003".
func IsStatementCompletionUnknown(code string) bool {
	return code == StatementCompletionUnknown
}

// IsDeadlockDetected returns true if the code equals "40004".
func IsDeadlockDetected(code string) bool {
	return code == DeadlockDetected
}

// IsSyntaxErrorOrAccessRuleViolation returns true if the code equals "42000".
func IsSyntaxErrorOrAccessRuleViolation(code string) bool {
	return code == SyntaxErrorOrAccessRuleViolation
}

// IsSyntaxError returns true if the code equals "42601".
func IsSyntaxError(code string) bool {
	return code == SyntaxError
}

// IsInsufficientPrivilege returns true if the code equals "42501".
func IsInsufficientPrivilege(code string) bool {
	return code == InsufficientPrivilege
}

// IsCannotCoerce returns true if the code equals "42846".
func IsCannotCoerce(code string) bool {
	return code == CannotCoerce
}

// IsGroupingError returns true if the code equals "42803".
func IsGroupingError(code string) bool {
	return code == GroupingError
}

// IsWindowingError returns true if the code equals "42P20".
func IsWindowingError(code string) bool {
	return code == WindowingError
}

// IsInvalidRecursion returns true if the code equals "42P19".
func IsInvalidRecursion(code string) bool {
	return code == InvalidRecursion
}

// IsInvalidForeignKey returns true if the code equals "42830".
func IsInvalidForeignKey(code string) bool {
	return code == InvalidForeignKey
}

// IsInvalidName returns true if the code equals "42602".
func IsInvalidName(code string) bool {
	return code == InvalidName
}

// IsNameTooLong returns true if the code equals "42622".
func IsNameTooLong(code string) bool {
	return code == NameTooLong
}

// IsReservedName returns true if the code equals "42939".
func IsReservedName(code string) bool {
	return code == ReservedName
}

// IsDatatypeMismatch returns true if the code equals "42804".
func IsDatatypeMismatch(code string) bool {
	return code == DatatypeMismatch
}

// IsIndeterminateDatatype returns true if the code equals "42P18".
func IsIndeterminateDatatype(code string) bool {
	return code == IndeterminateDatatype
}

// IsCollationMismatch returns true if the code equals "42P21".
func IsCollationMismatch(code string) bool {
	return code == CollationMismatch
}

// IsIndeterminateCollation returns true if the code equals "42P22".
func IsIndeterminateCollation(code string) bool {
	return code == IndeterminateCollation
}

// IsWrongObjectType returns true if the code equals "42809".
func IsWrongObjectType(code string) bool {
	return code == WrongObjectType
}

// IsGeneratedAlways returns true if the code equals "428C9".
func IsGeneratedAlways(code string) bool {
	return code == GeneratedAlways
}

// IsUndefinedColumn returns true if the code equals "42703".
func IsUndefinedColumn(code string) bool {
	return code == UndefinedColumn
}

// IsUndefinedFunction returns true if the code equals "42883".
func IsUndefinedFunction(code string) bool {
	return code == UndefinedFunction
}

// IsUndefinedTable returns true if the code equals "42P01".
func IsUndefinedTable(code string) bool {
	return code == UndefinedTable
}

// IsUndefinedParameter returns true if the code equals "42P02".
func IsUndefinedParameter(code string) bool {
	return code == UndefinedParameter
}

// IsUndefinedObject returns true if the code equals "42704".
func IsUndefinedObject(code string) bool {
	return code == UndefinedObject
}

// IsDuplicateColumn returns true if the code equals "42701".
func IsDuplicateColumn(code string) bool {
	return code == DuplicateColumn
}

// IsDuplicateCursor returns true if the code equals "42P03".
func IsDuplicateCursor(code string) bool {
	return code == DuplicateCursor
}

// IsDuplicateDatabase returns true if the code equals "42P04".
func IsDuplicateDatabase(code string) bool {
	return code == DuplicateDatabase
}

// IsDuplicateFunction returns true if the code equals "42723".
func IsDuplicateFunction(code string) bool {
	return code == DuplicateFunction
}

// IsDuplicatePreparedStatement returns true if the code equals "42P05".
func IsDuplicatePreparedStatement(code string) bool {
	return code == DuplicatePreparedStatement
}

// IsDuplicateSchema returns true if the code equals "42P06".
func IsDuplicateSchema(code string) bool {
	return code == DuplicateSchema
}

// IsDuplicateTable returns true if the code equals "42P07".
func IsDuplicateTable(code string) bool {
	return code == DuplicateTable
}

// IsDuplicateAlias returns true if the code equals "42712".
func IsDuplicateAlias(code string) bool {
	return code == DuplicateAlias
}

// IsDuplicateObject returns true if the code equals "42710".
func IsDuplicateObject(code string) bool {
	return code == DuplicateObject
}

// IsAmbiguousColumn returns true if the code equals "42702".
func IsAmbiguousColumn(code string) bool {
	return code == AmbiguousColumn
}

// IsAmbiguousFunction returns true if the code equals "42725".
func IsAmbiguousFunction(code string) bool {
	return code == AmbiguousFunction
}

// IsAmbiguousParameter returns true if the code equals "42P08".
func IsAmbiguousParameter(code string) bool {
	return code == AmbiguousParameter
}

// IsAmbiguousAlias returns true if the code equals "42P09".
func IsAmbiguousAlias(code string) bool {
	return code == AmbiguousAlias
}

// IsInvalidColumnReference returns true if the code equals "42P10".
func IsInvalidColumnReference(code string) bool {
	return code == InvalidColumnReference
}

// IsInvalidColumnDefinition returns true if the code equals "42611".
func IsInvalidColumnDefinition(code string) bool {
	return code == InvalidColumnDefinition
}

// IsInvalidCursorDefinition returns true if the code equals "42P11".
func IsInvalidCursorDefinition(code string) bool {
	return code == InvalidCursorDefinition
}

// IsInvalidDatabaseDefinition returns true if the code equals "42P12".
func IsInvalidDatabaseDefinition(code string) bool {
	return code == InvalidDatabaseDefinition
}

// IsInvalidFunctionDefinition returns true if the code equals "42P13".
func IsInvalidFunctionDefinition(code string) bool {
	return code == InvalidFunctionDefinition
}

// IsInvalidPreparedStatementDefinition returns true if the code equals "42P14".
func IsInvalidPreparedStatementDefinition(code string) bool {
	return code == InvalidPreparedStatementDefinition
}

// IsInvalidSchemaDefinition returns true if the code equals "42P15".
func IsInvalidSchemaDefinition(code string) bool {
	return code == InvalidSchemaDefinition
}

// IsInvalidTableDefinition returns true if the code equals "42P16".
func IsInvalidTableDefinition(code string) bool {
	return code == InvalidTableDefinition
}

// IsInvalidObjectDefinition returns true if the code equals "42P17".
func IsInvalidObjectDefinition(code string) bool {
	return code == InvalidObjectDefinition
}

// IsWithCheckOptionViolation returns true if the code equals "44000".
func IsWithCheckOptionViolation(code string) bool {
	return code == WithCheckOptionViolation
}

// IsInsufficientResources returns true if the code equals "53000".
func IsInsufficientResources(code string) bool {
	return code == InsufficientResources
}

// IsDiskFull returns true if the code equals "53100".
func IsDiskFull(code string) bool {
	return code == DiskFull
}

// IsOutOfMemory returns true if the code equals "53200".
func IsOutOfMemory(code string) bool {
	return code == OutOfMemory
}

// IsTooManyConnections returns true if the code equals "53300".
func IsTooManyConnections(code string) bool {
	return code == TooManyConnections
}

// IsConfigurationLimitExceeded returns true if the code equals "53400".
func IsConfigurationLimitExceeded(code string) bool {
	return code == ConfigurationLimitExceeded
}

// IsProgramLimitExceeded returns true if the code equals "54000".
func IsProgramLimitExceeded(code string) bool {
	return code == ProgramLimitExceeded
}

// IsStatementTooComplex returns true if the code equals "54001".
func IsStatementTooComplex(code string) bool {
	return code == StatementTooComplex
}

// IsTooManyColumns returns true if the code equals "54011".
func IsTooManyColumns(code string) bool {
	return code == TooManyColumns
}

// IsTooManyArguments returns true if the code equals "54023".
func IsTooManyArguments(code string) bool {
	return code == TooManyArguments
}

// IsObjectNotInPrerequisiteState returns true if the code equals "55000".
func IsObjectNotInPrerequisiteState(code string) bool {
	return code == ObjectNotInPrerequisiteState
}

// IsObjectInUse returns true if the code equals "55006".
func IsObjectInUse(code string) bool {
	return code == ObjectInUse
}

// IsCantChangeRuntimeParam returns true if the code equals "55P02".
func IsCantChangeRuntimeParam(code string) bool {
	return code == CantChangeRuntimeParam
}

// IsLockNotAvailable returns true if the code equals "55P03".
func IsLockNotAvailable(code string) bool {
	return code == LockNotAvailable
}

// IsOperatorIntervention returns true if the code equals "57000".
func IsOperatorIntervention(code string) bool {
	return code == OperatorIntervention
}

// IsQueryCanceled returns true if the code equals "57014".
func IsQueryCanceled(code string) bool {
	return code == QueryCanceled
}

// IsAdminShutdown returns true if the code equals "57P01".
func IsAdminShutdown(code string) bool {
	return code == AdminShutdown
}

// IsCrashShutdown returns true if the code equals "57P02".
func IsCrashShutdown(code string) bool {
	return code == CrashShutdown
}

// IsCannotConnectNow returns true if the code equals "57P03".
func IsCannotConnectNow(code string) bool {
	return code == CannotConnectNow
}

// IsDatabaseDropped returns true if the code equals "57P04".
func IsDatabaseDropped(code string) bool {
	return code == DatabaseDropped
}

// IsSystemError returns true if the code equals "58000".
func IsSystemError(code string) bool {
	return code == SystemError
}

// IsIoError returns true if the code equals "58030".
func IsIoError(code string) bool {
	return code == IoError
}

// IsUndefinedFile returns true if the code equals "58031".
func IsUndefinedFile(code string) bool {
	return code == UndefinedFile
}

// IsDuplicateFile returns true if the code equals "58032".
func IsDuplicateFile(code string) bool {
	return code == DuplicateFile
}

// IsConfigFileError returns true if the code equals "F0000".
func IsConfigFileError(code string) bool {
	return code == ConfigFileError
}

// IsLockFileExists returns true if the code equals "F0001".
func IsLockFileExists(code string) bool {
	return code == LockFileExists
}

// IsFdwError returns true if the code equals "HV000".
func IsFdwError(code string) bool {
	return code == FdwError
}

// IsFdwColumnNameNotFound returns true if the code equals "HV005".
func IsFdwColumnNameNotFound(code string) bool {
	return code == FdwColumnNameNotFound
}

// IsFdwDynamicParameterValueNeeded returns true if the code equals "HV002".
func IsFdwDynamicParameterValueNeeded(code string) bool {
	return code == FdwDynamicParameterValueNeeded
}

// IsFdwFunctionSequenceError returns true if the code equals "HV010".
func IsFdwFunctionSequenceError(code string) bool {
	return code == FdwFunctionSequenceError
}

// IsFdwInconsistentDescriptorInformation returns true if the code equals "HV021".
func IsFdwInconsistentDescriptorInformation(code string) bool {
	return code == FdwInconsistentDescriptorInformation
}

// IsFdwInvalidAttributeValue returns true if the code equals "HV024".
func IsFdwInvalidAttributeValue(code string) bool {
	return code == FdwInvalidAttributeValue
}

// IsFdwInvalidColumnName returns true if the code equals "HV007".
func IsFdwInvalidColumnName(code string) bool {
	return code == FdwInvalidColumnName
}

// IsFdwInvalidColumnNumber returns true if the code equals "HV008".
func IsFdwInvalidColumnNumber(code string) bool {
	return code == FdwInvalidColumnNumber
}

// IsFdwInvalidDataType returns true if the code equals "HV004".
func IsFdwInvalidDataType(code string) bool {
	return code == FdwInvalidDataType
}

// IsFdwInvalidDataTypeDescriptors returns true if the code equals "HV006".
func IsFdwInvalidDataTypeDescriptors(code string) bool {
	return code == FdwInvalidDataTypeDescriptors
}

// IsFdwInvalidDescriptorFieldIdentifier returns true if the code equals "HV091".
func IsFdwInvalidDescriptorFieldIdentifier(code string) bool {
	return code == FdwInvalidDescriptorFieldIdentifier
}

// IsFdwInvalidHandle returns true if the code equals "HV00B".
func IsFdwInvalidHandle(code string) bool {
	return code == FdwInvalidHandle
}

// IsFdwInvalidOptionIndex returns true if the code equals "HV00C".
func IsFdwInvalidOptionIndex(code string) bool {
	return code == FdwInvalidOptionIndex
}

// IsFdwInvalidOptionName returns true if the code equals "HV00D".
func IsFdwInvalidOptionName(code string) bool {
	return code == FdwInvalidOptionName
}

// IsFdwInvalidStringLengthOrBufferLength returns true if the code equals "HV090".
func IsFdwInvalidStringLengthOrBufferLength(code string) bool {
	return code == FdwInvalidStringLengthOrBufferLength
}

// IsFdwInvalidStringFormat returns true if the code equals "HV00A".
func IsFdwInvalidStringFormat(code string) bool {
	return code == FdwInvalidStringFormat
}

// IsFdwInvalidUseOfNullPointer returns true if the code equals "HV009".
func IsFdwInvalidUseOfNullPointer(code string) bool {
	return code == FdwInvalidUseOfNullPointer
}

// IsFdwTooManyHandles returns true if the code equals "HV014".
func IsFdwTooManyHandles(code string) bool {
	return code == FdwTooManyHandles
}

// IsFdwOutOfMemory returns true if the code equals "HV001".
func IsFdwOutOfMemory(code string) bool {
	return code == FdwOutOfMemory
}

// IsFdwNoSchemas returns true if the code equals "HV00P".
func IsFdwNoSchemas(code string) bool {
	return code == FdwNoSchemas
}

// IsFdwOptionNameNotFound returns true if the code equals "HV00J".
func IsFdwOptionNameNotFound(code string) bool {
	return code == FdwOptionNameNotFound
}

// IsFdwReplyHandle returns true if the code equals "HV00K".
func IsFdwReplyHandle(code string) bool {
	return code == FdwReplyHandle
}

// IsFdwSchemaNotFound returns true if the code equals "HV00Q".
func IsFdwSchemaNotFound(code string) bool {
	return code == FdwSchemaNotFound
}

// IsFdwTableNotFound returns true if the code equals "HV00R".
func IsFdwTableNotFound(code string) bool {
	return code == FdwTableNotFound
}

// IsFdwUnableToCreateExecution returns true if the code equals "HV00L".
func IsFdwUnableToCreateExecution(code string) bool {
	return code == FdwUnableToCreateExecution
}

// IsFdwUnableToCreateReply returns true if the code equals "HV00M".
func IsFdwUnableToCreateReply(code string) bool {
	return code == FdwUnableToCreateReply
}

// IsFdwUnableToEstablishConnection returns true if the code equals "HV00N".
func IsFdwUnableToEstablishConnection(code string) bool {
	return code == FdwUnableToEstablishConnection
}

// IsPlpgSqlError returns true if the code equals "P0000".
func IsPlpgSqlError(code string) bool {
	return code == PlpgSqlError
}

// IsRaiseException returns true if the code equals "P0001".
func IsRaiseException(code string) bool {
	return code == RaiseException
}

// IsNoDataFound returns true if the code equals "P0002".
func IsNoDataFound(code string) bool {
	return code == NoDataFound
}

// IsTooManyRows returns true if the code equals "P0003".
func IsTooManyRows(code string) bool {
	return code == TooManyRows
}

// IsAssertFailure returns true if the code equals "P0004".
func IsAssertFailure(code string) bool {
	return code == AssertFailure
}

// IsInternalError returns true if the code equals "XX000".
func IsInternalError(code string) bool {
	return code == InternalError
}

// IsDataCorrupted returns true if the code equals "XX001".
func IsDataCorrupted(code string) bool {
	return code == DataCorrupted
}

// IsIndexCorrupted returns true if the code equals "XX002".
func IsIndexCorrupted(code string) bool {
	return code == IndexCorrupted
}
