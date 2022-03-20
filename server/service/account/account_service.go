/*
 * Copyright Rivtower Technologies LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package account

import (
	"context"

	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pb "github.com/cita-cloud/operator-proxy/api/account"
	"github.com/cita-cloud/operator-proxy/server/kubeapi"
)

var _ pb.AccountServiceServer = &accountServer{}

type accountServer struct {
}

func NewAccountServer() pb.AccountServiceServer {
	return &accountServer{}
}

func (a accountServer) CreateAccount(ctx context.Context, account *pb.Account) (*pb.Account, error) {
	accountCr := &citacloudv1.Account{}
	accountCr.Name = account.GetName()
	accountCr.Namespace = account.GetNamespace()
	accountCr.Spec.Chain = account.GetChain()
	accountCr.Spec.KmsPassword = account.GetKmsPassword()
	accountCr.Spec.Role = convertProtoToSpec(account.Role)
	accountCr.Spec.Domain = account.Domain
	accountCr.Spec.Address = account.Address

	err := kubeapi.K8sClient.Create(ctx, accountCr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account cr", err)
	}
	return &pb.Account{
		Name:      account.GetName(),
		Namespace: account.GetNamespace(),
	}, status.New(codes.OK, "").Err()
}

func (a accountServer) ListAccount(ctx context.Context, request *pb.ListAccountRequest) (*pb.AccountList, error) {
	accountCrList := &citacloudv1.AccountList{}
	accountCrOpts := []client.ListOption{
		client.InNamespace(request.GetNamespace()),
	}
	if request.GetChain() != "" {
		accountCrOpts = append(accountCrOpts, client.MatchingFields{"spec.chain": request.GetChain()})
	}
	if err := kubeapi.K8sClient.List(ctx, accountCrList, accountCrOpts...); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list account cr", err)
	}
	accountList := make([]*pb.Account, 0)
	for _, accountCr := range accountCrList.Items {
		ac := &pb.Account{
			Name:        accountCr.Name,
			Namespace:   accountCr.Namespace,
			Chain:       accountCr.Spec.Chain,
			KmsPassword: accountCr.Spec.KmsPassword,
			Role:        convertSpecToProto(accountCr.Spec.Role),
			Domain:      accountCr.Spec.Domain,
		}
		accountList = append(accountList, ac)
	}
	return &pb.AccountList{Accounts: accountList}, status.New(codes.OK, "").Err()
}

func convertSpecToProto(role citacloudv1.Role) pb.Role {
	if role == citacloudv1.Admin {
		return pb.Role_Admin
	} else if role == citacloudv1.Consensus {
		return pb.Role_Consensus
	} else {
		return pb.Role_Ordinary
	}
}

func convertProtoToSpec(role pb.Role) citacloudv1.Role {
	if role == pb.Role_Admin {
		return citacloudv1.Admin
	} else if role == pb.Role_Consensus {
		return citacloudv1.Consensus
	} else {
		return citacloudv1.Ordinary
	}
}

func GetAdminAccountByChain(ctx context.Context, namespace, chainName string) (*pb.Account, error) {
	accountCrList := &citacloudv1.AccountList{}
	accountCrOpts := []client.ListOption{
		client.InNamespace(namespace),
	}
	accountCrOpts = append(accountCrOpts, client.MatchingFields{"spec.chain": chainName})
	if err := kubeapi.K8sClient.List(ctx, accountCrList, accountCrOpts...); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list account cr", err)
	}
	if len(accountCrList.Items) == 0 {
		return nil, nil
	}
	for _, account := range accountCrList.Items {
		if account.Spec.Role == citacloudv1.Admin {
			res := &pb.Account{
				Name:        account.Name,
				Namespace:   account.Namespace,
				Chain:       chainName,
				KmsPassword: account.Spec.KmsPassword,
				Role:        pb.Role_Admin,
				Domain:      account.Spec.Domain,
			}
			return res, nil
		}
	}
	return nil, nil
}
