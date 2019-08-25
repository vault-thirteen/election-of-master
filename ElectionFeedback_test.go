////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
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

//+build test

package eom

import (
	"testing"
	"time"

	"github.com/vault-thirteen/tester"
)

func Test_ParseRawData(t *testing.T) {

	var aTest *tester.Test
	var ef ElectionFeedback
	var err error
	var timeExpected time.Time

	aTest = tester.New(t)

	// Test #1. Junk.
	ef = ElectionFeedback{
		RawData: "xyz",
	}
	err = ef.ParseRawData()
	aTest.MustBeAnError(err)

	// Test #2. Bad Result.
	ef = ElectionFeedback{
		RawData: `(x,"something")`,
	}
	err = ef.ParseRawData()
	aTest.MustBeAnError(err)

	// Test #3. Bad Time.
	ef = ElectionFeedback{
		RawData: `(t,"something")`,
	}
	err = ef.ParseRawData()
	aTest.MustBeAnError(err)

	// Test #4. All clear.
	ef = ElectionFeedback{
		RawData: `(t,"1999-12-31 23:59:33+05")`,
	}
	err = ef.ParseRawData()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ef.Result, true)
	timeExpected, err = time.Parse(time.RFC3339, "1999-12-31T23:59:33+05:00")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ef.Time, timeExpected)
}
