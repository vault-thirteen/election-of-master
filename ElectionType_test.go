// ElectionType_test.go.

//+build test

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package eom

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_IsValid(t *testing.T) {

	var aTest *tester.Test
	var et ElectionType

	aTest = tester.New(t)

	// Test #1. Junk.
	et = ElectionType(0)
	aTest.MustBeEqual(et.IsValid(), false)

	// Test #1. Normal Type.
	et = ElectionType(ElectionTypeSingleMaster)
	aTest.MustBeEqual(et.IsValid(), true)
}
