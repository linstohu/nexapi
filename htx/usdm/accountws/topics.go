/*
 * Copyright (c) 2023, LinstoHu
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package accountws

import (
	"fmt"
)

func (m *AccountWsClient) GetIsolatedAccountUpdateTopic(contractCode string) (string, error) {
	if contractCode == "" {
		return "", fmt.Errorf("the contract_code field must be provided")
	}

	return fmt.Sprintf("accounts.%s", contractCode), nil
}

func (m *AccountWsClient) GetCrossAccountUpdateTopic(marginAccount string) (string, error) {
	if marginAccount == "" {
		return "", fmt.Errorf("the margin_account field must be provided")
	}

	return fmt.Sprintf("accounts_cross.%s", marginAccount), nil
}

func (m *AccountWsClient) GetUnifyAccountUpdateTopic() (string, error) {
	return "accounts_unify.USDT", nil
}
